package cgopp

import (
	"reflect"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
)

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

var ffiver = 3

func SwitchFfiver(v int) {
	if v == 2 || v == 3 {
		ffiver = v
	} else {
		gopp.Warn("Invalid ffiver[2, 3], but", v)
	}
}

func TestLitfficallz() {
	switch ffiver {
	case 2:
		TestLitffi2callz()
	default:
		TestLitffi3callz()
	}
}

// Ffi3Call 9999 23.591148ms 2.359µs 423908 /s
// Ffi2Call 9999 17.42031ms 1.742µs 574052 /s
func FfiCall[RETY any, FT voidptr | usize](fnptrx FT, args ...any) (rvx RETY) {
	switch ffiver {
	case 1:
		// rv := Litfficallg(fnptrx, args...)
		// rvx = RETY(rv)
	case 2:
		rvx = Ffi2Call[RETY](fnptrx, args...)
	default:
		rvx = Ffi3Call[RETY](fnptrx, args...)
	}
	return
}

// todo cxx
// 涉及的 CPP 的 name resolusion
func fficallcpp() {}

func FfiCallVoid[FT voidptr | usize](fnptrx FT, args ...any) {
	switch ffiver {
	case 2:
		Ffi2Call[int](fnptrx, args...)
	default:
		Ffi3Call[int](fnptrx, args...)
	}
}

func FfiCall0[T any](name string, args ...any) (rvx T) {
	fnsym := Dlsym0(name)
	gopp.NilPrint(fnsym, "symnil", name)
	switch ffiver {
	case 2:
		rvx = Ffi2Call[T](fnsym, args...)
	default:
		rvx = Ffi3Call[T](fnsym, args...)
	}
	return
}

func FfiCallVoid0(name string, args ...any) {
	fnsym := Dlsym0(name)
	gopp.NilPrint(fnsym, "symnil", name)
	switch ffiver {
	case 2:
		Ffi2Call[int](fnsym, args...)
	default:
		Ffi3Call[int](fnsym, args...)
	}
}

func Dlsym0(name string) voidptr {
	sym, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
	gopp.ErrPrint(err, name)
	gopp.NilPrint(sym, name)
	return voidptr(sym)
}

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
	// fnv  reflect.Value
	rvx RETY
}

func FfiCifNew[T any]() *FfiCif[T] {
	me := &FfiCif[T]{}
	me.rety = reflect.TypeOf(me.rvx)
	return me
}

// check argtype support state
func (me *FfiCif[T]) Prep(args ...any) error {
	for _, argx := range args {
		ty := reflect.TypeOf(argx)
		me.argtys = append(me.argtys, ty)
	}
	me.fnty = fntypebyargs(me.rety, args...)
	return nil
}

func (me *FfiCif[T]) Call(fnptrx any, args ...any) (rvx T) {
	gopp.FalsePrint(len(args) == len(me.argtys), "argc not match prep", me.argtys)
	fnv := reflect.New(me.fnty)

	var fnptr usize
	switch fn := fnptrx.(type) {
	case voidptr:
		fnptr = usize(fn)
	case usize:
		fnptr = fn
	default:
		gopp.Warn(reflect.TypeOf(fnptrx), len(args), args)
	}
	purego.RegisterFunc(fnv.Interface(), fnptr)
	gopp.NilPrint(fnv.Interface(), "regfunc failed/nil", fnv, fnv.Interface(), me.fnty)

	invals := make([]reflect.Value, len(args))
	for i, argx := range args {
		v := reflect.ValueOf(argx)
		ty := v.Type()
		switch ty.Kind() {
		case reflect.String:
			v = reflect.ValueOf(voidptr(CStringaf(argx.(string))))
		default:
		}
		invals[i] = v
	}

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
