package gopp

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type Pair[KT any, VT any] struct {
	Key KT
	Val VT
}

func PairNew[KT any, VT any](key KT, val VT /*, extra ...interface{}*/) *Pair[KT, VT] {
	p := &Pair[KT, VT]{key, val}
	return p
}
func PairNewInt[VT any](key int, val VT) *Pair[int, VT] {
	return PairNew(key, val)
}
func PairNewU64[VT any](key uint64, val VT) *Pair[uint64, VT] {
	return PairNew(key, val)
}
func PairNewStr[VT any](key string, val VT) *Pair[string, VT] {
	return PairNew(key, val)
}

// redeclare
// func Domap[T any](ins []T) {}

// 可以写成这个样子，但是挺麻烦的
func mapdo2[FT func(int, any, any) map[any]any |
	func(int, any, any) []any |
	func(int, any, any) any](ins any, fx FT) {
}

// 该函数没有办法终止，必须完全循环完成。。。
// 这个函数没法写成范型了吧
// 模板参数无法表达固定长度数组
// 返回 array/slice
// f proto:
//
//	func(any) any
//	func(any) []any
//	func(int,any) any
//	func(int,any) []any
//	func(int,any,any) any // map only
//	func(int,any,any) []any // map only
//
// 返回 maps/hashtable
//
//	func(any) (any,any)
//	func(int,any) (any,any)
//	func(int,any,any) (any,any)
//	func(any) map[any]any
//	func(int,any) map[any]any
//	func(int,any,any) map[any]any
//
// 可以n=>n, n=>n+,n=>n-，具有reduce功能，所以不需要单独的reduce
// 支持可以迭代的类型：结构体，slice，数组，字符串，map
// todo fx 的参数可以不用any的
// todo fx 可以没有返回值的
func Mapdo(ins any, fx any) (out any) {
	infxty := reflect.TypeOf(fx)

	outmap := make(map[any]any, 0)
	outarr := make([]any, 0)
	retmap := false
	ff := func(idx int, key any, val any) bool {
		switch f := fx.(type) {
		case func(any) (any, any):
			v0, v1 := f(val)
			outmap[v0] = v1
			retmap = true
		case func(any) map[any]any:
			vm := f(val)
			for k, v := range vm {
				outmap[k] = v
			}
			retmap = true
		case func(int, any) (any, any):
			v0, v1 := f(idx, val)
			outmap[v0] = v1
		case func(int, any) map[any]any:
			vm := f(idx, val)
			for k, v := range vm {
				outmap[k] = v
			}
			retmap = true
		case func(int, any, any) (any, any):
			v0, v1 := f(idx, key, val)
			outmap[v0] = v1
			retmap = true
		case func(int, any, any) map[any]any:
			vm := f(idx, key, val)
			for k, v := range vm {
				outmap[k] = v
			}
			retmap = true

		case func(any) []any:
			out := f(val)
			outarr = append(outarr, out...)
		case func(any) any:
			out := f(val)
			outarr = append(outarr, out)
		case func(any):
			f(val)

		case func(int, any) []any:
			out := f(idx, val)
			outarr = append(outarr, out...)
		case func(int, any) any:
			out := f(idx, val)
			outarr = append(outarr, out)
		case func(int, any):
			f(idx, val)

		case func(int, any, any) []any:
			out := f(idx, key, val)
			outarr = append(outarr, out...)
		case func(int, any, any) any:
			out := f(idx, key, val)
			outarr = append(outarr, out)
		case func(int, any, any):
			f(idx, key, val)

		default:
			// 处理参数为实际类型的情况
			if infxty.NumIn() == 1 {
				if reflect.TypeOf(val) == infxty.In(0) {
					// ok
					args := Sliceof(reflect.ValueOf(val))
					outsx := reflect.ValueOf(fx).Call(args)
					if outsx != nil {
						// todo, assign return value
					}
					break
				}
			} else if infxty.NumIn() == 2 {
				if reflect.TypeOf(val) == infxty.In(1) &&
					reflect.TypeOf(idx) == infxty.In(0) {
					// ok
					args := Sliceof(reflect.ValueOf(idx), reflect.ValueOf(val))
					outsx := reflect.ValueOf(fx).Call(args)
					if outsx != nil {
						// todo, assign return value
					}
					break
				}
			} else if infxty.NumIn() == 3 {
				if reflect.TypeOf(val) == infxty.In(2) &&
					reflect.TypeOf(key) == infxty.In(1) &&
					reflect.TypeOf(idx) == infxty.In(0) {
					// ok
					args := Sliceof(reflect.ValueOf(idx), reflect.ValueOf(key), reflect.ValueOf(val))
					outsx := reflect.ValueOf(fx).Call(args)
					if outsx != nil {
						// todo, assign return value
					}
					break
				}
			}

			tylst := fmt.Sprintf("(int?, %v?, %v)", reflect.TypeOf(key), reflect.TypeOf(val))
			Infop("invalid fxcb", idx, key, infxty, "want", tylst)
			return false
		}
		return true
	}

	tmpty := reflect.TypeOf(ins)
	switch tmpty.Kind() {
	case reflect.Map:
		tmpv := reflect.ValueOf(ins)
		for idx, vk := range tmpv.MapKeys() {
			vv := tmpv.MapIndex(vk).Interface()
			if !ff(idx, vk.Interface(), vv) {
				break
			}
		}
	case reflect.Slice, reflect.Array:
		tmpv := reflect.ValueOf(ins)
		for idx := 0; idx < tmpv.Len(); idx++ {
			e := tmpv.Index(idx).Interface()
			if !ff(idx, nil, e) {
				break
			}
		}
	case reflect.String:
		for idx, uc := range ins.(string) {
			if !ff(idx, nil, uc) {
				break
			}
		}
	case reflect.Struct:
		tmpv := reflect.ValueOf(ins)
		for idx := 0; idx < tmpv.NumField(); idx++ {
			key := tmpty.Field(idx).Name
			val := tmpv.Field(idx).Interface()
			if !ff(idx, key, val) {
				break
			}
		}
	case reflect.Interface:
		tmpv := reflect.ValueOf(ins)
		for idx := 0; idx < tmpv.NumMethod(); idx++ {
			key := tmpty.Method(idx).Name
			val := tmpty.Method(idx).Type.String()
			if !ff(idx, key, val) {
				break
			}
		}
	default:
		// the same as DomapTypeField
		if tmpty.Kind() == reflect.Ptr && tmpty.String() == "*reflect.rtype" {
			insty := ins.(reflect.Type)
			for idx := 0; idx < insty.NumField(); idx++ {
				field := insty.Field(idx)
				if !ff(idx, nil, field) {
					break
				}
			}
		} else { // possible crash here if not match
			insRanger := ins.([]any)
			for idx, in := range insRanger {
				if !ff(idx, nil, in) {
					break
				}
			}
		}
	}

	out = IfElse(retmap, outmap, outarr)
	return
}

