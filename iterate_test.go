package gopp

import (
	"reflect"
	"testing"
)

func TestIV2Str(t *testing.T) {
	{
		var vec = []any{"aaa", "bbb", "ccc"}
		ret := IV2Strings(vec)
		if reflect.TypeOf(ret).String() != "[]string" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}

	{
		var vec = []any{111, 222, 333}
		ret := IV2Strings(vec)
		if reflect.TypeOf(ret).String() != "[]string" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}
	{
		var vec = []any{111, 222, 333}
		ret := IVConvert[uint](vec)
		if reflect.TypeOf(ret).String() != "[]uint" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}
	{
		var vec = []any{111, 222, 333}
		ret := IVConvert[uint64](vec)
		if reflect.TypeOf(ret).String() != "[]uint64" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}
	{
		var vec = []any{111, 222, 333}
		ret := IVConvert[uintptr](vec)
		if reflect.TypeOf(ret).String() != "[]uintptr" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}
	{
		var vec = []any{"111", "222", "333"}
		ret := IVConvert[int](vec)
		if reflect.TypeOf(ret).String() != "[]int" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}
	{
		var vec = []any{"111", "222", "333"}
		ret := IVConvert[uint](vec)
		if reflect.TypeOf(ret).String() != "[]uint" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}

	{
		var vec = []any{111.111, 222.222, 333.333}
		ret := IV2Strings(vec)
		if reflect.TypeOf(ret).String() != "[]string" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}

	{
		var vec = []any{"111.111", "222.222", "333.333"}
		ret := IVConvert[float64](vec)
		if reflect.TypeOf(ret).String() != "[]float64" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}
	{
		var vec = []any{"111.111", "222.222", "333.333"}
		ret := IVConvert[float32](vec)
		if reflect.TypeOf(ret).String() != "[]float32" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}

	{
		var vec = []any{"111.111", "222.222", "333.333"}
		ret := IVConvert[int](vec)
		if reflect.TypeOf(ret).String() != "[]int" || len(ret) != len(vec) {
			t.Error(ret, reflect.TypeOf(ret), reflect.TypeOf(vec))
		}
	}

}
