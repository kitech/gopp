package cgopp

import (
	"log"
	"reflect"
	"sync"
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

// called by c like other lang
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
	if true {
		// return
	}
	// log.Println(*funcname) //, ins, out)
	ins2 := *(*[]reflect.Value)(ins)
	out2 := *(*[]reflect.Value)(out)

	irgo_ffi_call2(*funcname, ins2, out2)
}

func irgo_ffi_call2(funcname string, ins []reflect.Value, out []reflect.Value) {
	// log.Println(len(ins), cap(ins), unsafe.Sizeof(reflect.Value{}))
	// log.Println(ins[0].Type())
	// // log.Println(ins)
	// fno := reflect.ValueOf(demoffifn)
	// log.Println(fno, fno.Type())
	// o2 := fno.Call(ins)
	// log.Println(o2)
	var out2 []reflect.Value
	fnx, ok := irgocbfns.Load(funcname)
	// gopp.FalsePrint(ok, "irgofn 404 in", irgocbfnscnt, funcname, "Try Dlsymgo mode...")
	if !ok {
		pc := Dlsymgo(funcname)
		gopp.ZeroPrint(pc, "cannot get symbol:", funcname, "in", irgocbfnscnt)
		if pc == 0 {
			return
		}
		fnv := MakeFuncByargs(pc, ins, out)
		IrgoRegistFunc(funcname, fnv.Elem().Interface())
		out2 = fnv.Elem().Call(ins)
	} else {
		fno := fnx.(reflect.Value)
		out2 = fno.Call(ins)
	}
	for i, o := range out2 {
		out[i] = o
	}
}

var irgocbfns = sync.Map{} // string => func value
var irgocbfnscnt = 0

func init() { IrgoRegistFunc("demoffifn", demoffifn) }
func demoffifn(a0 int32, a1 string, a2 float64) {
	// log.Println(a0, a1, a2)
}

func IrgoRegistFunc(funcname string, fno any) bool {
	if vx, ok := irgocbfns.Load(funcname); ok {
		log.Println("Exist?", funcname, vx)
		return false
	}
	irgocbfns.Store(funcname, reflect.ValueOf(fno))
	irgocbfnscnt += 1
	gopp.TruePrint(irgocbfnscnt > 9999, "so many!!!", irgocbfnscnt)
	return true
}
func IrgoUnregFunc(funcname string) bool {
	irgocbfns.Delete(funcname)
	return true
}
func IrgoRegistCount() int { return irgocbfnscnt }

func MakeFuncByargs(pc usize, ins []reflect.Value, out []reflect.Value) reflect.Value {
	// reflect.FuncOf()
	//reflect.MakeFunc()

	intys := make([]reflect.Type, len(ins))
	outys := make([]reflect.Type, len(out))
	for i, tv := range ins {
		intys[i] = tv.Type()
		// log.Println(i, tv.Type(), tv)
	}
	for i, tv := range out {
		if !tv.IsValid() {
			outys = outys[:i]
			break
		}
		outys[i] = tv.Type()
	}

	fnty := reflect.FuncOf(intys, outys, false)
	fnv := reflect.New(fnty)
	CreateFuncForCodePtr(fnv.Interface(), pc)
	return fnv
}

// Convenience struct for modifying the underlying code pointer of a function
// value. The actual struct has other values, but always starts with a code
// pointer.
type Func struct {
	codePtr uintptr
}

// CreateFuncForCodePtr is given a code pointer and creates a function value
// that uses that pointer. The outFun argument should be a pointer to a function
// of the proper type (e.g. the address of a local variable), and will be set to
// the result function value.
func CreateFuncForCodePtr(outFuncPtr interface{}, codePtr uintptr) {
	outFuncVal := reflect.ValueOf(outFuncPtr).Elem()
	// Use reflect.MakeFunc to create a well-formed function value that's
	// guaranteed to be of the right type and guaranteed to be on the heap
	// (so that we can modify it). We give a nil delegate function because
	// it will never actually be called.
	newFuncVal := reflect.MakeFunc(outFuncVal.Type(), nil)
	// Use reflection on the reflect.Value (yep!) to grab the underling
	// function value pointer. Trying to call newFuncVal.Pointer() wouldn't
	// work because it gives the code pointer rather than the function value
	// pointer. The function value is a struct that starts with its code
	// pointer, so we can swap out the code pointer with our desired value.
	funcValuePtr := reflect.ValueOf(newFuncVal).FieldByName("ptr").Pointer()
	funcPtr := (*Func)(unsafe.Pointer(funcValuePtr))
	funcPtr.codePtr = codePtr
	outFuncVal.Set(newFuncVal)
}
