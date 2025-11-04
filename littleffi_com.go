package gopp

import (
	"reflect"

	"github.com/ebitengine/purego"
	// "github.com/kitech/gopp"
)

// 不好，purego只支持arm64,amd64, 不支持arm32,amd32. see purego/func.go:416
// purego 也没说RegisterFunc的函数是不是threadsafe
// _ "github.com/ebitengine/purego"

const (
	FFITY_NONE = iota
	FFITY_INT
	FFITY_INT64
	FFITY_STRING // charptr
	FFITY_FLOAT32
	FFITY_FLOAT64
	FFITY_POINTER
	FFITY_USIZE
)

var ffiver = FfiPurego

const Ffinone = 0
const FfiLita6 = 1
const FfiPurego = 3
const (
	// maybe go native support is comming
	// https://github.com/golang/go/issues/975
	// https://github.com/golang/go/issues/16623
	FfiVaarg = iota + 4
	FfiVtcm  // vtmpl + cmacro
)

func SwitchFfiver(v int) {
	if v == ffiver {
		return
	}
	if v == 2 || v == 3 {
		ffiver = v
	} else if v == FfiLita6 {
		Warn("Only voidptr param supported")
		ffiver = v
	} else {
		Warn("Invalid ffiver[2, 3], but", v)
	}
}

func TestLitfficallz() {
	switch ffiver {
	case 2:
		// TestLitffi2callz()
	default:
		// TestLitffi3callz()
	}
}

// Ffi3 Prep Call 9999 9.54493ms 954ns 1048218 /s
// Ffi3Call 9999 23.591148ms 2.359µs 423908 /s
// Ffi2Call 9999 17.42031ms 1.742µs 574052 /s
func FfiCall[RETY any, FT voidptr | usize | *[0]byte](fnptrx FT, args ...any) (rvx RETY) {
	switch ffiver {
	case 1:
		// rv := Litfficallg(fnptrx, args...)
		// rvx = RETY(rv)
		// rvx = *(*RETY)(voidptr(&rv))
	case 2:
		// rvx = Ffi2Call[RETY](fnptrx, args...)
	default:
		// rvx = Ffi3Call[RETY](fnptrx, args...)
	}
	return
}

// todo cxx
// 涉及的 CPP 的 name resolusion
func fficallcpp() {}

func FfiCallVoid[FT voidptr | usize | *[0]byte](fnptrx FT, args ...any) {
	switch ffiver {
	case 2:
		// Ffi2Call[int](fnptrx, args...)
	default:
		Ffi3Call[int](fnptrx, args...)
	}
}

func FfiCall0[T any](name string, args ...any) (rvx T) {
	fnsym := Dlsym0(name)
	NilPrint(fnsym, "symnil", name)
	switch ffiver {
	case 2:
		// rvx = Ffi2Call[T](fnsym, args...)
	default:
		rvx = Ffi3Call[T](fnsym, args...)
	}
	return
}

func FfiCallVoid0(name string, args ...any) {
	fnsym := Dlsym0(name)
	NilPrint(fnsym, "symnil", name)
	switch ffiver {
	case 2:
		// Ffi2Call[int](fnsym, args...)
	default:
		Ffi3Call[int](fnsym, args...)
	}
}

// func Dlsym0(name string) voidptr {
// 	// name := name[1:] // for macos???
// 	sym, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
// 	gopp.ErrPrint(err, name)
// 	gopp.NilPrint(sym, name)
// 	return voidptr(sym)
// }

// /// libffi like prepare method
type FfiCif[RETY any] struct {
	abi        int
	preped     bool
	isvar      bool
	nfixedargc int
	ntotalargc int

	rety   reflect.Type
	argtys []reflect.Type
	fnty   reflect.Type

	invals []reflect.Value
	fnv    reflect.Value
	rvx    RETY
}

func FfiCifNew[T any]() *FfiCif[T] {
	me := &FfiCif[T]{}
	me.rety = reflect.TypeOf(me.rvx)
	return me
}

// check argtype support state
func (me *FfiCif[T]) Prep(fnptrx any, args ...any) error {
	me.fnty, me.argtys = fntypebyargs(me.rety, args...)
	me.fnv = reflect.New(me.fnty)
	me.invals = make([]reflect.Value, len(args)) // 这个分配内存影响10%的效率差不多

	var ifv = (*GoIface)((voidptr)(&fnptrx))
	var fnptr = usize(*((*voidptr)(ifv.Data)))
	// var fnptr = reflect.ValueOf(fnptrx).Convert(gopp.UsizeTy).Interface().(usize)
	// switch fn := fnptrx.(type) {
	// case voidptr:
	// 	fnptr = usize(fn)
	// case usize:
	// 	fnptr = fn
	// default:
	// 	gopp.Warn(reflect.TypeOf(fnptrx), len(args), args)
	// }

	// 这个调用对效率的影响挺大的，可能有20%的次序影响
	var fnv = me.fnv
	purego.RegisterFunc(fnv.Interface(), fnptr)
	NilPrint(fnv.Interface(), "regfunc failed/nil", fnv, fnv.Interface(), me.fnty)

	return nil
}

func (me *FfiCif[T]) Call(fnptrx any, args ...any) (rvx T) {
	TruePrint(len(args) != len(me.argtys), "argc not match", len(args), len(me.argtys))

	// invals := make([]reflect.Value, len(args))
	invals := me.invals
	for i, argty := range me.argtys {
		if argty.Kind() == reflect.String {
			// invals[i] = reflect.ValueOf((CStringaf(argx.(string)))
			invals[i] = reflect.ValueOf(CStringaf(args[i].(string)))
		} else {
			invals[i] = reflect.ValueOf(args[i])
		}
	}

	fnv := me.fnv
	outvals := fnv.Elem().Call(invals)
	// log.Println("fficalldone", outvals)
	rvx = outvals[0].Interface().(T)

	return
}
func (me *FfiCif[T]) Call0(fnsymname string, args ...any) (rvx T) {
	fnsym := Dlsym0(fnsymname)
	rvx = me.Call(fnsym, args...)
	return
}

// varidic args
func (me *FfiCif[T]) PrepVar() {

}
