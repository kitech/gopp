package gopp

import (
	"fmt"
	"log"
	mrand "math/rand"
	"slices"
	"sync"

	"github.com/dolthub/maphash"
	// _ "github.com/elliotchance/orderedmap/v2"
)

// hash 任意类型，来自go内部实现
// https://github.com/dolthub/maphash

const ListMap0_SHARD_COUNT = 32

// ordered map, thread safe, reversemap
type ListMap0[KT comparable, VT comparable] struct {
	mu sync.RWMutex
	m0 map[uint64]*Pair[KT, VT]
	a0 []uint64                 // todo how know map key hash value
	mr map[uint64]*Pair[KT, VT] // hash(val) => Pair
	// mr2 map[VT][]KT // TODO, reverse map conflict

	reversemap bool
	lockless   bool // todo

	hhker maphash.Hasher[KT] // todo
	hhver maphash.Hasher[VT]
}

type Stringer interface {
	fmt.Stringer
	comparable
}

func ListMap0NewAny[VT comparable]() *ListMap0[any, VT] {
	return ListMap0New[any, VT]()
}
func ListMap0NewInt[VT comparable]() *ListMap0[int, VT] {
	return ListMap0New[int, VT]()
}
func ListMap0NewStr[VT comparable]() *ListMap0[string, VT] {
	return ListMap0New[string, VT]()
}
func ListMap0Newr[KT comparable, VT comparable]() *ListMap0[KT, VT] {
	me := ListMap0New[KT, VT]()
	me.reversemap = true

	return me
}

func ListMap0New[KT comparable, VT comparable]() *ListMap0[KT, VT] {
	me := &ListMap0[KT, VT]{}
	me.m0 = make(map[uint64]*Pair[KT, VT], 8)
	me.mr = make(map[uint64]*Pair[KT, VT], 8)
	me.a0 = make([]uint64, 0, 8)
	me.hhker = maphash.Hasher[KT](maphash.NewHasher[KT]())
	me.hhver = maphash.Hasher[VT](maphash.NewHasher[VT]())

	return me
}

// /////////
func ListMap0From[KT comparable, VT comparable](m map[KT]VT) *ListMap0[KT, VT] {
	me := ListMap0New[KT, VT]()
	me.PutMany(m)
	return me
}

/////////

func (me *ListMap0[KT, VT]) Count() int { return me.Len() }
func (me *ListMap0[KT, VT]) Len() int {
	me.mu.RLock()
	l := len(me.a0)
	me.mu.RUnlock()
	return l
}
func (me *ListMap0[KT, VT]) lennolock() int {
	lena0 := len(me.a0)
	lenm0 := len(me.m0)
	Assert(lena0 == lena0, "a0/m0 not consistant", lena0, lenm0)
	return IfElse2(mrand.Int()%2 == 0, lena0, lenm0)
}

func (me *ListMap0[KT, VT]) putnolock(key KT, val VT) (exist, ok bool) {
	hkey := me.hhker.Hash(key)

	kv, exist := me.m0[hkey]
	if exist {
		kv.Val = val
		// 25530 ns/op
		// 应该是slice index 效率差
		// idx := slices.Index(me.a0, hkey)         // olog(n)
		// me.a0 = slices.Delete(me.a0, idx, idx+1) // olog(n)
		// _ = idx
	} else {
		kv = PairNew(key, val)
		me.m0[hkey] = kv
		me.a0 = append(me.a0, hkey)
	}

	ok = true
	if me.reversemap {
		hval := me.hhver.Hash(val)
		me.mr[hval] = kv
	}

	return
}

func (me *ListMap0[KT, VT]) Put(key KT, val VT) (ok bool) {
	me.mu.Lock()
	exist, ok := me.putnolock(key, val)
	_, _ = exist, ok
	me.mu.Unlock()

	return
}
func (me *ListMap0[KT, VT]) PutMany(kvs map[KT]VT) {
	me.mu.Lock()
	for k, v := range kvs {
		me.putnolock(k, v)
	}
	me.mu.Unlock()

	return
}
func (me *ListMap0[KT, VT]) PutMany3(keys []KT, vals []VT) {
	if len(keys) != len(vals) {
		log.Println("keys and vals len not same", len(keys), len(vals))
		return
	}
	me.mu.Lock()
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := vals[i]
		me.putnolock(key, val)
	}

	me.mu.Unlock()

	return
}
func (me *ListMap0[KT, VT]) PutMany2(kvs ...any) {
	if len(kvs)%2 != 0 {
		Warnp("kvs len not 2x", len(kvs))
	}
	me.mu.Lock()
	for i := 0; i < len(kvs)/2; i++ {
		key := kvs[i*2].(KT)
		val := kvs[i*2+1].(VT)
		me.putnolock(key, val)
	}

	me.mu.Unlock()

	return
}

