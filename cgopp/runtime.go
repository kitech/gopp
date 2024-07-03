package cgopp

/*
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/kitech/gopp"
)

// \see https://andrestc.com/post/go-memory-allocation-pt1/

// voidptr is runtime._type, and can be nil

//go:linkname mymallocgc runtime.mallocgc
func mymallocgc(n usize, rttypeptr voidptr, zero bool) voidptr

// rttypeptr cannot be nil

//go:linkname mynewobject runtime.newobject
func mynewobject(rttypeptr voidptr) voidptr

// 需要关闭的对象的自动处理
// *os.File, *http.Response
func Deferx(objx interface{}) {
	if objx == nil {
		return
	}

	switch obj := objx.(type) {
	case *C.char:
		runtime.SetFinalizer(objx, func(objx interface{}) { C.free(unsafe.Pointer(obj)) })
	default:
		gopp.Deferx(objx)
	}
}
