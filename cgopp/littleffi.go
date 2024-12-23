package cgopp

import (
	"log"
	"reflect"

	"github.com/kitech/gopp"

	// "github.com/ebitengine/purego"
	// _ "github.com/ebitengine/purego"
)

/*
#cgo LDFLAGS: -Wl,--export-dynamic

#ifndef EXPORT_API
#define EXPORT_API __attribute__ ((visibility("default")))
#endif

#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

EXPORT_API
void* litfficall(void* fnptr, int argc, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5) {
	// printf("cgopp.C.%s: fnptr=%p,argc=%d,arg0=%p\n", __FUNCTION__, fnptr, argc, arg0);
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
	typedef void* (*fnargc5)(void*, void*, void*, void*, void*) ;
	fnargc5 fn5 = fnptr;
	typedef void* (*fnargc6)(void*, void*, void*, void*, void*, void*) ;
	fnargc6 fn6 = fnptr;

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
	case 5:
		retptr = fn5(arg0, arg1, arg2, arg3, arg4);
		break;
	case 6:
		retptr = fn6(arg0, arg1, arg2, arg3, arg4, arg5);
		break;
	default:
		printf("cgopp.C.%s: casedft fnptr=%p,argc=%d,arg0=%p\n", __FUNCTION__, fnptr, argc, arg0);
	}

	return retptr;
}

*/
import "C"

var litfficallfnc func(voidptr, int, voidptr, voidptr, voidptr, voidptr, voidptr, voidptr) voidptr

// //go:linkname gointernal_dlsym dlsym
// var gointernal_dlsym uintptr

func init() {
	// gopp.ZeroPrint(gointernal_dlsym, "go:linkname not worked")
	// sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "litfficall")
	// gopp.ErrPrint(err)
	// purego.RegisterFunc(&litfficallfnc, sym)
}

// note: 所有参数必须全是指针类型
// not support len(args) <= 6
func Litfficall(fnptrx voidptr, args ...voidptr) voidptr {
	if len(args) > 6 {
		log.Println("too many args, max", 6, ", but", len(args))
	}
	var argc = len(args)
	var argv = [6]voidptr{}
	for i, argx := range args {
		if i > 5 {
			break
		}
		argv[i] = argx
	}

	rv := C.litfficall(fnptrx, cint(argc), argv[0], argv[1], argv[2], argv[3], argv[4], argv[5])
	// rv := litfficallfnc(fnptrx, argc, argv[0], argv[1], argv[2], argv[3], argv[4])
	return rv
}

// \see type go2cfnty
// note: 所有参数必须全是指针类型
// argsx, must can convert to voidptr
// not support len(args) <= 6
func Litfficallg[FT voidptr | uintptr | *[0]byte](fnptrx FT, argsx ...any) voidptr {
	if len(argsx) > 6 {
		log.Println("too many args, max", 6, ", but", len(argsx))
	}

	var fnptr voidptr
	switch ptrx := any(fnptrx).(type) {
	case voidptr:
		fnptr = ptrx
	case uintptr:
		fnptr = voidptr(ptrx)
	case *[0]byte:
		fnptr = voidptr(ptrx)
	}
	gopp.SetPin(fnptr, true)

	var argc = len(argsx)
	var argv = [6]voidptr{}
	for i, argx := range argsx {
		if i > 5 {
			break
		}
		if argx == nil {
			argv[i] = voidptr(nil)
			continue
		}
		switch vx := argx.(type) {
		case voidptr:
			argv[i] = vx
		case *C.char:
			argv[i] = voidptr(vx)
		case bool:
			argv[i] = voidptr(usize(gopp.IfElse2(vx, 1, 0)))
		default:
			// try reflect convert
			var voidptrty = gopp.VoidpTy
			var argty = reflect.TypeOf(argx)
			if argty.Kind() == reflect.Pointer {
				var refval = reflect.ValueOf(argx)
				argv[i] = refval.UnsafePointer()
			} else if argty.ConvertibleTo(voidptrty) {
				var refval = reflect.ValueOf(argx).Convert(voidptrty)
				argv[i] = refval.UnsafePointer()
			} else if argty.Size() == gopp.VoidpTy.Size() {
				log.Println("maybe convable", argty.Size(), argty.String())
			} else {
				log.Println("ffiarg cannot convto voidptr", i, reflect.TypeOf(argx))
			}
		}
		gopp.SetPin(argv[i], true)
	}

	// todo panic: runtime error: cgo argument has Go pointer to unpinned Go pointer
	rv := C.litfficall(fnptr, cint(argc), argv[0], argv[1], argv[2], argv[3], argv[4], argv[5])
	// rv := litfficallfnc(fnptr, argc, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5])
	return rv
}