func (me *ListMap0[KT, VT]) Add(key KT, val VT) bool {
	return me.PutTry(key, val)
}
func (me *ListMap0[KT, VT]) PutTry(key KT, val VT) bool {
	hkey := me.hhker.Hash(key)

	var putok bool
	me.mu.Lock()
	_, exist := me.m0[hkey]
	putok = !exist
	if !exist {
		_, putok = me.putnolock(key, val)
	}
	me.mu.Unlock()

	return putok
}

// ///// range/iterate
func (me *ListMap0[KT, VT]) snapshotOrder() (keys []KT, vals []VT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		keys = make([]KT, len(me.a0), len(me.a0))
		vals = make([]VT, len(me.a0), len(me.a0))
		for i, key := range me.a0 {
			kv := me.m0[key]
			keys[i] = kv.Key
			vals[i] = kv.Val
		}
	}
	me.mu.RUnlock()
	return
}

func (me *ListMap0[KT, VT]) EachOrder(fx func(idx int, key KT, val VT) bool) {
	keys, vals := me.snapshotOrder()
	for i := 0; i < len(keys); i++ {
		ok := fx(i, keys[i], vals[i])
		if !ok {
			break
		}
	}
}
func (me *ListMap0[KT, VT]) snapshot() (kvs map[KT]VT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		kvs = make(map[KT]VT, len(me.a0))
		for _, kv := range me.m0 {
			kvs[kv.Key] = kv.Val
		}
	}
	me.mu.RUnlock()
	return
}

func (me *ListMap0[KT, VT]) Each(fx func(idx int, key KT, val VT) bool) {
	kvs := me.snapshot()
	for k, v := range kvs {
		ok := fx(0, k, v)
		if !ok {
			break
		}
	}
}

/////// array method

func (me *ListMap0[KT, VT]) GetIndex(idx int) (rk KT, rv VT, exist bool) {
	if idx < 0 || idx >= len(me.a0) {
		return
	}
	me.mu.RLock()
	rk, rv, exist = me.indexatnolock(idx)
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) indexatnolock(idx int) (rk KT, rv VT, exist bool) {
	if idx < 0 || idx >= len(me.a0) {
		return
	}

	hkey := me.a0[idx]
	kv := me.m0[hkey]
	rk = kv.Key
	rv = kv.Val
	exist = len(me.a0) == len(me.m0)

	return
}

func (me *ListMap0[KT, VT]) DelIndex(idx int) (rv VT, exist bool) {
	// Assert(idx >= 0, "idx must >= 0")
	if idx < 0 || idx >= len(me.a0) {
		return
	}
	me.mu.Lock()
	key := me.a0[idx]
	me.a0 = slices.Delete(me.a0, idx, idx+1)
	delete(me.m0, key)
	exist = len(me.a0) == len(me.m0)
	me.mu.Unlock()

	return
}
func (me *ListMap0[KT, VT]) DelIndexN(idx int, n int) (rv *ListMap0[KT, VT]) {
	Assert(idx >= 0, "idx must >= 0")
	Assert(n >= 0, "n must >= 0")
	rv = ListMap0New[KT, VT]()
	me.mu.Lock()
	eidx := idx + n
	eidx = IfElse2(eidx > len(me.a0)-1, len(me.a0)-1, eidx)
	for i := idx; i < eidx; i++ {
		hkey := me.a0[i]
		kv := me.m0[hkey]
		delete(me.m0, hkey)
		rv.Put(kv.Key, kv.Val)

		if me.reversemap {
			hval := me.hhver.Hash(kv.Val)
			delete(me.mr, hval)
		}
	}
	me.a0 = slices.Delete(me.a0, idx, eidx)
	me.mu.Unlock()

	if rv.Count() == 0 {
		rv = nil
	}

	return
}
func (me *ListMap0[KT, VT]) DelIndexN2(idx int, n int) {
	Assert(idx >= 0, "idx must >= 0")
	Assert(n >= 0, "n must >= 0")
	me.mu.Lock()
	eidx := idx + n
	eidx = IfElse2(eidx > len(me.a0)-1, len(me.a0)-1, eidx)
	for i := idx; i < eidx; i++ {
		hkey := me.a0[i]
		kv := me.m0[hkey]
		delete(me.m0, hkey)

		if me.reversemap {
			hval := me.hhver.Hash(kv.Val)
			delete(me.mr, hval)
		}
	}
	me.a0 = slices.Delete(me.a0, idx, eidx)
	me.mu.Unlock()

	return
}

