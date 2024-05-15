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
	m0 map[KT]VT
	a0 []KT      // todo how know map key hash value
	mr map[VT]KT //
	// mr2 map[VT][]KT // TODO, reverse map conflict

	reversemap bool

	hher maphash.Hasher[KT] // todo
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
	me.m0 = make(map[KT]VT, 8)
	me.mr = make(map[VT]KT, 8)
	me.a0 = make([]KT, 0, 8)

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

	_, exist := me.m0[key]
	if exist {
		idx := slices.Index(me.a0, key)
		me.a0 = slices.Delete(me.a0, idx, idx+1)
	}
	me.m0[key] = val
	me.a0 = append(me.a0, key)

	if me.reversemap {
		me.mr[val] = key
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
	var putok bool
	me.mu.Lock()
	_, exist := me.m0[key]
	if !exist {
		me.m0[key] = val
		me.a0 = append(me.a0, key)

		if me.reversemap {
			me.mr[val] = key
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
			keys[i] = key
			vals[i] = me.m0[key]
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
		for _, key := range me.a0 {
			kvs[key] = me.m0[key]
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
	rk = me.a0[idx]
	rv = me.m0[rk]
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
		key := me.a0[i]
		val := me.m0[key]
		delete(me.m0, key)
		rv.Put(key, val)

		if me.reversemap {
			delete(me.mr, val)
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
		key := me.a0[i]
		val := me.m0[key]
		delete(me.m0, key)

		if me.reversemap {
			delete(me.mr, val)
		}
	}
	me.a0 = slices.Delete(me.a0, idx, eidx)
	me.mu.Unlock()

	return
}

func (me *ListMap[KT, VT]) firstnolock() (key KT, val VT, exist bool) {
	if len(me.a0) > 0 {
		exist = true
		key = me.a0[0]
		val, exist = me.m0[key]
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
		key := me.a0[len(me.a0)-1]
		val, exist = me.m0[key]
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
		for i, key := range me.a0 {
			keys[i] = key
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) ValuesOrder() (vals []VT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		vals = make([]VT, len(me.a0), len(me.a0))
		for i, key := range me.a0 {
			val := me.m0[key]
			vals[i] = val
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
		for key, _ := range me.m0 {
			keys = append(keys, key)
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) Values() (vals []VT) {
	me.mu.RLock()
	if len(me.a0) > 0 {
		vals = make([]VT, 0, len(me.a0))
		for _, val := range me.m0 {
			vals = append(vals, val)
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) Get(key KT) (rv VT, exist bool) {
	me.mu.RLock()
	rv, exist = me.m0[key]
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetMust(key KT) (rv VT) {
	me.mu.RLock()
	rv, _ = me.m0[key]
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
	me.mu.RLock()
	_, exist = me.m0[key]
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) HasMany(keys ...KT) (rv map[KT]bool) {
	me.mu.RLock()
	rv = make(map[KT]bool, len(keys))
	for _, key := range keys {
		_, exist := me.m0[key]
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

	_, exist = me.m0[key]
	if exist {
		idx := slices.Index(me.a0, key)
		me.a0 = slices.Delete(me.a0, idx, idx+1)
		val := me.m0[key]
		delete(me.m0, key)

		if me.reversemap {
			delete(me.mr, val)
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
	me.mu.RLock()
	key, exist = me.mr[val]
	me.mu.RUnlock()
	return
}
func (me *ListMap[KT, VT]) Hasr(val VT) (exist bool) {
	me.mu.RLock()
	_, exist = me.mr[val]
	me.mu.RUnlock()
	return
}
func (me *ListMap[KT, VT]) Delr(val VT) (exist bool) {
	me.mu.RLock()
	key, exist := me.mr[val]
	if exist {
		me.delnolock(key)
	}
	me.mu.RUnlock()
	return
}

////////

// //// other

func (me *ListMap[KT, VT]) RandKey() (rv KT) {
	me.mu.RLock()
	for rv, _ = range me.m0 {
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
	for _, rv = range me.m0 {
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
		_, exist := me.m0[key]
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
		_, exist := me.m0[key]
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
