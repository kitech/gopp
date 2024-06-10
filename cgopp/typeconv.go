package cgopp

/*
#include <string.h>
#include <stdlib.h>

static void carr_set_item(char** pp, int idx, char*p)
{ pp[idx] = p; }
static char* carr_get_item(char** pp, int idx)
{ return pp[idx]; }

*/
import "C"
import (
	"math"
	"unsafe"

	"github.com/kitech/gopp"
)

func StringSliceToCCharPP(ss []string) unsafe.Pointer {
	var tp *C.char
	p := C.calloc(C.size_t(len(ss)+1), C.size_t(unsafe.Sizeof(tp)))
	var pp **C.char = (**C.char)(p)

	for i, _ := range ss {
		s := C.CString(ss[i])
		C.carr_set_item(pp, C.int(i), s)
		C.carr_set_item(pp, C.int(i+1), nil)
	}

	return p
}
func CCharPPToStringSlice(charpp unsafe.Pointer) []string {
	ss := []string{}
	var pp **C.char = (**C.char)(charpp)
	for i := 0; i < math.MaxInt32; i++ {
		p := C.carr_get_item(pp, C.int(i))
		if p == nil {
			break
		}
		ss = append(ss, C.GoString(p))
	}
	return ss
}

// x64
// note: C.int != go int
type Cint C.int
type Cgoint int32
type Clong C.long
type Cgolong int64

// x32
// note: C.int == go int
/*
type Cint C.int
type Cgoint int32
type Clong C.long
type Cgolong = int64
*/

// => char**
type CStrArr struct {
	carr  unsafe.Pointer
	calen int
	garr  []*string
}

func CStrArrFromu8(arr **uint8, n int) *CStrArr {
	return CStrArrFromp(unsafe.Pointer(arr), n)
}
func CStrArrFromc8(arr **int8, n int) *CStrArr {
	return CStrArrFromp(unsafe.Pointer(arr), n)
}

// must a (u)char**
func CStrArrFromp(arr unsafe.Pointer, n int) *CStrArr {
	this := &CStrArr{}
	this.carr = arr
	return this
}

func (this *CStrArr) ToGo() (rets []string) {
	for i := 0; i < this.calen; i++ {
		ep := unsafe.Pointer(uintptr(this.carr) + unsafe.Sizeof((uintptr(0)))*uintptr(i))
		e := C.GoString((*C.char)(ep))
		rets = append(rets, e)
	}
	return
}

func CStrArrFromStrs(arr []string) *CStrArr {
	this := &CStrArr{}
	for _, e := range arr {
		t := e + "\x00"
		this.garr = append(this.garr, &t)
	}
	this.garr = append(this.garr, nil)
	return this
}

func (this *CStrArr) ToC() unsafe.Pointer {
	return (unsafe.Pointer)(&this.garr[0])
}

func (this *CStrArr) Append(s string) {
	if this.garr == nil {
		// think as from c
		strs := this.ToGo()
		tarr := CStrArrFromStrs(strs)
		this.garr = tarr.garr
	}
	e := s + "\x00"
	this.garr = append(this.garr, &e)
}

func GoStrArr2c(arr []string) uintptr {
	if len(arr) == 0 {
		return 0
	}

	pv := make([]unsafe.Pointer, len(arr)+1)
	for i, v := range arr {
		pv[i] = unsafe.Pointer(C.CString(v))
	}
	sz := int(unsafe.Sizeof(uintptr(0))) * (len(arr) + 1)
	rv := C.calloc(1, C.ulong(sz))
	C.memcpy(rv, unsafe.Pointer(&pv[0]), C.ulong(sz))
	return uintptr(rv)
}

func AddrAdd(addr voidptr, n usize) voidptr {
	rv := voidptr(usize(addr) + n)
	return rv
}

type gostrin struct {
	ptr *C.char
	len usize
	cap usize
}

// note nocopy
func StrtoCharpRef(s *string) charptr {
	gopp.FalsePrint(StrIsNilTail(s), "not safe case, gostring not null terminated")
	sp := (*gostrin)(voidptr(s))
	return (charptr)(sp.ptr)
}

// note nocopy
func StrtoVptrRef(s *string) vptr {
	gopp.FalsePrint(StrIsNilTail(s), "not safe case, gostring not null terminated")
	sp := (*gostrin)(voidptr(s))
	return (vptr)(sp.ptr)
}

func StrIsNilTail(s *string) bool {
	sp := (*gostrin)(voidptr(s))
	if false {
		(*sp.ptr) = 0
	}
	idx := sp.len
	// log.Println(idx, ch, sp.ptr, AddrAdd(voidptr(sp.ptr), 1))
	if sp.ptr != nil {
		p1 := (*[1 << 16]byte)(voidptr(sp.ptr))[: sp.len+1 : sp.len+1]
		// log.Println(1<<16, sp.len, reflect.TypeOf(p1), len(p1))
		return p1[idx] == 0
		// log.Println(p1)
	}
	return true
}

// note: 不能处理常量字符串。最大64KB
func StrNilTail(s *string) {
	sp := (*gostrin)(voidptr(s))
	if false {
		(*sp.ptr) = 0
	}
	idx := sp.len
	// log.Println(idx, ch, sp.ptr, AddrAdd(voidptr(sp.ptr), 1))
	if sp.ptr != nil {
		p1 := (*[1 << 20]byte)(voidptr(sp.ptr))[: sp.len+1 : sp.len+1]
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
	sp := (*gostrin)(voidptr(s))
	if false {
		(*sp.ptr) = 0
	}
	// log.Println(idx, ch, sp.ptr, AddrAdd(voidptr(sp.ptr), 1))
	if sp.ptr != nil {
		p1 := (*[1 << 20]byte)(voidptr(sp.ptr))[:sp.len:sp.len]
		// log.Println(1<<16, sp.len, reflect.TypeOf(p1), len(p1))
		p1[idx] = ch
		// log.Println(p1)
	}
}

func GoString(ptr voidptr) string {
	return C.GoString((*C.char)(ptr))
}
func GoStringN(ptr voidptr, len usize) string {
	return C.GoStringN((*C.char)(ptr), (C.int)(len))
}
func CString(s string) voidptr {
	return voidptr(C.CString(s))
}