// return false if exist
func (me *ListMap0[KT, VT]) InsertAt(idx int, key KT, val VT) (ok bool) {
	hkey := me.hhker.Hash(key)
	me.mu.Lock()
	_, exist := me.m0[hkey]
	ok = !exist
	if !exist {
		kv := PairNew(key, val)
		me.a0 = slices.Insert(me.a0, idx, hkey)
		me.m0[hkey] = kv

		if me.reversemap {
			hval := me.hhver.Hash(val)
			me.mr[hval] = kv
		}
	}
	me.mu.Unlock()
	return
}

// there this no append because Put/Add is append already
// return false if exist
// \see InsertAt
func (me *ListMap0[KT, VT]) Prepend(idx int, key KT, val VT) (ok bool) {
	return me.InsertAt(0, key, val)
}

func (me *ListMap0[KT, VT]) firstnolock() (key KT, val VT, exist bool) {
	if len(me.a0) > 0 {
		exist = true
		hkey := me.a0[0]
		kv, ok := me.m0[hkey]
		exist = ok
		key, val = kv.Key, kv.Val
	}

	return
}
func (me *ListMap0[KT, VT]) FirstPair() (key KT, val VT, exist bool) {
	me.mu.RLock()
	key, val, exist = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) FirstKey() (key KT, exist bool) {
	me.mu.RLock()
	key, _, exist = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) FirstKeyMust() (key KT) {
	me.mu.RLock()
	key, _, _ = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) First() (val VT, exist bool) {
	me.mu.RLock()
	_, val, exist = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) FirstMust() (rv VT) {
	me.mu.RLock()
	_, rv, _ = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) FirstN(n int) (rv *ListMap0[KT, VT]) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) FirstKeysN(n int) (rv *ListMap0[KT, VT]) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}

func (me *ListMap0[KT, VT]) lastnolock() (key KT, val VT, exist bool) {
	if len(me.a0) > 0 {
		exist = true
		hkey := me.a0[len(me.a0)-1]
		kv, ok := me.m0[hkey]
		exist = ok
		key, val = kv.Key, kv.Val
	}

	return
}
func (me *ListMap0[KT, VT]) LastPair() (key KT, val VT, exist bool) {
	me.mu.RLock()
	key, val, exist = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) LastKey() (key KT, exist bool) {
	me.mu.RLock()
	key, _, exist = me.lastnolock()
	me.mu.RUnlock()

	return
}

// todo
func (me *ListMap0[KT, VT]) LastKeysN(n int) (keys []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) LastKeyMust() (key KT) {
	me.mu.RLock()
	key, _, _ = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) Last() (rv VT, exist bool) {
	me.mu.RLock()
	_, rv, exist = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) LastMust() (rv VT) {
	me.mu.RLock()
	_, rv, _ = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) LastN(n int) (rv *ListMap0[KT, VT]) {
	me.mu.RLock()
	me.mu.RUnlock()

	return
}

func (me *ListMap0[KT, VT]) Mid(idx int, n int) (rv *ListMap0[KT, VT]) {
	me.mu.RLock()
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) KeysOrder() (keys []KT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		keys = make([]KT, len(me.a0), len(me.a0))
		for i, hkey := range me.a0 {
			kv, _ := me.m0[hkey]
			keys[i] = kv.Key
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) ValuesOrder() (vals []VT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		vals = make([]VT, len(me.a0), len(me.a0))
		for i, hkey := range me.a0 {
			kv := me.m0[hkey]
			vals[i] = kv.Val
		}
	}
	me.mu.RUnlock()

	return
}

