package gopp

import (
	"fmt"
	"reflect"
	"strconv"
)

var ZeroStr string

func MustInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	n, err := strconv.Atoi(s)
	ErrPrint(err, s)
	return n
}

func MustUint32(s string) uint32 {
	if len(s) == 0 {
		return 0
	}
	n, err := strconv.Atoi(s)
	ErrPrint(err, s)
	return uint32(n)
}

func MustInt64(s string) int64 {
	if len(s) == 0 {
		return 0
	}
	n, err := strconv.ParseInt(s, 10, 64)
	ErrPrint(err, s)
	return n
}

func MustFloat64(s string) float64 {
	if len(s) == 0 {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	ErrPrint(err, s)
	return v
}
func MustFloat32(s string) float32 {
	if len(s) == 0 {
		return 0
	}
	v, err := strconv.ParseFloat(s, 32)
	ErrPrint(err, s)
	return float32(v)
}

func ToStr(v interface{}) string { return fmt.Sprintf("%v", v) }
func ToStrs(args ...interface{}) (rets []string) {
	for _, arg := range args {
		rets = append(rets, ToStr(arg))
	}
	return
}
func ToStrs2(argsx any) (rets []string) {
	args, ok := argsx.([]any)
	if !ok {
		return
	}
	rets = ToStrs(args...)
	return
}

// support bool, string, *int*, uintptr, unsafe.Pointer, float*
func Toint(vx any) int {
	var rv int
	switch v := vx.(type) {
	case bool:
		if v {
			rv = 1
		} else {
			rv = 0
		}
	case string:
		rv = MustInt(v)

	default:
		rvty := reflect.TypeOf(rv)
		vv := reflect.ValueOf(vx)
		if vv.CanConvert(rvty) {
			rv = (vv.Convert(rvty).Interface()).(int)
		} else {
			Infop("not support", reflect.TypeOf(vx), vx)
		}
	}
	return rv
}
