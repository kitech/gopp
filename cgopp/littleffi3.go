package cgopp

import (
	"log"
	"reflect"

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
litffi3_test1(double a, void*b, int64_t c) {
	// printf("%s: %f, %d, %p=%ld, %lld, \n", __FUNCTION__, a, (int)a, b, (uintptr_t)b, c);
	return a;
    return (int)(a);
}

float
litffi3_test2(float a) {
	printf("%s: %f\n", __FUNCTION__, a);
	return a+1;
}

*/
import "C"

func TestLitffi3callz() {
	sym, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi3_test1")
	// sym2, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi_test2")
	log.Println(sym, voidptr(sym))
	// todo how let go not checkptr, 用 usize传递
	// -race failed: fatal error: checkptr: pointer arithmetic computed bad pointer value
	// argp0 := voidptr(usize(3309))
	// argp0 := voidptr(usize(0)) // ok
	argp0 := usize(3309)
	x := Ffi3Call[float64](sym, float64(123.2345), argp0, uint64(386))
	log.Println("ret1", gopp.IfElse2(x == 123.2345, "OK", "unwant"), x)

	// float**的参数和返回值类型一定要与C函数匹配
	{
		// 这个应该也是不支持的，不支持为什么结果正确
		x := Ffi3Call0[float32]("litffi3_test1", float32(123.2345))
		log.Println("ret2", gopp.IfElse2(x == 123.2345, "OK", "unwant"), x)
	}
	{
		// 不支持的，返回结果未定义
		x := Ffi3Call0[float64]("litffi3_test1", float32(123.2345))
		log.Println("ret3", gopp.IfElse2(gopp.FloatIsZero(x), "OK", "unwant"), x)
	}
	{
		v := float32(1.23)
		x := Ffi3Call0[float32]("litffi3_test2", v)
		log.Println("ret4", gopp.IfElse2(x == 2.23, "OK", "unwant"), x)
	}
}

func BMLitffi3callz() {
	fnsym, _ := purego.Dlsym(purego.RTLD_DEFAULT, "litffi3_test1")

	gopp.Benchfn(func() {
		argp0 := usize(3309)
		x := Ffi3Call[float64](fnsym, float64(123.2345), argp0, uint64(386))
		_ = x
	}, 9999, gopp.MyFuncName())
}

// //////////
// todo 也许还需要做类型转换，像Pointer类型，可能要转换为voidptr
// todo 实现为prepare，也许可以提高一点效率
// 支持浮点数返回值
// 支持primitive类型参数，以及 Pointer/voidptr, C.type
// 不支持结构体类型，类RECORD类型
// 不支持go的map/slice/chan
// 支持go能够支持的任意多个参数
// 支持go string 传递参数，自动转换为charptr。但是C端不要持有该字符串指针，函数调用完成释放掉的
// go string 转换为 const char*，如果C端要持有该参数所有权，则调用前自己分配char*参数
// 如果没有返回值，使用FfiCall[int]()即可
// Usage1: FfiCall[float64]()
func Ffi3Call[RETY any, FT voidptr | usize](fnptrx FT, args ...any) (rvx RETY) {

	rety := reflect.TypeOf(rvx)
	fnty := fntypebyargs(rety, args...)
	fnv := reflect.New(fnty)
	// log.Println(fnv.UnsafeAddr()) // not works
	// log.Println(fnv.UnsafePointer()) // works but useless
	purego.RegisterFunc(fnv.Interface(), usize(fnptrx))
	gopp.NilPrint(fnv.Interface(), "regfunc failed/nil", fnv, fnv.Interface(), fnty)

	invals := make([]reflect.Value, len(args))
	for i, argx := range args {
		v := reflect.ValueOf(argx)
		ty := v.Type()
		switch ty.Kind() {
		case reflect.String:
			v = reflect.ValueOf(voidptr(CStringaf(argx.(string))))
		default:
		}
		invals[i] = v
	}
	outvals := fnv.Elem().Call(invals)
	// log.Println("fficalldone", outvals)
	rvx = outvals[0].Interface().(RETY)
	return
}

func Ffi3Call0[T any](name string, args ...any) T {
	sym := Dlsym0(name)
	return Ffi3Call[T](sym, args...)
}

// rety 如果是空，则设置为int
func fntypebyargs(rety reflect.Type, args ...any) reflect.Type {
	intys := make([]reflect.Type, len(args))
	outtys := []reflect.Type{rety}

	for i, argx := range args {
		intys[i] = reflect.TypeOf(argx)
	}

	fnty := reflect.FuncOf(intys, outtys, false)
	return fnty
}
