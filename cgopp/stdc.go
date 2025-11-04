package cgopp

import (
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
)

/*
// fix macos nix clang16.0 cgo CoreFoundation and libobjc.A.dylib not found
#cgo darwin LDFLAGS: -F/Library/Developer/CommandLineTools/SDKs/MacOSX11.sdk/System/Library/Frameworks -L/Library/Developer/CommandLineTools//SDKs/MacOSX11.sdk/usr/lib
#cgo CFLAGS: -O0

#include <string.h>
#include <stdlib.h>
#include <stdint.h>
// macos not found
// #include <malloc.h>
// #include <memory.h>

*/
import "C"

// std c library functions
// 这么封装一次，引用的包不需要再显式的引入"C"包。让CGO代码转换不传播到引用的代码中
//
// Deprecated: Use gopp.Cmemcpy instead.
func Cmemcpy(dst voidptr, src voidptr, n isize) voidptr {
	return gopp.Cmemcpy(dst, src, n)
	// rv := C.memcpy(dst, src, C.size_t(n))
	// return rv
}
func cfree_voidptr(ptr voidptr) { C.free(ptr) }
func Cfree(ptrx any) {
	var ptr0 voidptr
	switch ptr := ptrx.(type) {
	case unsafe.Pointer:
		ptr0 = ptr
	case uintptr:
		ptr0 = voidptr(ptr)
	case *C.char:
		ptr0 = voidptr(ptr)
	case *C.schar:
		ptr0 = voidptr(ptr)
	case *C.uchar:
		ptr0 = voidptr(ptr)
	case C.uintptr_t:
		ptr0 = voidptr(usize(ptr))
	default:
		ty := reflect.TypeOf(ptrx)
		if ty.ConvertibleTo(gopp.VoidpTy) {
			tv := reflect.ValueOf(ptrx)
			ptr0 = tv.Convert(gopp.VoidpTy).Interface().(voidptr)
		} else if ty.Kind() == reflect.Pointer &&
			strings.HasSuffix(ty.String(), "._Ctype_char") {
			var addr = (*gopp.GoIface)(voidptr(&ptrx))
			ptr0 = addr.Data
		} else {
			panic("unimpl " + ty.String() + " " + ty.Kind().String())
		}
	}
	if ptr0 != nil {
		gopp.Cfree(ptr0)
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

func Malloc(n int) voidptr {
	rv := C.malloc(csizet(n))
	return rv
}

var cmallocfn func(csizet) voidptr
var ccallocfn func(csizet, csizet) voidptr
var creallocfn func(voidptr, csizet) voidptr
var cfreefn func(voidptr)
var cmemsetfn func(voidptr, cint, csizet) voidptr

func init() {
	{
		fnadr, _ := purego.Dlsym(purego.RTLD_DEFAULT, "malloc")
		purego.RegisterFunc(&cmallocfn, fnadr)
	}
	{
		fnadr, _ := purego.Dlsym(purego.RTLD_DEFAULT, "calloc")
		purego.RegisterFunc(&ccallocfn, fnadr)
	}
	{
		fnadr, _ := purego.Dlsym(purego.RTLD_DEFAULT, "realloc")
		purego.RegisterFunc(&creallocfn, fnadr)
	}
	{
		fnadr, _ := purego.Dlsym(purego.RTLD_DEFAULT, "free")
		purego.RegisterFunc(&cfreefn, fnadr)
	}
	{
		fnadr, _ := purego.Dlsym(purego.RTLD_DEFAULT, "memset")
		purego.RegisterFunc(&cmemsetfn, fnadr)
	}
}
func Mallocpg(n int) voidptr {
	rv := cmallocfn(csizet(n))
	return rv
}
func Cfreepg(ptr voidptr) { cfreefn(ptr) }

func RttypeOf(v any) voidptr {
	var typtr voidptr = ((*gopp.GoIface)(voidptr(&v))).Type
	return typtr
}

// *byte's runtime._type instance
var gobyterttype = RttypeOf(byte(0))

// note: return ptr is Pinned
// 可以用作存储C字符串？
// 不可以用作存储C++对象，无法调用destructor
func Mallocgc(n int) voidptr {
	ptr := mymallocgc(usize(n), gobyterttype, true)
	//	setPinned(ptr, true)
	// must a type not unsafe.Pointer
	runtime.SetFinalizer((*byte)(ptr), func(obj any) {
		if false {
			log.Println("dtor", obj, n)
		}
	})
	return ptr
}

// keep it, cpp need

//export cgoppMallocgc
func cgoppMallocgc(n cint) voidptr {
	return Mallocgc(int(n))
}

const CBoolTySz = gopp.Int32TySz
const CppBoolTySz = gopp.Int8TySz

// macos not this func
// let freed memory really given back to OS
// func MallocTrim() int { return int(C.malloc_trim(0)) }

// Deprecated: Use gopp.GoString instead.
func GoString[T voidptr | charptr](ptr T) string {
	return gopp.GoString(voidptr(ptr))
	// return C.GoString((*C.char)(ptr))
}

// Deprecated: Use gopp.GoStringN instead.
func GoStringN[T voidptr | charptr](ptr T, len isize) string {
	return gopp.GoStringN(voidptr(ptr), len)
	// return C.GoStringN((*C.char)(ptr), (C.int)(len))
}

// Deprecated: Use gopp.CString instead.
func CString(s string) voidptr {
	return gopp.CString(s)
	// return voidptr(C.CString(s))
}

// too slow, 480ns/op, C.CString不过 100ns/op
// auto free after timeout
func CStringaf(s string) voidptr {
	ptr := voidptr(C.CString(s))

	time.AfterFunc(gopp.DurandSec(3, 5), func() { cfree_voidptr(ptr) })
	return ptr
}

// using go's mallocgc version. go memcpy version.
//
// dddDeprecated: Use gopp.CStringgc instead.
func CStringgc(s string) voidptr {
	ptr := Mallocgc(len(s) + 1)
	return gopp.Cmemcpy(ptr, gopp.StringData(s), len(s))
	// return gopp.Cmemcpy(ptr, (*gopp.GoStringIn)(voidptr(&s)).Ptr, len(s))
}

// using go's mallocgc version. C memcpy version
//
// dddDeprecated: Use gopp.Cmemcpy instead.
func CStringgc2(s string) voidptr {
	ptr := Mallocgc(len(s) + 1)
	o := (*gopp.GoStringIn)((voidptr)(&s))
	Cmemcpy(ptr, o.Ptr, o.Len)

	return ptr
}

// \see strings.Clone
//
// dddDeprecated: Use gopp.Cmemcpy instead.
func Gostrdup(s string) string {
	return strings.Clone(s)

	// ptr := CStringgc(s)
	// var rv string
	// o := ((*gopp.GoStringIn)((voidptr)(&rv)))
	// o.Ptr = ptr
	// o.Len = usize(len(s))

	// return rv
}

type GortType struct {
}

// 从cgo生成的临时文件中取得
// CString converts the Go string s to a C string.
//
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func _Cfunc_CString_demoo(s string) charptr {
	if len(s)+1 <= 0 {
		panic("string too large")
	}
	p := _cgo_cmalloc(uint64(len(s) + 1))
	sliceHeader := struct {
		p   unsafe.Pointer
		len int
		cap int
	}{p, len(s) + 1, len(s) + 1}
	b := *(*[]byte)(unsafe.Pointer(&sliceHeader))
	copy(b, s)
	b[len(s)] = 0
	return (charptr)(p)
}