func DomapTypeField(ty reflect.Type, f func(reflect.StructField) interface{}) (outs []interface{}) {
	outs = make([]interface{}, 0)

	for idx := 0; idx < ty.NumField(); idx++ {
		field := ty.Field(idx)
		out := f(field)
		outs = append(outs, out)
	}

	return
}

// ///// todo more...
var vecmapconvfns = map[string]func(any) any{
	"any2string": func(vx any) any { return ToStr(vx) },
	// "any2bool":
	"bool2int":   func(vx any) any { return IfElse2(vx.(bool), 1, 0) },
	"bool2int64": func(vx any) any { return IfElse2(vx.(bool), int64(1), int64(0)) },

	"string2int64": func(vx any) any {
		s := vx.(string)
		if IsNumberic(s) {
			if strings.Contains(s, ".") {
				v := MustFloat64(s)
				return int64(v)
			} else {
				v := MustInt64(s)
				return v
			}
		}
		return -1
	},

	"string2real": func(vx any) any {
		v := MustFloat64(vx.(string))
		return v
	},
	"string2float64": func(vx any) any {
		v := MustFloat64(vx.(string))
		return v
	},
	"string2float32": func(vx any) any {
		v := MustFloat32(vx.(string))
		return v
	},
}

func init() {
	var str2i64fn = vecmapconvfns["string2int64"]
	vecmapconvfns["string2int"] = func(vx any) any {
		return int(str2i64fn(vx).(int64))
	}
	vecmapconvfns["string2uint"] = func(vx any) any {
		return uint(str2i64fn(vx).(int64))
	}
	vecmapconvfns["string2uint64"] = func(vx any) any {
		return uint64(str2i64fn(vx).(int64))
	}
	vecmapconvfns["string2int32"] = func(vx any) any {
		return int32(str2i64fn(vx).(int64))
	}
	vecmapconvfns["string2uint32"] = func(vx any) any {
		return uint32(str2i64fn(vx).(int64))
	}
	vecmapconvfns["string2uintptr"] = func(vx any) any {
		return uintptr(str2i64fn(vx).(int64))
	}
}

