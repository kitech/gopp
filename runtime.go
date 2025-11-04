package gopp

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"sync"
	"syscall"

	// "runtime/rtin"
	_ "unsafe"
)

// \see https://andrestc.com/post/go-memory-allocation-pt1/

// voidptr is runtime._type, and can be nil

////go:linkname mymallocgc runtime.mallocgc
//func mymallocgc(n usize, rttypeptr voidptr, zero bool) voidptr

// rttypeptr cannot be nil

////go:linkname mynewobject runtime.newobject
//func mynewobject(rttypeptr voidptr) voidptr

// 这个函数只负责固定指针地址不移动，但是并不负责持有指针引用
// 如果需要持有引用，直接使用runtime.Pinner.Pin

// go:linkname not need cgo enabled
// only allowed in Go files that import "unsafe"
//
////go:linkname setPinned runtime.setPinned
//func setPinned(ptr voidptr, pin bool) bool

////go:linkname acquirem runtime.acquirem
//func acquirem() (mp voidptr)

////go:linkname releasem runtime.releasem
//func releasem(mp voidptr)

////go:linkname firstmoduledata runtime.firstmoduledata
//var firstmoduledata moduledata

////go:linkname lastmoduledatap runtime.lastmoduledatap
//var lastmoduledatap *moduledata

// go1.22.3
// moduledata records information about the layout of the executable
// image. It is written by the linker. Any changes here must be
// matched changes to the code in cmd/link/internal/ld/symtab.go:symtab.
// moduledata is stored in statically allocated non-pointer memory;
// none of the pointers here are visible to the garbage collector.
type moduledata struct {
	_ voidptr //sys.NotInHeap // NotInHeap // sys.NotInHeap // Only in static data

	pcHeader     voidptr //  *pcHeader
	funcnametab  []byte
	cutab        []uint32
	filetab      []byte
	pctab        []byte
	pclntable    []byte
	ftab         []functab
	findfunctab  uintptr
	minpc, maxpc uintptr

	text, etext           uintptr
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	covctrs, ecovctrs     uintptr
	end, gcdata, gcbss    uintptr
	types, etypes         uintptr
	rodata                uintptr
	gofunc                uintptr // go.func.*

	textsectmap []int     // []textsect
	typelinks   []int32   // offsets from types
	itablinks   []voidptr // []*itab

	ptab []int // []ptabEntry

	pluginpath string
	pkghashes  []int // []modulehash

	// This slice records the initializing tasks that need to be
	// done to start up the program. It is built by the linker.
	inittasks []int // []*initTask

	modulename   string
	modulehashes []int // []modulehash

	hasmain uint8 // 1 if module contains the main function, 0 otherwise

	gcdatamask, gcbssmask bitvector

	typemap map[int]voidptr
	// typemap map[typeOff]*_type // offset to *_rtype in previous module

	bad bool // module failed to load and should be ignored

	next *moduledata
}

// /// type moduledata deps
// type NotInHeap = rtin.NotInHeap

// Information from the compiler about the layout of stack frames.
// Note: this type must agree with reflect.bitVector.
type bitvector struct {
	n        int32 // # of bits
	bytedata *uint8
}

type functab struct {
	entryoff uint32 // relative to runtime.text
	funcoff  uint32
}

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

// dont move
func SetPin(ptr voidptr, pin bool) {
	//	setPinned(ptr, pin)
}

func IsAndroid() bool { return runtime.GOOS == "android" }
func IsWindows() bool { return runtime.GOOS == "windows" }

// 需要关闭的对象的自动处理
// *os.File, *http.Response
func Deferx(objx interface{}) {
	if objx == nil {
		return
	}

	switch obj := objx.(type) {
	case *os.File:
		runtime.SetFinalizer(objx, func(objx interface{}) { obj.Close() })
	case *http.Response:
		runtime.SetFinalizer(objx, func(objx interface{}) { obj.Body.Close() })
	case *sync.Mutex:
		runtime.SetFinalizer(objx, func(objx interface{}) { obj.Unlock() })
	case *sync.RWMutex:
		runtime.SetFinalizer(objx, func(objx interface{}) { obj.Unlock() })
		// move to cgopp
	// case *C.char:
	//	runtime.SetFinalizer(objx, func(objx interface{}) { C.free(unsafe.Pointer(obj)) })
	case io.Closer:
		runtime.SetFinalizer(objx, func(objx interface{}) { obj.Close() })
		// TODO chan, context?
	default:
		objty := reflect.TypeOf(objx)
		log.Println("Unsupported:", objty.Kind().String())
	}
}

// runtime.setCgoTrace
// https://github.com/ianlancetaylor/cgosymbolizer
// https://github.com/benesch/cgosymbolizer

/*
	added GOROOT/src/runtime/gettid.go
	\\ see goinct.git/runtime/gettid.go

package runtime

// return os thread no, 0+, not really os thread id

	func Gettno() int32 {
		// struct runtime.p means process
		return getg().m.p.ptr().id
	}

// this is really pthread_self() value

	func Gettid() uint64 {
		// struct runtime.p means process
		return getg().m.procid
	}
*/
// func Gettid() uint64 { return runtime.Gettid() }
// func Gettno() int32  { return runtime.Gettno() }

func Gettid() int {
	tid := syscall.Gettid()
	return tid
}
