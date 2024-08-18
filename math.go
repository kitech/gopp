package gopp

import (
	"log"
	"math"
	"reflect"
	// "golang.org/x/exp/constraints"
)

func Abs[T int | int64 | int32 | float32 | float64](vx T) T {
	if vx < 0 {
		return -vx
	}
	return vx
}

func AbsNum(x interface{}) interface{} {
	f1 := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	f2 := func(x float64) float64 {
		if x < 0.0 {
			return -x
		}
		return x
	}

	fns := []interface{}{f1, f2}
	fnidx := SymbolResolveFns([]interface{}{x}, fns)
	if fnidx == -1 {
		log.Panicln("Unresolved")
	}

	out := overloadRcCall([]interface{}{x}, fns[fnidx])
	return out[0].Interface()
}

func AbsI64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsI32(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxU32(nums []uint32) uint32 {
	var ret uint32
	for _, n := range nums {
		if n > ret {
			ret = n
		}
	}
	return ret
}

func MinU32(nums []uint32) uint32 {
	var ret uint32 = math.MaxUint32
	for _, n := range nums {
		if n < ret {
			ret = n
		}
	}
	return ret
}

func Maxv(a0, a1 interface{}, elses ...interface{}) interface{} {
	return Max(append(elses, a0, a1))
}

// support, slice of number/string
func Max(arr interface{}) interface{} {
	arrv := reflect.ValueOf(arr)
	arrty := arrv.Type()
	switch arrty.Kind() {
	case reflect.Array, reflect.Slice:
	default:
		return nil
	}

	switch arrty.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Float32, reflect.Float64:
	case reflect.String:
	default:
		return nil
	}

	var retv reflect.Value
	for i := 0; i < arrv.Len(); i++ {
		itemv := arrv.Index(i)
		if i == 0 {
			retv = itemv
		} else {
			switch arrty.Elem().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if itemv.Int() > retv.Int() {
					retv = itemv
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if itemv.Uint() > retv.Uint() {
					retv = itemv
				}
			case reflect.Float32, reflect.Float64:
				if itemv.Float() > retv.Float() {
					retv = itemv
				}
			case reflect.String:
				if itemv.String() > retv.String() {
					retv = itemv
				}
			}
		}
	}

	return retv.Interface()
}
func Minv(a0, a1 interface{}, elses ...interface{}) interface{} {
	return Min(append(elses, a0, a1))
}
func Min(arr interface{}) interface{} {
	arrv := reflect.ValueOf(arr)
	arrty := arrv.Type()
	switch arrty.Kind() {
	case reflect.Array, reflect.Slice:
	default:
		return nil
	}

	switch arrty.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Float32, reflect.Float64:
	case reflect.String:
	default:
		return nil
	}

	var retv reflect.Value
	for i := 0; i < arrv.Len(); i++ {
		itemv := arrv.Index(i)
		if i == 0 {
			retv = itemv
		} else {
			switch arrty.Elem().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if itemv.Int() > retv.Int() {
					retv = itemv
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if itemv.Uint() > retv.Uint() {
					retv = itemv
				}
			case reflect.Float32, reflect.Float64:
				if itemv.Float() > retv.Float() {
					retv = itemv
				}
			case reflect.String:
				if itemv.String() > retv.String() {
					retv = itemv
				}
			}
		}
	}

	return retv.Interface()
}

type Minmaxer[T Ordered] struct {
	Min    T
	Max    T
	inited bool
}

func NewMinmaxer[T Ordered]() *Minmaxer[T] {
	me := &Minmaxer[T]{}
	return me
}

func (me *Minmaxer[T]) Input(v T) {
	if me.inited {
		// why syntax not work???
		if v > me.Max {
			me.Max = v
		} else if v < me.Min {
			me.Min = v
		}
	} else {
		me.inited = true
		me.Min, me.Max = v, v
	}
}
func (me *Minmaxer[T]) Inputs(vs ...T) {
	for _, v := range vs {
		me.Input(v)
	}
}

func (me *Minmaxer[T]) Get() (T, T) {
	return me.Min, me.Max
}
