package gopp

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"sync"

	_ "unsafe"
)

//go:linkname mallocgc runtime.mallocgc
var mallocgc func(int) voidptr

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
