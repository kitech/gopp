package gopp

import (
	"fmt"
	"log"
	"slices"
	"sync"

	"github.com/dolthub/maphash"
)

// hash 任意类型，来自go内部实现
// https://github.com/dolthub/maphash

const LISTMAP_SHARD_COUNT = 32

// ordered map, thread safe, reversemap
type ListMap[KT comparable, VT comparable] struct {
	mu sync.RWMutex
	m0 map[uint64]*Pair[KT, VT]
	a0 []uint64                 // todo how know map key hash value
	mr map[uint64]*Pair[KT, VT] // hash(val) => Pair
	// mr2 map[VT][]KT // TODO, reverse map conflict

	reversemap bool

	hhker maphash.Hasher[KT] // todo
	hhver maphash.Hasher[VT]
}

type Stringer interface {
	fmt.Stringer
	comparable
}

func ListMapNewAny[VT comparable]() *ListMap[any, VT] {
	return ListMapNew[any, VT]()
}
func ListMapNewInt[VT comparable]() *ListMap[int, VT] {
	return ListMapNew[int, VT]()
}
func ListMapNewStr[VT comparable]() *ListMap[string, VT] {
	return ListMapNew[string, VT]()
}
func ListMapNewr[KT comparable, VT comparable]() *ListMap[KT, VT] {
	me := ListMapNew[KT, VT]()
	me.reversemap = true

	return me
}

func ListMapNew[KT comparable, VT comparable]() *ListMap[KT, VT] {
	me := &ListMap[KT, VT]{}
	me.m0 = make(map[uint64]*Pair[KT, VT], 8)
	me.mr = make(map[uint64]*Pair[KT, VT], 8)
	me.a0 = make([]uint64, 0, 8)
	me.hhker = maphash.Hasher[KT](maphash.NewHasher[KT]())
	me.hhver = maphash.Hasher[VT](maphash.NewHasher[VT]())

	return me
}

// /////////
func ListMapFrom[KT comparable, VT comparable](m map[KT]VT) *ListMap[KT, VT] {
	me := ListMapNew[KT, VT]()
	me.PutMany(m)
	return me
}

/////////

func (me *ListMap[KT, VT]) Count() int {
	return len(me.a0)
}

