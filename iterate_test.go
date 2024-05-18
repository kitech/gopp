package gopp

import (
	"reflect"
	"strings"
	"testing"
)

func TestMapdo1(t *testing.T) {
	{
		// map 1=>1
		var vec = []string{"http://x1.com", "http://x2.com", "http://x3.com"}
		var epval = []any{"x1.com", "x2.com", "x3.com"}
		var rv = Mapdo(vec, func(vx any) any {
			var v = vx.(string)
			return v[7:]
		})
		if !reflect.DeepEqual(epval, rv) {
			t.Error("not equal", rv, epval)
		}
		// log.Println(rv)
	}
	{
		// map 1=>n
		var vec = []string{"http://x1.com", "http://x2.com", "http://x3.com"}
		var epval = []any{"x1", "com", "x2", "com", "x3", "com"}
		var rv = Mapdo(vec, func(vx any) []any {
			var v = vx.(string)
			v = v[7:]
			arr := strings.Split(v, ".")
			return []any{arr[0], arr[1]}
		})
		if !reflect.DeepEqual(epval, rv) {
			t.Error("not equal", rv, epval)
		}
		// log.Println(rv)
	}
	{
		// reduce 1
		var vec = []string{"http://x1.com", "http://x2.com", "http://x3.com"}
		var epval = []any{"x1.com", "x3.com"}
		var rv = Mapdo(vec, func(vx any) []any {
			var v = vx.(string)
			v = v[7:]
			if v == "x2.com" {
				return nil
			}
			return []any{v}
		})
		if !reflect.DeepEqual(epval, rv) {
			t.Error("not equal", rv, epval)
		}
		// log.Println(rv)
	}

}

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

// go test -v -bench RangeA -run RangeA

func BenchmarkRangeA(b *testing.B) {
	// for n := 0; n < b.N; n++ {
	// 	fib(30) // run fib(30) b.N times
	// }
	// log.Println(b.N)
	for n := 0; n < b.N; n++ {
		for range _RangeA(33) {
		}
	}
}
func BenchmarkRangeA2(b *testing.B) {
	// for n := 0; n < b.N; n++ {
	// 	fib(30) // run fib(30) b.N times
	// }
	// log.Println(b.N)
	for n := 0; n < b.N; n++ {
		for range RangeA(33) {
		}
	}
}
func BenchmarkRangeC(b *testing.B) {
	// for n := 0; n < b.N; n++ {
	// 	fib(30) // run fib(30) b.N times
	// }
	// log.Println(b.N)
	for n := 0; n < b.N; n++ {
		for range RangeC(33) {
		}
	}
}
