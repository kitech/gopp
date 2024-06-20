package cgopp

import (
	"log"
	"reflect"
	"strings"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
)

/*
//
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>

// int
double
litffi_test1(double a) {
	printf("%f, %d\n", a, (int)a);
	return a;
    return (int)(a);
}

*/
import "C"

const (
	FFITY_NONE = iota
	FFITY_INT
	FFITY_INT64
	FFITY_STRING // charptr
	FFITY_FLOAT32
	FFITY_FLOAT64
	FFITY_POINTER
	FFITY_USIZE
)

// //////////
// 支持浮点数返回值
// 支持最大5个参数
// 如果没有返回值，使用[int]即可
// Usage1: FfiCall[float64]()
// Usage1: FfiCall(FFITY_FLOAT)
func FfiCall[T any](fnptrx voidptr, args ...any) (rvx T) {
	if len(args) == 0 {
		var fnv func() T
		purego.RegisterFunc(&fnv, uintptr(fnptrx))
		rvx = fnv()
	}

	var tystrs []string
	for _, arg := range args {
		ty := reflect.TypeOf(arg)
		tystrs = append(tystrs, ty.String())
	}

	var tystr = strings.Join(tystrs, "_")
	var tycrc uint64
	tycrc = gopp.Crc64Str(tystr)

	log.Println(tystrs, tycrc, tystr)
	var rv = litfficallgenimpl[T](tycrc, uintptr(fnptrx), args...)
	gopp.GOUSED(rv)
	var retptr = &rvx
	*retptr = *((*T)(voidptr(&rv)))

	return
}

func FfiCall0[T any](name string, args ...any) T {
	sym := Dlsym0(name)
	return FfiCall[T](sym, args...)
}
func Dlsym0(name string) voidptr {
	sym, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
	gopp.ErrPrint(err, name)
	return voidptr(sym)
}

func TestLitfficallz() {
	sym, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi_test1")
	// log.Println(sym)
	x := FfiCall[float64](voidptr(sym), float64(123.2345))
	log.Println(x)
	{
		x := FfiCall0[float64]("litffi_test1", float64(123.2345))
		log.Println(x)
	}
}