///////  map method

func (me *ListMap0[KT, VT]) Keys() (keys []KT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		keys = make([]KT, 0, len(me.a0))
		for _, kv := range me.m0 {
			keys = append(keys, kv.Key)
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) Values() (vals []VT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		vals = make([]VT, 0, len(me.a0))
		for _, kv := range me.m0 {
			vals = append(vals, kv.Val)
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) Get(key KT) (rv VT, exist bool) {
	hkey := me.hhker.Hash(key)

	me.mu.RLock()
	kv, ok := me.m0[hkey]
	exist = ok
	if ok {
		rv = kv.Val
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) GetMust(key KT) (rv VT) {
	hkey := me.hhker.Hash(key)

	me.mu.RLock()
	kv, ok := me.m0[hkey]
	if ok {
		rv = kv.Val
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) GetMany(keys ...KT) (rv *ListMap0[KT, VT]) {
	rv = ListMap0New[KT, VT]()
	me.mu.RLock()
	for _, key := range keys {
		hkey := me.hhker.Hash(key)
		kv, ok := me.m0[hkey]
		if ok {
			rv.putnolock(key, kv.Val)
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) GetMany2(keys ...KT) (rv map[KT]VT) {
	rv = map[KT]VT{}
	me.mu.RLock()
	for _, key := range keys {
		hkey := me.hhker.Hash(key)
		kv, ok := me.m0[hkey]
		if ok {
			rv[kv.Key] = kv.Val
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) Has(key KT) (exist bool) {
	hkey := me.hhker.Hash(key)

	me.mu.RLock()
	_, exist = me.m0[hkey]
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) HasMany(keys ...KT) (rv map[KT]bool) {
	me.mu.RLock()
	rv = make(map[KT]bool, len(keys))
	for _, key := range keys {
		hkey := me.hhker.Hash(key)
		_, exist := me.m0[hkey]
		rv[key] = exist
	}
	me.mu.RUnlock()

	return
}

func (me *ListMap0[KT, VT]) Del(key KT) (exist bool) {
	me.mu.Lock()
	exist = me.delnolock(key)
	me.mu.Unlock()

	return
}
func (me *ListMap0[KT, VT]) delnolock(key KT) (exist bool) {
	hkey := me.hhker.Hash(key)
	exist = me.delnolock2(hkey)
	return
}

// del 不如链表版本
func (me *ListMap0[KT, VT]) delnolock2(hkey uint64) (exist bool) {

	kv, ok := me.m0[hkey]
	exist = ok
	if ok {
		idx := slices.Index(me.a0, hkey)
		me.a0 = slices.Delete(me.a0, idx, idx+1)
		delete(me.m0, hkey)

		if me.reversemap {
			hval := me.hhver.Hash(kv.Val)
			delete(me.mr, hval)
		}
	}

	return
}
func (me *ListMap0[KT, VT]) DelMany(keys ...KT) (rv int) {
	me.mu.Lock()
	for _, k := range keys {
		exist := me.delnolock(k)
		rv += IfElseInt(exist, 1, 0)
	}
	me.mu.Unlock()

	return
}

// /// binsearch, 找到插入位置,保持有序
// 返回该元素的可插入位置
func (me *ListMap0[KT, VT]) BinFind(v VT, cmpfn func(v0 VT, v1 VT) int) (inspos int) {
	me.mu.RLock()
	inspos = me.binfindnolock(v, cmpfn)
	me.mu.RUnlock()
	return
}
func (me *ListMap0[KT, VT]) binfindnolock(v VT, cmpfn func(v0 VT, v1 VT) int) (inspos int) {
	inspos = -1
	begin := 0
	end := me.lennolock()

	var iter = 0
	for iter = 0; iter < 999; iter++ {
		pos := (begin + end) / 2

		_, v0, e0 := me.indexatnolock(pos - 1)
		_, v1, e1 := me.indexatnolock(pos)

		// permit v0 or v1 nil now
		if any(v1) == nil {
			// log.Println(me.lennolock(), begin, end, pos)
		}
		var cmpval0 int
		var cmpval1 int
		if !e0 {
			cmpval0 = 1
		} else {
			cmpval0 = cmpfn(v, v0)
		}
		if !e1 {
			cmpval1 = -1
		} else {
			cmpval1 = cmpfn(v, v1)
		}
		if cmpval0 < 0 && cmpval1 < 0 {
			end = pos
		} else if cmpval0 > 0 && cmpval1 > 0 {
			begin = pos + 1
		} else if cmpval0 >= 0 && cmpval1 <= 0 {
			inspos = pos
			goto endfor
		} else {
			Warn("Invalid cmpval", cmpval0, cmpval1)
		}

		// Debug("i", iter, "pos", pos, "cnt", me.lennolock(), "bgn", begin, "end", end, v0, "v1:", v1)

	}
endfor:
	TruePrint(inspos == -1, "i", iter, "cnt", me.lennolock(), "bgn", begin, "end", end)
	return
}

// ///// reverse operation
func (me *ListMap0[KT, VT]) Getr(val VT) (key KT, exist bool) {
	hval := me.hhver.Hash(val)

	me.mu.RLock()
	kv, ok := me.mr[hval]
	exist = ok
	if ok {
		key = kv.Key
	}
	me.mu.RUnlock()
	return
}
func (me *ListMap0[KT, VT]) GetrMany(vals ...VT) (keys []KT) {

	me.mu.RLock()
	for _, val := range vals {
		hval := me.hhver.Hash(val)
		kv, ok := me.mr[hval]
		if ok {
			keys = append(keys, kv.Key)
		}
	}
	me.mu.RUnlock()
	return
}
func (me *ListMap0[KT, VT]) Hasr(val VT) (exist bool) {
	hval := me.hhver.Hash(val)

	me.mu.RLock()
	_, exist = me.mr[hval]
	me.mu.RUnlock()
	return
}
func (me *ListMap0[KT, VT]) Delr(val VT) (exist bool) {
	hval := me.hhver.Hash(val)

	me.mu.Lock()
	kv, ok := me.mr[hval]
	exist = ok
	if ok {
		me.delnolock(kv.Key)
	}
	me.mu.Unlock()
	return
}
func (me *ListMap0[KT, VT]) DelrMany(vals ...VT) (rv int) {

	me.mu.Lock()
	for _, val := range vals {
		hval := me.hhver.Hash(val)
		kv, ok := me.mr[hval]
		if ok {
			rv += Toint(me.delnolock(kv.Key))
		}
	}
	me.mu.Unlock()
	return
}

////////

// //// other

func (me *ListMap0[KT, VT]) RandKey() (rv KT) {
	me.mu.RLock()
	for _, kv := range me.m0 {
		rv = kv.Key
		break
	}
	me.mu.RUnlock()
	return
}

// todo
func (me *ListMap0[KT, VT]) RandKeys(n int) (rv []KT) {
	me.mu.RLock()
	me.mu.RUnlock()
	return
}
func (me *ListMap0[KT, VT]) RandVal() (rv VT) {
	me.mu.RLock()
	for _, kv := range me.m0 {
		rv = kv.Val
		break
	}
	me.mu.RUnlock()
	return
}

// todo
func (me *ListMap0[KT, VT]) RandVals(n int) (rv []VT) {
	me.mu.RLock()
	me.mu.RUnlock()
	return
}

func (me *ListMap0[KT, VT]) Notin(keys ...KT) (rv []KT) {
	me.mu.RLock()
	for _, key := range keys {
		hkey := me.hhker.Hash(key)
		_, exist := me.m0[hkey]
		if !exist {
			rv = append(rv, key)
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap0[KT, VT]) In(keys ...KT) (rv []KT) {
	me.mu.RLock()
	for _, key := range keys {
		hkey := me.hhker.Hash(key)
		_, exist := me.m0[hkey]
		if exist {
			rv = append(rv, key)
		}
	}
	me.mu.RUnlock()

	return
}

// todo
// not change exist one
func (me *ListMap0[KT, VT]) Merge(m *ListMap0[KT, VT]) (rv []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}

// todo
// change exist one
func (me *ListMap0[KT, VT]) Update(m *ListMap0[KT, VT]) (rv []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}

// todo
func (me *ListMap0[KT, VT]) Diff(m *ListMap0[KT, VT]) (rv []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}
