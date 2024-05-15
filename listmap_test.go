package gopp

import (
	"slices"
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

	}
}
