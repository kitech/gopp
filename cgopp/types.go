package cgopp

/*
#include <string.h>
#include <stdint.h>
*/
import "C"
import "unsafe"

// 把浮点数存储在uint64中
func Float64AsInt(n float64) (rv uint64) {
	C.memcpy((unsafe.Pointer(&rv)), (unsafe.Pointer(&n)), 8)
	return
}
func Float32AsInt(n float32) (rv uint64) {
	C.memcpy((unsafe.Pointer(&rv)), (unsafe.Pointer(&n)), 4)
	return
}
func IntAsFloat64(v uint64) (n float64) {
	C.memcpy((unsafe.Pointer(&n)), (unsafe.Pointer(&v)), 8)
	return
}
func IntAsFloat32(v uint64) (n float32) {
	C.memcpy((unsafe.Pointer(&n)), (unsafe.Pointer(&v)), 4)
	return
}
func U64ToPtr(v uint64) unsafe.Pointer    { return unsafe.Pointer(uintptr(v)) }
func U64OfPtr(vptr unsafe.Pointer) uint64 { return uint64(uintptr(vptr)) }

func C2goBool(ok C.int) bool {
	if ok == 0 {
		return false
	}
	return true
}
func Go2cBool(ok bool) C.int {
	if ok {
		return 1
	}
	return 0
}

type go2cfnty *[0]byte

// 参数怎么传递
// Go2cfn(cptr)
func Go2cfnp(fn unsafe.Pointer) *[0]byte {
	return go2cfnty(fn)
}

// Go2cfn(C.hello)
func Go2cfn(fn interface{}) *[0]byte {
	// assert(reflect.TypeOf(fn).Kind == reflect.Ptrx)
	return Go2cfnp(fn.(unsafe.Pointer))
}

// make go compiler happy, without show has unpined go pointer
func Anyptr2i[T any](v *T) uintptr {
	return uintptr(unsafe.Pointer(v))
}
func Anyptr2ci[T any](v *T) C.uintptr_t {
	return C.uintptr_t(Anyptr2i(v))
}
func Anyptr2cvptr[T any](v *T) *C.void {
	return (*C.void)(unsafe.Pointer(v))
}

func Iptr2Any[T any](v uintptr) *T {
	return (*T)(unsafe.Pointer(v))
}
func Ciptr2Any[T any](v C.uintptr_t) *T {
	return (*T)(unsafe.Pointer((uintptr)(v)))
}
func Cvptr2Any[T any](v *C.void) *T {
	return (*T)(unsafe.Pointer(v))
}
