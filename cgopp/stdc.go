package cgopp

import (
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/kitech/gopp"
)

/*
#include <string.h>
#include <stdlib.h>
// macos not found
// #include <malloc.h>
// #include <memory.h>

*/
import "C"

type GoIface struct {
	Type voidptr
	Data voidptr
}

// std c library functions
// 这么封装一次，引用的包不需要再显式的引入"C"包。让CGO代码转换不传播到引用的代码中
func Cmemcpy(dst voidptr, src voidptr, n usize) voidptr {
	rv := C.memcpy(dst, src, C.size_t(n))
	return rv
}
func cfree_voidptr(ptr voidptr) { C.free(ptr) }
func Cfree(ptrx any) {
	switch ptr := ptrx.(type) {
	case unsafe.Pointer:
		cfree_voidptr(ptr)
	case uintptr:
		p := (voidptr(ptr))
		cfree_voidptr(p)
	case *C.char:
		p := (voidptr(ptr))
		cfree_voidptr(p)
	case *C.schar:
		p := (voidptr(ptr))
		cfree_voidptr(p)
	case *C.uchar:
		p := (voidptr(ptr))
		cfree_voidptr(p)
	case C.uintptr_t:
		p := (voidptr(usize(ptr)))
		cfree_voidptr(p)
	default:
		ty := reflect.TypeOf(ptrx)
		if ty.ConvertibleTo(gopp.VoidpTy()) {
			tv := reflect.ValueOf(ptrx)
			p := tv.Convert(gopp.VoidpTy()).Interface().(voidptr)
			cfree_voidptr(p)
		} else if ty.Kind() == reflect.Pointer &&
			strings.HasSuffix(ty.String(), "._Ctype_char") {
			var addr = (*GoIface)(voidptr(&ptrx))
			cfree_voidptr(addr.Data)
		} else {
			panic("unimpl " + ty.String() + " " + ty.Kind().String())
		}
	}
}

func Calloc(c, n int) voidptr {
	return C.calloc(csizet(c), csizet(n))
}
func Memset(ptr voidptr, c, n int) voidptr {
	return C.memset(ptr, cint(c), csizet(n))
}
func MemZero(ptr voidptr, n int) voidptr {
	return C.memset(ptr, 0, csizet(n))
}
func Strcpy() {}
func Strdup() {}

func Malloc(n int) voidptr {
	rv := C.malloc(csizet(n))
	return rv
}

// 可以用作存储C字符串？
// 不可以用作存储C++对象，无法调用destructor
func Mallocgc(n int) voidptr {
	return mymallocgc(usize(n), nil, true)
}

const CBoolTySz = gopp.Int32TySz
const CppBoolTySz = gopp.Int8TySz

// macos not this func
// let freed memory really given back to OS
// func MallocTrim() int { return int(C.malloc_trim(0)) }

func GoString(ptr voidptr) string {
	return C.GoString((*C.char)(ptr))
}
func GoStringN(ptr voidptr, len usize) string {
	return C.GoStringN((*C.char)(ptr), (C.int)(len))
}
func CString(s string) voidptr {
	return voidptr(C.CString(s))
}

// auto free after timeout
func CStringaf(s string) voidptr {
	ptr := voidptr(C.CString(s))

	time.AfterFunc(gopp.DurandSec(3, 5), func() { cfree_voidptr(ptr) })

	return ptr
}