func (me *ListMap[KT, VT]) putnolock(key KT, val VT) {
	hkey := me.hhker.Hash(key)

	_, exist := me.m0[hkey]
	if exist {
		idx := slices.Index(me.a0, hkey)
		me.a0 = slices.Delete(me.a0, idx, idx+1)
	}
	kv := PairNew(key, val)
	me.m0[hkey] = kv
	me.a0 = append(me.a0, hkey)

	if me.reversemap {
		hval := me.hhver.Hash(val)
		me.mr[hval] = kv
	}

	return
}
func (me *ListMap[KT, VT]) Put(key KT, val VT) {
	me.mu.Lock()
	me.putnolock(key, val)
	me.mu.Unlock()

	return
}
func (me *ListMap[KT, VT]) PutMany(kvs map[KT]VT) {
	me.mu.Lock()
	for k, v := range kvs {
		me.putnolock(k, v)
	}
	me.mu.Unlock()

	return
}
func (me *ListMap[KT, VT]) PutMany3(keys []KT, vals []VT) {
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
func (me *ListMap[KT, VT]) PutMany2(kvs ...any) {
	if len(kvs)%2 != 0 {
		log.Println("kvs len not 2x", len(kvs))
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

func (me *ListMap[KT, VT]) PutTry(key KT, val VT) bool {
	hkey := me.hhker.Hash(key)

	var putok bool
	me.mu.Lock()
	_, exist := me.m0[hkey]
	if !exist {
		kv := PairNew(key, val)
		me.m0[hkey] = kv
		me.a0 = append(me.a0, hkey)

		if me.reversemap {
			hval := me.hhver.Hash(val)
			me.mr[hval] = kv
		}
	}
	putok = !exist
	me.mu.Unlock()

	return putok
}

// ///// range/iterate
func (me *ListMap[KT, VT]) snapshotOrder() (keys []KT, vals []VT) {
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

func (me *ListMap[KT, VT]) EachOrder(fx func(idx int, key KT, val VT) bool) {
	keys, vals := me.snapshotOrder()
	for i := 0; i < len(keys); i++ {
		ok := fx(i, keys[i], vals[i])
		if !ok {
			break
		}
	}
}
func (me *ListMap[KT, VT]) snapshot() (kvs map[KT]VT) {
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

func (me *ListMap[KT, VT]) Each(fx func(idx int, key KT, val VT) bool) {
	kvs := me.snapshot()
	for k, v := range kvs {
		ok := fx(0, k, v)
		if !ok {
			break
		}
	}
}

/////// array method

func (me *ListMap[KT, VT]) GetIndex(idx int) (rk KT, rv VT, exist bool) {
	if idx < 0 || idx >= len(me.a0) {
		return
	}
	me.mu.RLock()
	hkey := me.a0[idx]
	kv := me.m0[hkey]
	rk = kv.Key
	rv = kv.Val
	exist = len(me.a0) == len(me.m0)
	me.mu.RUnlock()

	return
}

func (me *ListMap[KT, VT]) DelIndex(idx int) (rv VT, exist bool) {
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
func (me *ListMap[KT, VT]) DelIndexN(idx int, n int) (rv *ListMap[KT, VT]) {
	Assert(idx >= 0, "idx must >= 0")
	Assert(n >= 0, "n must >= 0")
	rv = ListMapNew[KT, VT]()
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
func (me *ListMap[KT, VT]) DelIndexN2(idx int, n int) {
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

func (me *ListMap[KT, VT]) firstnolock() (key KT, val VT, exist bool) {
	if len(me.a0) > 0 {
		exist = true
		hkey := me.a0[0]
		kv, ok := me.m0[hkey]
		exist = ok
		key, val = kv.Key, kv.Val
	}

	return
}
func (me *ListMap[KT, VT]) FirstPair() (key KT, val VT, exist bool) {
	me.mu.RLock()
	key, val, exist = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) FirstKey() (key KT, exist bool) {
	me.mu.RLock()
	key, _, exist = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) FirstKeyMust() (key KT) {
	me.mu.RLock()
	key, _, _ = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) First() (val VT, exist bool) {
	me.mu.RLock()
	_, val, exist = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) FirstMust() (rv VT) {
	me.mu.RLock()
	_, rv, _ = me.firstnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) FirstN(n int) (rv *ListMap[KT, VT]) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) FirstKeysN(n int) (rv *ListMap[KT, VT]) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}

func (me *ListMap[KT, VT]) lastnolock() (key KT, val VT, exist bool) {
	if len(me.a0) > 0 {
		exist = true
		hkey := me.a0[len(me.a0)-1]
		kv, ok := me.m0[hkey]
		exist = ok
		key, val = kv.Key, kv.Val
	}

	return
}
func (me *ListMap[KT, VT]) LastPair() (key KT, val VT, exist bool) {
	me.mu.RLock()
	key, val, exist = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) LastKey() (key KT, exist bool) {
	me.mu.RLock()
	key, _, exist = me.lastnolock()
	me.mu.RUnlock()

	return
}

// todo
func (me *ListMap[KT, VT]) LastKeysN(n int) (keys []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) LastKeyMust() (key KT) {
	me.mu.RLock()
	key, _, _ = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) Last() (rv VT, exist bool) {
	me.mu.RLock()
	_, rv, exist = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) LastMust() (rv VT) {
	me.mu.RLock()
	_, rv, _ = me.lastnolock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) LastN(n int) (rv *ListMap[KT, VT]) {
	me.mu.RLock()
	me.mu.RUnlock()

	return
}

func (me *ListMap[KT, VT]) Mid(idx int, n int) (rv *ListMap[KT, VT]) {
	me.mu.RLock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) KeysOrder() (keys []KT) {
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
func (me *ListMap[KT, VT]) ValuesOrder() (vals []VT) {
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

func (me *ListMap[KT, VT]) Keys() (keys []KT) {
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
func (me *ListMap[KT, VT]) Values() (vals []VT) {
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
func (me *ListMap[KT, VT]) Get(key KT) (rv VT, exist bool) {
	hkey := me.hhker.Hash(key)

	me.mu.RLock()
	kv, ok := me.m0[hkey]
	rv = kv.Val
	exist = ok
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetMust(key KT) (rv VT) {
	hkey := me.hhker.Hash(key)

	me.mu.RLock()
	kv := me.m0[hkey]
	rv = kv.Val
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetMany(key ...KT) (rv *ListMap[KT, VT]) {
	me.mu.RLock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetMany2(key ...KT) (rv map[KT]VT) {
	me.mu.RLock()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) Has(key KT) (exist bool) {
	hkey := me.hhker.Hash(key)

	me.mu.RLock()
	_, exist = me.m0[hkey]
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) HasMany(keys ...KT) (rv map[KT]bool) {
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

func (me *ListMap[KT, VT]) Del(key KT) (exist bool) {
	me.mu.Lock()
	exist = me.delnolock(key)
	me.mu.Unlock()

	return
}
func (me *ListMap[KT, VT]) delnolock(key KT) (exist bool) {
	hkey := me.hhker.Hash(key)
	exist = me.delnolock2(hkey)
	return
}
func (me *ListMap[KT, VT]) delnolock2(hkey uint64) (exist bool) {

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
func (me *ListMap[KT, VT]) DelMany(keys ...KT) (rv int) {
	me.mu.Lock()
	for _, k := range keys {
		exist := me.delnolock(k)
		rv += IfElseInt(exist, 1, 0)
	}
	me.mu.Unlock()

	return
}

// ///// reverse operation
func (me *ListMap[KT, VT]) Getr(val VT) (key KT, exist bool) {
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
func (me *ListMap[KT, VT]) GetrMany(vals ...VT) (keys []KT) {

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
func (me *ListMap[KT, VT]) Hasr(val VT) (exist bool) {
	hval := me.hhver.Hash(val)

	me.mu.RLock()
	_, exist = me.mr[hval]
	me.mu.RUnlock()
	return
}
func (me *ListMap[KT, VT]) Delr(val VT) (exist bool) {
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
func (me *ListMap[KT, VT]) DelrMany(vals ...VT) (rv int) {

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

func (me *ListMap[KT, VT]) RandKey() (rv KT) {
	me.mu.RLock()
	for _, kv := range me.m0 {
		rv = kv.Key
		break
	}
	me.mu.RUnlock()
	return
}

// todo
func (me *ListMap[KT, VT]) RandKeys(n int) (rv []KT) {
	me.mu.RLock()
	me.mu.RUnlock()
	return
}
func (me *ListMap[KT, VT]) RandVal() (rv VT) {
	me.mu.RLock()
	for _, kv := range me.m0 {
		rv = kv.Val
		break
	}
	me.mu.RUnlock()
	return
}

// todo
func (me *ListMap[KT, VT]) RandVals(n int) (rv []VT) {
	me.mu.RLock()
	me.mu.RUnlock()
	return
}

func (me *ListMap[KT, VT]) Notin(keys ...KT) (rv []KT) {
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
func (me *ListMap[KT, VT]) In(keys ...KT) (rv []KT) {
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
func (me *ListMap[KT, VT]) Merge(m *ListMap[KT, VT]) (rv []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}

// todo
// change exist one
func (me *ListMap[KT, VT]) Update(m *ListMap[KT, VT]) (rv []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}

// todo
func (me *ListMap[KT, VT]) Diff(m *ListMap[KT, VT]) (rv []KT) {
	me.mu.RLock()

	me.mu.RUnlock()

	return
}
