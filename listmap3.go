package gopp

import (
	"log"
	"sync"
	_ "unsafe"

	"github.com/dolthub/maphash"
	// _ "github.com/elliotchance/orderedmap/v2"
)

// hash 任意类型，来自go内部实现
// https://github.com/dolthub/maphash

const ListMap_SHARD_COUNT = 32

// ordered map, thread safe, reversemap
// 实现从头部删除，或者尾部删除
// 按照随机index删除速度就慢了
type ListMap[KT comparable, VT comparable] struct {
	mu sync.RWMutex
	m0 *OrderedMap[KT, VT]
	mr map[uint64]*Element[KT, VT] // hash(val)=>

	reversemap bool
	lockless   bool // todo

	// hhker maphash.Hasher[KT] // todo
	hhver maphash.Hasher[VT]
}

// type Stringer interface {
// 	fmt.Stringer
// 	comparable
// }

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
	me.m0 = NewOrderedMap[KT, VT]()
	me.mr = map[uint64]*Element[KT, VT]{}

	// me.hhker = maphash.Hasher[KT](maphash.NewHasher[KT]())
	me.hhver = maphash.Hasher[VT](maphash.NewHasher[VT]())
	return me
}

// /////////
func ListMapFrom[KT comparable, VT comparable](m map[KT]VT) *ListMap[KT, VT] {
	me := ListMapNew[KT, VT]()
	me.putmanynolock(m)
	return me
}

/////////

func (me *ListMap[KT, VT]) Count() int {
	return me.m0.Len()
}
func (me *ListMap[KT, VT]) Len() int {
	return me.m0.Len()
}

func (me *ListMap[KT, VT]) putnolock(key KT, val VT) (exist, ok bool) {

	elem := me.m0.GetElement(key)
	exist = elem != nil
	ok = me.m0.Set(key, val)

	if me.reversemap {
		elem := me.m0.GetElement(key)

		hval := me.hhver.Hash(val)
		me.mr[hval] = elem
	}

	return
}

