package cgopp

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
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
litffi2_test1(double a, void*b, int64_t c) {
	printf("%s: %f, %d, %p=%ld, %lld, \n", __FUNCTION__, a, (int)a, b, (uintptr_t)b, c);
	return a;
    return (int)(a);
}

float
litffi2_test2(float a) {
	printf("%s: %f\n", __FUNCTION__, a);
	return a+1;
}

*/
import "C"

func TestLitffi2callz() {
	sym, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi2_test1")
	// sym2, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi_test2")
	// log.Println(sym)
	x := Ffi2Call[float64](sym, float64(123.2345), voidptr(uintptr(3309)), uint64(386))
	log.Println("ret1", gopp.IfElse2(x == 123.2345, "OK", "unwant"), x)
	{
		// 如果C中是double返回值，则用float32接收为0
		x := Ffi2Call0[float32]("litffi2_test1", float32(123.2345))
		log.Println("ret2", gopp.IfElse2(gopp.FloatIsZero(x), "OK", "unwant"), x, x > 0, x > 1.0)
	}
	{
		x := Ffi2Call0[float64]("litffi2_test1", float32(123.2345))
		log.Println("ret3", gopp.IfElse2(x == 123.2345, "OK", "unwant"), x)
	}
	{
		// 如果C中是float32的参数，则无法传递，无法支持调用这种C函数
		v := float32(1.23)
		x := Ffi2Call0[float32]("litffi2_test2", v)
		log.Println("ret4", gopp.IfElse2(x > 58486031497623648.0, "OK", "unwant"), x)
	}
}
func BMLitffi2callz() {
	fnsym, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi3_test1")

	gopp.Benchfn(func() {
		argp0 := usize(3309)
		x := Ffi2Call[float64](fnsym, float64(123.2345), argp0, uint64(386))
		_ = x
	}, 99999, gopp.MyFuncName())
}

// ///////////

// //////////
// 支持浮点数返回值
// 支持最大5个参数
// 不支持参数包含float32的C函数
// 支持go string 传递参数，自动转换为charptr。但是C端不要持有该字符串指针，函数调用完成释放掉的
// 如果没有返回值，使用[int]即可
// Usage1: FfiCall[float64]()
func Ffi2Call[RETY any, FT voidptr | usize | *[0]byte](fnptrx FT, args ...any) (rvx RETY) {
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
		case reflect.Float64:
			args2[i] = args[i]
		case reflect.Float32: // TODO 转换了C函数接收到也是错误的
			// args2[i] = float64(args[i].(float32)) // 这个转换就不太准确
			tvx := fmt.Sprintf("%v", args[i])
			tv, _ := strconv.ParseFloat(tvx, 64) // 这样准确
			args2[i] = tv
			ty = gopp.Float64Ty
			// log.Println(args2[i], args[i].(float32), tvx, tv)
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

	// log.Println(len(args), tystrs, tycrc, tystr)
	var rv = litfficallgenimpl[RETY](tycrc, uintptr(voidptr(fnptrx)), args2...)
	gopp.GOUSED(rv)
	// var retptr = &rvx
	// *retptr = *((*T)(voidptr(&rv)))
	rvx = rv

	return
}

func Ffi2Call0[T any](name string, args ...any) T {
	sym := Dlsym0(name)
	return Ffi2Call[T](sym, args...)
}
