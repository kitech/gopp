package main
// dont modify main

/*
#include <stdint.h>
*/
import "C"
import "unsafe"

type i32 = int32
type i64 = int64
type u32 = uint32
type u64 = uint64
type f32 = float32
type f64 = float64
type usize = uintptr
type vptr = unsafe.Pointer
type cuptr = C.uintptr_t
type cvptr = *C.void

	func anyptr2uptr[T any](p *T) usize {
		var pp = usize(vptr(p))
		return pp
		}
		

		func anyptr2uptrc[T any](p *T) cuptr{
			var pp = uintptr(vptr(p))
			return cuptr(pp)
		}		
			