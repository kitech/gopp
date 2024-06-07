package cgopp

/*
#include <stdio.h>
#include <stdlib.h>

void* litfficall(void* fnptr, int argc, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4) {
	printf("fnptr=%p,argc=%d,arg0=%p\n", fnptr, argc, arg0);
	typedef void* (*fnargc0)() ;
	fnargc0 fn0 = fnptr;
	typedef void* (*fnargc1)(void*) ;
	fnargc1 fn1 = fnptr;
	typedef void* (*fnargc2)(void*, void*) ;
	fnargc2 fn2 = fnptr;
	typedef void* (*fnargc3)(void*, void*, void*) ;
	fnargc3 fn3 = fnptr;
	typedef void* (*fnargc4)(void*, void*, void*, void*) ;
	fnargc4 fn4 = fnptr;

	void* retptr = 0;
	switch (argc) {
	case 0:
		retptr = fn0();
		break;
	case 1:
		retptr = fn1(arg0);
		break;
	case 2:
		retptr = fn2(arg0, arg1);
		break;
	case 3:
		retptr = fn3(arg0, arg1, arg2);
		break;
	case 4:
		retptr = fn4(arg0, arg1, arg2, arg3);
		break;
	}

	return retptr;
}

*/
import "C"
import (
	"log"
	"reflect"

	"github.com/kitech/gopp"
)

// note: 所有参数必须全是指针类型
// not support len(args) <= 5
func Litfficall(fnptrx voidptr, args ...voidptr) voidptr {
	if len(args) > 4 {
		log.Println("too many args, max", 5, ", but", len(args))
	}
	var argc = len(args)
	var argv = [5]voidptr{}
	for i, argx := range args {
		if i > 4 {
			break
		}
		argv[i] = argx
	}

	rv := C.litfficall(fnptrx, cint(argc), argv[0], argv[1], argv[2], argv[3], argv[4])
	return rv
}

// \see type go2cfnty
// note: 所有参数必须全是指针类型
// argsx, must can convert to voidptr
// not support len(args) <= 5
func Litfficallg[FT voidptr | *[0]byte](fnptrx FT, argsx ...any) voidptr {
	if len(argsx) > 4 {
		log.Println("too many args, max", 5, ", but", len(argsx))
	}

	var fnptr voidptr
	switch ptrx := any(fnptrx).(type) {
	case voidptr:
		fnptr = ptrx
	case *[0]byte:
		fnptr = voidptr(ptrx)
	}

	var argc = len(argsx)
	var argv = [5]voidptr{}
	for i, argx := range argsx {
		if i > 4 {
			break
		}
		switch vx := argx.(type) {
		case voidptr:
			argv[i] = vx
		case *C.char:
			argv[i] = voidptr(vx)
		default:
			// try reflect convert
			var voidptrty = gopp.VoidpTy()
			var argty = reflect.TypeOf(argx)
			if argty.Kind() == reflect.Pointer {
				var refval = reflect.ValueOf(argx)
				argv[i] = refval.UnsafePointer()
			} else if argty.ConvertibleTo(voidptrty) {
				var refval = reflect.ValueOf(argx).Convert(voidptrty)
				argv[i] = refval.UnsafePointer()
			} else if argty.Size() == gopp.VoidpTy().Size() {
				log.Println("maybe convable", argty.Size(), argty.String())
			} else {
				log.Println("ffiarg cannot convto voidptr", i, reflect.TypeOf(argx))
			}
		}
	}

	rv := C.litfficall(fnptr, cint(argc), argv[0], argv[1], argv[2], argv[3], argv[4])
	return rv
}
