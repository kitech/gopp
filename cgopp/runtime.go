package cgopp

/*
#include <stdlib.h>
*/
import "C"

import (
	"runtime"

	"github.com/kitech/gopp"
)

// \see https://andrestc.com/post/go-memory-allocation-pt1/

// voidptr is runtime._type, and can be nil

//go:linkname mymallocgc runtime.mallocgc
func mymallocgc(n usize, rttypeptr voidptr, zero bool) voidptr

// rttypeptr cannot be nil

//go:linkname mynewobject runtime.newobject
func mynewobject(rttypeptr voidptr) voidptr

// 这个函数只负责固定指针地址不移动，但是并不负责持有指针引用
// 如果需要持有引用，直接使用runtime.Pinner.Pin

//go:linkname setPinned runtime.setPinned
func setPinned(ptr voidptr, pin bool) bool

//go:linkname acquirem runtime.acquirem
func acquirem() (mp voidptr)

//go:linkname releasem runtime.releasem
func releasem(mp voidptr)

type FuncInfo struct {
	F  voidptr // *_func
	MD voidptr // datap   *moduledata
}

//go:linkname Rtfindfunc runtime.findfunc
func Rtfindfunc(uintptr) FuncInfo

//go:linkname Rtfuncname runtime.funcname
func Rtfuncname(FuncInfo) string

//go:linkname Rtfuncpkgpath runtime.funcpkgpath
func Rtfuncpkgpath(FuncInfo) string

//go:linkname Rtfuncfile runtime.funcfile
func Rtfuncfile(f FuncInfo, fileno int32) string

// this not work for
// Undefined symbols for architecture x86_64: "_runtime.getg"

// go:linkname getg runtime.getg
// func getg() (gr voidptr)

func SetPin(ptr voidptr, pin bool) {
	setPinned(ptr, pin)
}

// 需要关闭的对象的自动处理
// *os.File, *http.Response
func Deferx(objx interface{}) {
	if objx == nil {
		return
	}

	switch obj := objx.(type) {
	case charptr: // *C.char:
		runtime.SetFinalizer(objx, func(objx interface{}) { C.free(voidptr(obj)) })
	default:
		gopp.Deferx(objx)
	}
}
