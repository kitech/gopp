package gopp

import (
	"log"
	"unsafe"

	"github.com/ebitengine/purego"
)

var Dlopen = purego.Dlopen
var Dlsym = purego.Dlsym
var Dlclose = purego.Dlclose

type Dlerror = purego.Dlerror

// / stdc
type cint = int32 // for x64
type cuint = uint32
type char = int8
type uchar = uint8
type short = int16
type ushort = uint16
type float = float32
type double = float64
type charptr = *int8
type byteptr = *uint8
type wcharptr = *uint16

var Cmalloc func(sizet) voidptr
var Ccalloc func(sizet, sizet) voidptr
var Crealloc func(voidptr, sizet) voidptr
var cfreefn func(voidptr)
var Cmemset func(voidptr, cint, sizet) voidptr

func Cfree[T uintptr | voidptr](ptr T) { cfreefn(voidptr(ptr)) }

// var Cmemcpy func(voidptr, voidptr, sizet) voidptr
// var Cmemdup func(voidptr, sizet) voidptr

func init() {
	{
		fnadr, err := Dlsym(purego.RTLD_DEFAULT, "malloc")
		ErrPrint(err)
		purego.RegisterFunc(&Cmalloc, fnadr)
	}
	{
		fnadr, err := Dlsym(purego.RTLD_DEFAULT, "calloc")
		ErrPrint(err)
		purego.RegisterFunc(&Ccalloc, fnadr)
	}
	{
		fnadr, err := Dlsym(purego.RTLD_DEFAULT, "realloc")
		ErrPrint(err)
		purego.RegisterFunc(&Crealloc, fnadr)
	}
	{
		fnadr, err := Dlsym(purego.RTLD_DEFAULT, "free")
		ErrPrint(err)
		purego.RegisterFunc(&cfreefn, fnadr)
	}
	{
		fnadr, err := Dlsym(purego.RTLD_DEFAULT, "memset")
		ErrPrint(err)
		purego.RegisterFunc(&Cmemset, fnadr)
	}
}

func Cstrlen[T voidptr | charptr](ptr T) int { return cstrlen(ptr) }
func cstrlen[T voidptr | charptr](ptr T) int {
	if ptr == nil {
		return 0
	}
	v := gosliceref[char](voidptr(ptr), 1<<20)

	for i := 0; i < len(v); i++ {
		if v[i] == 0 {
			return i
		}
	}
	return 0
}
func cstrcpy[T voidptr | charptr](dst, src T) T {
	n := 1 << 30
	dv := gosliceref[byte](voidptr(dst), n)
	sv := gosliceref[byte](voidptr(src), n)
	for i := 0; i < n; i++ {
		dv[i] = sv[i]
		if sv[i] == 0 {
			break
		}
	}
	return dst
}
func cstrdup[T voidptr | charptr](ptr T) voidptr {
	n := cstrlen(ptr)
	dst := Cmalloc(n + 1)
	return Cmemcpy(dst, voidptr(ptr), n+1)
}
func cstrdupgc[T voidptr | charptr](ptr T) voidptr {
	n := cstrlen(ptr)
	dst := Mallocgc(n + 1)
	return Cmemcpy(dst, voidptr(ptr), n+1)
}

// cannot append
func gosliceref[T any](ptr voidptr, n sizet) []T {
	return *(*[]T)(voidptr(&GoSlice{ptr, n, n}))
}

// n index
func carrset[T any](ptr voidptr, idx sizet, v T) T {
	tv := gosliceref[T](ptr, idx+1)
	ov := tv[idx]
	tv[idx] = v
	return ov
}

func Cmemcpy[T voidptr | charptr](dst, src T, n sizet) T {
	dv := gosliceref[byte](voidptr(dst), n)
	sv := gosliceref[byte](voidptr(src), n)
	c := copy(dv, sv)
	if c != n {
	}
	return dst
}
func Cmemdup[T voidptr | charptr](ptr T, n sizet) voidptr {
	dst := Cmalloc(n + 1)
	return Cmemcpy(dst, voidptr(ptr), n)
}
func CmemdupAsstr[T voidptr | charptr](ptr T, n sizet) voidptr {
	dst := Cmalloc(n + 1)
	t := gosliceref[byte](dst, n+1)
	t[n] = 0
	return Cmemcpy(dst, voidptr(ptr), n)
}
func Cmemdupgc[T voidptr | charptr](ptr T, n sizet) voidptr {
	dst := Mallocgc(n + 1)
	return Cmemcpy(dst, voidptr(ptr), n)
}

func Mallocgc(n sizet) voidptr {
	if n < 0 {
		return nil
	}
	tmp := make([]byte, 0, n)
	rv := (*GoSlice)(voidptr(&tmp)).Data
	SetPin(rv, true)
	return rv
}

