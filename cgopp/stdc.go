package cgopp

/*
#include <string.h>
#include <stdlib.h>
// macos not found
// #include <malloc.h>
// #include <memory.h>

*/
import "C"

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/kitech/gopp"
)

type GoIface struct {
	Type voidptr
	Data voidptr
}

// std c library functions
// 这么封装一次，引用的包不需要再显式的引入"C"包。让CGO代码转换不传播到引用的代码中
func Cmemcpy()                  {}
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

func Calloc()   {}
func CMemset()  {}
func CMemZero() {}
func CStrcpy()  {}
func CStrdup()  {}

const CBoolTySz = gopp.Int32TySz
const CppBoolTySz = gopp.Int8TySz

// macos not this func
// let freed memory really given back to OS
// func MallocTrim() int { return int(C.malloc_trim(0)) }
