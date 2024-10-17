package cgopp

import (
	"log"
	"reflect"
	"unsafe"

	"github.com/kitech/gopp"
)

/*
// passive mode

*/
import "C"

// 两种模式，
// go作为exe入口，go是主调方
// go作为库使用，go是被调方
// 第三种方式，使用dlsym方式，这样不需要注册函数指针，只需要配置函数名字
// 只要链接到一起，不需要关心调用的顺序问题。
// \see github.com/ebitengine/purego

// 两种传参模式，
// 简单的JSON字段中
// 兼容的结构体模式

// passive mode

type GointeropFntype func(voidptr) voidptr

var irpasvcbfns = map[any]voidptr{}

// this must call nongo scope,
// such as c/rs/cxx
//
//export gointerop_set_passive_cbfn
func gointerop_set_passive_cbfn(name4c *C.char, fnptr voidptr) {
	var name = C.GoString(name4c)
	irpasvcbfns[name] = fnptr
	gopp.NilPrint(fnptr, "cannot set nil fnptr")
}

func GoirGetPasvCbfn[T string | int](name T) voidptr {
	if fnptr, ok := irpasvcbfns[name]; ok {
		return fnptr
	}
	return nil
}

// active mode

// dlsym mode

// irgo type funcs

//
//export irgo_get_gotype_object
func irgo_get_gotype_object(i int32, more *int32) uintptr {
	k := reflect.Kind(i)
	*more = 1
	if k > reflect.UnsafePointer {
		*more = 0
		return uintptr(0)
	}
	rvx := gopp.InvalidTy
	if vx, ok := gopp.RefKindTys[k]; ok {
		rvx = vx
	}
	// log.Println(i, reflect.TypeOf(rvx)) // *reflect._rtype
	rv := (*gopp.GoIface)(unsafe.Pointer(&rvx))

	if vx, ok := gopp.RefKindVals[k]; ok && vx.IsValid() {
		vp := (*gopp.Value)(voidptr(&vx))
		t0 := (*gopp.Abitype)(rv.Data)
		gopp.G_USED(vp, t0)
		// log.Println(i, rv.Data, voidptr(vp.Typ_), rv.Data == voidptr(vp.Typ_), t0.Kind_, vp.Typ_.Kind_)
		// stdout: 14 0x10c9b0940 0x10ca15160 false 14 54
		// return usize(voidptr(vp.Typ_))
		// note: reflect.TypeOf(v) != abi.Type
	}

	// ??? why rv.Data, should rv.Type???
	// yeap, rv.Data is type， rv.Type is type's type!!!
	return uintptr(rv.Data)
}

//export irgo_ffi_call
func irgo_ffi_call(funcname *string, ins voidptr, out voidptr) {
	log.Println(*funcname) //, ins, out)
	ins2 := *(*[]reflect.Value)(ins)
	out2 := *(*[]reflect.Value)(out)

	irgo_ffi_call2(*funcname, ins2, out2)
}

func demoffifn(a0, a1 int32, a2 float64) {
	log.Println(a0, a1, a2)
}
func irgo_ffi_call2(funcname string, ins []reflect.Value, out []reflect.Value) {
	log.Println(len(ins), cap(ins), unsafe.Sizeof(reflect.Value{}))
	log.Println(ins[0].Type())
	// log.Println(ins)
	fno := reflect.ValueOf(demoffifn)
	log.Println(fno, fno.Type())
	o2 := fno.Call(ins)
	log.Println(o2)
}