// same as officical go
func GoString[T voidptr | charptr](ptr T) (rv string) {
	sz := cstrlen(voidptr(ptr))
	tmp := Mallocgc(sz + 1)
	s := (*GoStringIn)(voidptr(&rv))
	s.Ptr = Cmemcpy(tmp, voidptr(ptr), sz)
	s.Len = sz
	return
}
func GoStringRef[T voidptr | charptr](ptr T) (rv string) {
	s := (*GoStringIn)(voidptr(&rv))
	s.Ptr = voidptr(ptr)
	s.Len = cstrlen(ptr)
	return
}

// same as officical go
func GoStringN[T voidptr | charptr](ptr T, len isize) (rv string) {
	dst := Mallocgc(len + 1)
	carrset(dst, len, byte(0))
	s := (*GoStringIn)(voidptr(&rv))
	s.Ptr = Cmemcpy(dst, voidptr(ptr), len)
	s.Len = len
	return
}
func GoStringNRef[T voidptr | charptr](ptr T, len isize) (rv string) {
	s := (*GoStringIn)(voidptr(&rv))
	s.Ptr = voidptr(ptr)
	s.Len = len
	return
}

// same as officical go
func CString(s string) voidptr {
	sz := len(s)
	dst := Cmalloc(sz + 1)
	carrset[byte](dst, sz, 0)
	sp := (*GoStringIn)(voidptr(&s))
	return Cmemcpy(dst, sp.Ptr, sz)
}

// maybe not null terminate when fake nil tail, but should be very rare
func CStringRef(s string) (voidptr, sizet) {
	p := &s
	if StrIsNilTail(p) {
		sp := (*GoStringIn)(voidptr(p))
		return (voidptr)(sp.Ptr), sp.Len
	}
	return CStringgc(*p), len(s) // still use go mem
	// sp := (*GoStringIn)(voidptr(&s))
	// return sp.Ptr, sp.Len
}
func CStringgc(s string) voidptr {
	sz := len(s)
	rp := Mallocgc(sz + 1)
	sp := (*GoStringIn)(voidptr(&s))
	return Cmemcpy(rp, sp.Ptr, sz)
}

func StringData(s string) voidptr    { return voidptr(unsafe.StringData(s)) }
func SliceData[T any](s []T) voidptr { return voidptr(unsafe.SliceData(s)) }

// note nocopy

// 更安全的refc字符串.
// 但是也还不够安全,有可能是临时变量,必须确保生命周期足够长.
// 在调用C函数的时候使用,在返回值的时候最好不用.
// 如果null结尾,则直接返回ref.
// 如果不是null结尾的,则返回clone版本.但是使用方需要在3秒内使用,否则内存会被自动翻译.
func StrtoRefc(s *string) voidptr {
	if StrIsNilTail(s) {
		sp := (*GoStringIn)(voidptr(s))
		return (voidptr)(sp.Ptr)
	}
	return CStringgc(*s) // still use go mem
	// underly C memory with auto free
	// s4c := CStringaf(*s)
	// return s4c
}

// 常量字符串失败
func StrChkNilTail(s *string) {
	v := StrIsNilTail(s)
	if !v {
		log.Println("gostr not nil tail", SubStr(*s, 32))
	}
}

func StrIsNilTail(s *string) bool {
	sp := (*GoStringIn)(voidptr(s))
	if false {
		// (*sp.ptr) = 0
	}
	idx := sp.Len
	// log.Println(idx, ch, sp.ptr, AddrAdd(voidptr(sp.ptr), 1))
	if sp.Ptr != nil {
		p1 := (*[1 << 16]byte)(voidptr(sp.Ptr))[: sp.Len+1 : sp.Len+1]
		// log.Println(1<<16, sp.len, reflect.TypeOf(p1), len(p1))
		return p1[idx] == 0
		// log.Println(p1)
	}
	return true
}

// note: 不能处理常量字符串。最大64KB
func StrNilTail(s *string) {
	sp := (*GoStringIn)(voidptr(s))
	if false {
		// (*sp.Ptr) = 0
	}
	idx := sp.Len
	// log.Println(idx, ch, sp.ptr, AddrAdd(voidptr(sp.ptr), 1))
	if sp.Ptr != nil {
		p1 := (*[1 << 20]byte)(voidptr(sp.Ptr))[: sp.Len+1 : sp.Len+1]
		// log.Println(1<<16, sp.len, reflect.TypeOf(p1), len(p1))
		// log.Println(p1[idx])
		if p1[idx] != 0 {
			p1[idx] = 0
		}
		// log.Println(p1)
	}
}

// note: 不能处理常量字符串。最大64KB
func StrAltChar(s *string, idx int, ch byte) {
	sp := (*GoStringIn)(voidptr(s))
	if false {
		// (*sp.Ptr) = 0
	}
	// log.Println(idx, ch, sp.ptr, AddrAdd(voidptr(sp.ptr), 1))
	if sp.Ptr != nil {
		p1 := (*[1 << 20]byte)(sp.Ptr)[:sp.Len:sp.Len]
		// log.Println(1<<16, sp.len, reflect.TypeOf(p1), len(p1))
		p1[idx] = ch
		// log.Println(p1)
	}
}