func (me *ListMap[KT, VT]) Put(key KT, val VT) {
	me.mu.Lock()
	exist, ok := me.putnolock(key, val)
	_, _ = exist, ok
	me.mu.Unlock()

	return
}
func (me *ListMap[KT, VT]) PutMany(kvs map[KT]VT) {
	me.mu.Lock()
	me.putmanynolock(kvs)
	me.mu.Unlock()

	return
}
func (me *ListMap[KT, VT]) putmanynolock(kvs map[KT]VT) {
	for k, v := range kvs {
		me.putnolock(k, v)
	}
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

func (me *ListMap[KT, VT]) Add(key KT, val VT) bool {
	return me.PutTry(key, val)
}
func (me *ListMap[KT, VT]) PutTry(key KT, val VT) bool {
	var putok = false
	me.mu.Lock()
	_, ok := me.m0.Get(key)
	if ok {
		putok = false
	} else {
		_, putok = me.putnolock(key, val)
	}
	me.mu.Unlock()
	return putok
}

// ///// range/iterate
func (me *ListMap[KT, VT]) snapshotOrder() (keys []KT, vals []VT) {
	me.mu.RLock()
	if me.m0.Len() > 0 {
		keys = make([]KT, len(keys), len(keys))
		vals = make([]VT, len(keys), len(keys))
		var i = 0
		for el := me.m0.Front(); el != nil; el = el.Next() {
			keys[i] = el.Key
			vals[i] = el.Value
			i++
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
	if me.m0.Len() > 0 {
		kvs = make(map[KT]VT, me.m0.Len())
		for el := me.m0.Front(); el != nil; el = el.Next() {
			kvs[el.Key] = el.Value
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

func (me *ListMap[KT, VT]) elematnolock(idx int) (elem *Element[KT, VT]) {
	var i = 0
	for el := me.m0.Front(); el != nil; el = el.Next() {
		if i == idx {
			return el
		}
		i++
	}
	return nil
}

// O(log(n)), 少用
func (me *ListMap[KT, VT]) GetIndex(idx int) (rk KT, rv VT, exist bool) {
	if idx < 0 || idx >= me.m0.Len() {
		return
	}
	me.mu.RLock()

	el := me.elematnolock(idx)
	if el != nil {
		rk = el.Key
		rv = el.Value
	}
	exist = el != nil

	me.mu.RUnlock()

	return
}

// O(log(n)), 少用
func (me *ListMap[KT, VT]) DelIndex(idx int) (rv VT, exist bool) {
	// Assert(idx >= 0, "idx must >= 0")
	if idx < 0 || idx >= me.m0.Len() {
		return
	}
	me.mu.Lock()
	el := me.elematnolock(idx)
	if el != nil {
		me.delnolock(el.Key)
	}
	exist = el != nil
	me.mu.Unlock()

	return
}
func (me *ListMap[KT, VT]) DelIndexN(idx int, n int) (rv *ListMap[KT, VT]) {
	Assert(idx >= 0, "idx must >= 0")
	Assert(n >= 0, "n must >= 0")
	rv = ListMapNew[KT, VT]()

	var elems = make([]*Element[KT, VT], n, n)
	me.mu.Lock()

	eidx := idx + n
	eidx = IfElse2(eidx > me.m0.Len()-1, me.m0.Len()-1, eidx)
	var i = me.m0.Len() - 1
	var j = 0
	for el := me.m0.Back(); el != nil; el = el.Prev() {
		if i >= idx && i <= eidx {
			elems[j] = el
		}
		i++
		j++
	}
	for _, el := range elems {
		rv.putnolock(el.Key, el.Value)
		me.delnolock(el.Key)
	}

	me.mu.Unlock()

	if rv.Count() == 0 {
		rv = nil
	}

	return
}
func (me *ListMap[KT, VT]) DelIndexN2(idx int, n int) {
	me.DelIndexN(idx, n)
	return
}

// return false if exist
func (me *ListMap[KT, VT]) _InsertAt(idx int, key KT, val VT) (ok bool) {
	panic("not support")
	// return
}

// there this no append because Put/Add is append already
// return false if exist
// \see InsertAt
func (me *ListMap[KT, VT]) Prepend(key KT, val VT) (ok bool) {

	me.mu.Lock()
	ok = me.m0.Prepend(key, val)
	if me.reversemap {
		hval := me.hhver.Hash(val)
		me.mr[hval] = me.m0.GetElement(key)
	}
	me.mu.Unlock()
	return
}

func (me *ListMap[KT, VT]) firstnolock() (key KT, val VT, exist bool) {
	if me.m0.Len() > 0 {
		exist = true
		el := me.m0.Front()
		key = el.Key
		val = el.Value
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
	if me.m0.Len() > 0 {
		el := me.m0.Back()
		exist = true
		key, val = el.Key, el.Value
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
	keys = me.m0.Keys()
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) ValuesOrder() (vals []VT) {
	me.mu.RLock()
	if me.m0.Len() > 0 {
		vals = make([]VT, me.m0.Len(), me.m0.Len())
		var i = 0
		for el := me.m0.Front(); el != nil; el = el.Next() {
			vals[i] = el.Value
			i++
		}
	}
	me.mu.RUnlock()

	return
}

///////  map method

func (me *ListMap[KT, VT]) Keys() (keys []KT) {
	return me.KeysOrder()
}
func (me *ListMap[KT, VT]) Values() (vals []VT) {
	return me.ValuesOrder()
}
func (me *ListMap[KT, VT]) Get(key KT) (rv VT, exist bool) {

	me.mu.RLock()
	rv, exist = me.m0.Get(key)
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetOr(key KT, dftval VT) (rv VT) {

	me.mu.RLock()
	rv = me.m0.GetOrDefault(key, dftval)
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetMust(key KT) (rv VT) {
	me.mu.RLock()
	rv2, ok := me.m0.Get(key)
	if ok {
		rv = rv2
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetMany(keys ...KT) (rv *ListMap[KT, VT]) {
	rv = ListMapNew[KT, VT]()

	me.mu.RLock()
	for _, key := range keys {
		val, ok := me.m0.Get(key)
		if ok {
			rv.putnolock(key, val)
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) GetMany2(keys ...KT) (rv map[KT]VT) {
	rv = map[KT]VT{}

	me.mu.RLock()
	for _, key := range keys {
		val, ok := me.m0.Get(key)
		if ok {
			rv[key] = val
		}
	}
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) Has(key KT) (exist bool) {

	me.mu.RLock()
	_, exist = me.m0.Get(key)
	me.mu.RUnlock()

	return
}
func (me *ListMap[KT, VT]) HasMany(keys ...KT) (rv map[KT]bool) {
	me.mu.RLock()
	rv = make(map[KT]bool, len(keys))
	for _, key := range keys {
		_, exist := me.m0.Get(key)
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

	if me.reversemap {
		var val VT
		val, exist = me.m0.Get(key)
		hval := me.hhver.Hash(val)
		delete(me.mr, hval)
	}
	exist = me.m0.Delete(key)

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
func (me *ListMap[KT, VT]) DelFirst() (ok bool) {
	me.mu.Lock()
	elem := me.m0.Front()
	ok = elem != nil
	if ok {
		ok = me.delnolock(elem.Key)
	}
	me.mu.Unlock()
	return
}
func (me *ListMap[KT, VT]) DelFirstN(n int) (rvlm *ListMap[KT, VT]) {
	rvlm = ListMapNew[KT, VT]()
	var elems = make([]*Element[KT, VT], n, n)
	me.mu.Lock()
	for i, el := 0, me.m0.Front(); i < n && el != nil; i, el = i+1, el.Next() {
		elems = append(elems, el)
	}

	for _, el := range elems {
		me.delnolock(el.Key)

		rvlm.putnolock(el.Key, el.Value)
	}
	me.mu.Unlock()
	return
}
func (me *ListMap[KT, VT]) DelLast() (ok bool) {
	me.mu.Lock()
	elem := me.m0.Back()
	ok = elem != nil
	if ok {
		ok = me.delnolock(elem.Key)
	}
	me.mu.Unlock()
	return
}
func (me *ListMap[KT, VT]) DelLastN(n int) (rvlm *ListMap[KT, VT]) {
	rvlm = ListMapNew[KT, VT]()
	var elems = make([]*Element[KT, VT], n, n)
	me.mu.Lock()
	for i, el := 0, me.m0.Back(); i < n && el != nil; i, el = i+1, el.Prev() {
		elems = append(elems, el)
	}
	for _, el := range elems {
		me.delnolock(el.Key)

		rvlm.putnolock(el.Key, el.Value)
	}
	me.mu.Unlock()
	return
}

// ///// reverse operation
func (me *ListMap[KT, VT]) checkreversemap() bool {
	if !me.reversemap {
		Warnp("reversemap not enabled")
	}
	return me.reversemap
}
func (me *ListMap[KT, VT]) Getr(val VT) (key KT, exist bool) {
	if !me.checkreversemap() {
		return
	}

	hval := me.hhver.Hash(val)
	me.mu.RLock()
	elem, exist := me.mr[hval]
	if exist {
		key = elem.Key
	}
	me.mu.RUnlock()
	return
}
func (me *ListMap[KT, VT]) GetrMany(vals ...VT) (keys []KT) {
	if !me.checkreversemap() {
		return
	}

	me.mu.RLock()
	for _, val := range vals {
		hval := me.hhver.Hash(val)
		elem, ok := me.mr[hval]
		if ok {
			keys = append(keys, elem.Key)
		}
	}
	me.mu.RUnlock()
	return
}
func (me *ListMap[KT, VT]) Hasr(val VT) (exist bool) {
	if !me.checkreversemap() {
		return
	}

	hval := me.hhver.Hash(val)
	me.mu.RLock()
	_, exist = me.mr[hval]
	me.mu.RUnlock()
	return
}
func (me *ListMap[KT, VT]) Delr(val VT) (exist bool) {
	if !me.checkreversemap() {
		return
	}

	hval := me.hhver.Hash(val)

	me.mu.Lock()
	elem, exist := me.mr[hval]
	if exist {
		me.delnolock(elem.Key)
	}
	me.mu.Unlock()
	return
}
func (me *ListMap[KT, VT]) DelrMany(vals ...VT) (rv int) {
	if !me.checkreversemap() {
		return
	}

	me.mu.Lock()
	for _, val := range vals {
		hval := me.hhver.Hash(val)
		elem, ok := me.mr[hval]
		if ok {
			rv += Toint(me.delnolock(elem.Key))
		}
	}
	me.mu.Unlock()
	return
}

////////

// //// other

func (me *ListMap[KT, VT]) RandKey() (rv KT) {

	me.mu.RLock()
	for key, _ := range me.m0.kv {
		rv = key
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
	for _, elem := range me.m0.kv {
		rv = elem.Value
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

// more more more ...

func (me *ListMap[KT, VT]) Notin(keys ...KT) (rv []KT) {
	panic("not support")
	// return
}
func (me *ListMap[KT, VT]) In(keys ...KT) (rv []KT) {
	panic("not support")
	// return
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