// primity type
// vecmapconvertvalueg[int](1.23)
func vecmapconvertvalueg[T any](vx any) (rv T, ok bool) {
	ety := reflect.TypeOf(vx)
	toty := reflect.TypeFor[T]()

	rvx, okx := vecmapconvertvalue(vx, ety, toty)
	ok = okx
	rv = rvx.(T)
	return
}

// primity type
// vecmapconvertvalue(1.23).(int)
func vecmapconvertvalue(vx any, ety, toty reflect.Type) (any, bool) {
	var rvx any

	if reflect.DeepEqual(ety, toty) {
		rvx = vx
	} else if ety.ConvertibleTo(toty) || ety.AssignableTo(toty) {
		rvx = reflect.ValueOf(vx).Convert(toty).Interface()
	} else {
		fnname := fmt.Sprintf("%s2%s", ety.String(), toty.String())
		if toty.Kind() == reflect.String {
			fnname = fmt.Sprintf("any2%s", toty.String())
		}
		if fn, ok := vecmapconvfns[fnname]; ok {
			rvx = fn(vx)
		} else {
			Infop("IVConvert failed", fnname)
			return rvx, false
		}
	}
	return rvx, true
}

func MapConvert[KT comparable, VT any](items map[any]any) (outs map[KT]VT) {
	if items == nil {
		return nil
	}
	outs = make(map[KT]VT, 0)
	if len(items) == 0 {
		return
	}

	kety := reflect.TypeOf(items).Key()
	vety := reflect.TypeOf(items).Elem()

	ktoty := reflect.TypeOf(outs).Key()
	vtoty := reflect.TypeOf(outs).Elem()

	for kx, vx := range items {
		tokeyx, keyconvok := vecmapconvertvalue(kx, kety, ktoty)
		tovalx, valconvok := vecmapconvertvalue(vx, vety, vtoty)
		if !keyconvok || !valconvok {
			break
		}
		outs[tokeyx.(KT)] = tovalx.(VT)
	}

	return
}

// 转换成显式类型
func IVConvert[T any](items []any) (outs []T) {
	if items == nil {
		return nil
	}
	outs = make([]T, 0)
	if len(items) == 0 {
		return outs
	}

	ety := reflect.TypeOf(items[0])
	toty := reflect.TypeOf(outs).Elem()

	for _, vx := range items {
		rvx, ok := vecmapconvertvalue(vx, ety, toty)
		if !ok {
			break
		}
		outs = append(outs, rvx.(T))
	}

	return
}

// interface vector to strings
func IV2Strings(items []any) []string {
	return IVConvert[string](items)
}
func IV2Ints(items []any) []int {
	return IVConvert[int](items)
}
func Strs2IV(items []string) []any {
	outx := Mapdo(items, func(s any) any { return s })
	return outx.([]any)
}

// enumerate类似功能
// 第一种方式，采用数组,可能用内存比较多
// usage: for i := range gopp.Range(5){}
// 88.88 ns/op
func _RangeA(n int) (rg []int) {
	rg = make([]int, n)
	for i := 0; i < n; i++ {
		rg[i] = i
	}
	return
}

var rangeaunderarr = make([]int, 0, 3)
var rangeaunderfillpos = 0

// 25.55 ns/op 不新分配内存就是快
// 但是最好不要太大的n，内存一直在
func RangeA[T int | float64](nx T) (rg []int) {
	n := int(nx)
	growed := false
	if n > cap(rangeaunderarr) {
		growed = true
		rangeaunderarr = slices.Grow(rangeaunderarr, n)
	}
	// rg = make([]int, n)
	rg = rangeaunderarr[:n]
	if growed || n > rangeaunderfillpos {
		rangeaunderfillpos = n - 1
		for i := 0; i < n; i++ {
			rg[i] = i
		}
	}
	return
}

// 第二种方式，采用channel。由于用到一个goroutine，可能效率慢
// 果然这个速度慢，不要使用，8290 ns/op
func RangeC(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// TODO
type Iterable interface {
	Iter() any
	Next() any
}

// string/map/slice/struct or implementation Iterable
func CanIter(v interface{}) bool {
	return false
}
