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
litffi_test1(double a, void*b, int64_t c) {
	printf("%f, %d, %p=%d, %d, \n", a, (int)a, b, (int)b, c);
	return a;
    return (int)(a);
}

float
litffi_test2(float a) {
	printf("%f\n",a);
	return a+1;
}

*/
import "C"

func TestLitfficallz() {
	sym, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi_test1")
	// sym2, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi_test2")
	// log.Println(sym)
	x := FfiCall[float64](voidptr(sym), float64(123.2345), voidptr(uintptr(3309)), uint64(386))
	log.Println(x)
	{
		x := FfiCall0[float32]("litffi_test1", float32(123.2345))
		log.Println(x)
	}
	{
		v := int32(0)
		*((*float32)(voidptr(&v))) = 1.23
		x := FfiCall0[float32]("litffi_test2", v)
		log.Println(x, v)
	}
}

// ///////////
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
// 支持go string 传递参数，自动转换为charptr。但是C端不要持有该字符串指针，函数调用完成释放掉的
// 如果没有返回值，使用[int]即可
// Usage1: FfiCall[float64]()
// Usage1: FfiCall(FFITY_FLOAT)
func FfiCall[T any](fnptrx voidptr, args ...any) (rvx T) {
	var args2 = make([]any, 5)
	var tystrs = make([]string, 5)
	for i, arg := range args {
		ty := reflect.TypeOf(arg)
		switch ty.Kind() {
		case reflect.Int, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Int16, reflect.Uint16, reflect.Int8, reflect.Uint8:
			if ty.Size() <= 4 {
				tv := reflect.ValueOf(arg).Convert(gopp.Int32Ty).Interface().(int32)
				args2[i] = tv
			} else {
				tv := reflect.ValueOf(arg).Convert(gopp.Int64Ty).Interface().(int64)
				args2[i] = tv
			}
			ty = gopp.IfElse2(ty.Size() == 4, gopp.Int32Ty, gopp.Int64Ty)
		case reflect.Bool: // nice, its works
			tv := int32(Go2cBool(arg.(bool)))
			args2[i] = tv
			ty = gopp.Int32Ty

		case reflect.UnsafePointer:
			if ty.Size() == 4 {
				tv := int32(usize(arg.(voidptr)))
				args2[i] = tv
			} else {
				tv := int64(usize(arg.(voidptr)))
				args2[i] = tv
			}
			ty = gopp.IfElse2(ty.Size() == 4, gopp.Int32Ty, gopp.Int64Ty)

		case reflect.String:
			tv := CString(arg.(string))
			defer cfree_voidptr(tv)
			args2[i] = voidptr(tv)
			ty = gopp.IfElse2(ty.Size() == 4, gopp.Int32Ty, gopp.Int64Ty)

		case reflect.Int32, reflect.Int64: // just fine
		case reflect.Float64, reflect.Float32:
		default:
			gopp.Info(ty.String(), arg)
		}
		tystrs[i] = ty.String()
	}
	for i := len(args); i < 5; i++ {
		tv := int32(0)
		args2[i] = tv
		tystrs[i] = gopp.Int32Ty.String()
	}
	gopp.FalsePrint(len(tystrs) == 5 && len(args2) == 5, "some error", len(tystrs), len(args))

	var tystr = strings.Join(tystrs, "_")
	var tycrc uint64
	tycrc = gopp.Crc64Str(tystr)

	// log.Println(tystrs, tycrc, tystr)
	var rv = litfficallgenimpl[T](tycrc, uintptr(fnptrx), args2...)
	gopp.GOUSED(rv)
	// var retptr = &rvx
	// *retptr = *((*T)(voidptr(&rv)))
	rvx = rv

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
