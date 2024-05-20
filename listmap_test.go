package gopp

import (
	mrand "math/rand"
	"slices"
	"sync"
	"testing"
)

func TestListMap0(t *testing.T) {
	{
		lm := ListMapNewInt[string]()
		{
			_, ok := lm.First()
			if ok {
				t.Error("empty cannot ok", ok)
			}
		}
		{
			_, ok := lm.Last()
			if ok {
				t.Error("empty cannot ok", ok)
			}
		}

		lm.Put(3, "s333")
		lm.Put(1, "s111")
		lm.Put(2, "s222")

		keys := lm.KeysOrder()
		etkeys := []int{3, 1, 2}
		if !slices.Equal(keys, etkeys) {
			t.Error("error noteq", keys, etkeys)
		}
		vals := lm.ValuesOrder()
		etvals := []string{"s333", "s111", "s222"}
		if !slices.Equal(vals, etvals) {
			t.Error("error noteq", vals, etvals)
		}

		if lm.Count() != 3 {
			t.Error("count must 3", lm.Count())
		}

		if !lm.Has(3) {
			t.Error("must has 3")
		}
		if lm.Has(5) {
			t.Error("must not has 5")
		}

		if !lm.Del(3) {
			t.Error("del key error", 3)
		}
		if lm.Has(3) {
			t.Error("must not has 3")
		}
		if lm.Count() != 2 {
			t.Error("count must 2", lm.Count())
		}

		if _, ok := lm.DelIndex(30); ok {
			t.Error("must not ok", ok)
		}
		if _, ok := lm.DelIndex(1); !ok {
			t.Error("must ok", ok, lm.Count())
		}
		if lm.Count() != 1 {
			t.Error("count must 1", lm.Count())
		}

		if _, ok := lm.First(); !ok {
			t.Error("must ok", ok)
		}
		if _, ok := lm.Last(); !ok {
			t.Error("must ok", ok)
		}
		if lm.FirstMust() != lm.LastMust() {
			t.Error("now First must eq Last", lm.FirstMust(), lm.LastMust())
		}

		// log.Println(lm.Keys(), lm.Count()) //
		// left key=1 now
		if lm.Hasr("s111") {
			t.Error("must has not s111 now") // not set reverse map
		}
	}
	//
	// log.Println("=======")
	{
		lm := ListMapNewr[int, string]()

		lm.Put(3, "s333")
		lm.Put(1, "s111")
		lm.Put(2, "s222")

		if !lm.Hasr("s111") {
			t.Error("must has s111 now")
		}

		if !lm.Delr("s111") {
			t.Error("must del s111")
		}

		if lm.Hasr("s111") {
			t.Error("must has not s111 now, after delete")
		}
		if lm.Has(1) {
			t.Error("must has not key=1 now, after delete")
		}

		if key, ok := lm.Getr("s222"); !ok || key != 2 {
			t.Error("must has and val=2", ok, key)
		}
		if key, ok := lm.Getr("s999"); ok {
			t.Error("must not has", ok, key)
		}
	}
}

func BenchmarkListMap0(b *testing.B) {
	lm := ListMapNew[int, int]()
	{
		for i := 0; i < b.N; i++ {
			lm.Has(123)
		}
	}
}
func BenchmarkListMap1(b *testing.B) {
	lm := ListMapNew[int, int]()
	{
		for i := 0; i < b.N; i++ {
			lm.Put(mrand.Int()%123456, mrand.Int())
		}
	}
}
func BenchmarkListMap12(b *testing.B) {
	lm := ListMapNew[int, int]()
	{
		for i := 0; i < b.N; i++ {
			lm.Put(mrand.Int()%123456, mrand.Int())
		}
	}
}
func BenchmarkListMap11(b *testing.B) {
	// 这个速度可以，100ns/op
	lm := orderedMapNew[int, int]()
	{
		var mu sync.RWMutex
		for i := 0; i < b.N; i++ {
			mu.Lock()
			lm.Set(mrand.Int()%123456, mrand.Int())
			mu.Unlock()
			// lm.Put(mrand.Int()%123456, mrand.Int())
		}
	}
}

func BenchmarkListMap2(b *testing.B) {
	v := []int{}
	{
		for i := 0; i < b.N; i++ {
			v = append(v, mrand.Int()%123456)
		}
	}
}

func TestListMap3(t *testing.T) {
	lm := ListMapNew[int, int]()
	lm.Put(555, 666)
	lm.Prepend(111, 222)
}
