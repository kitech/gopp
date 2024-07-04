package cgopp

import (
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
	return voidptr(sym)
}
