package gopp

import (
	"log"
	"sync"

	"github.com/ebitengine/purego"
)


var ffiloadonce sync.Once

// var ffi_prepare =

var ffisoh usize
func ffilib_init() {
	ffiloadonce.Do(func() {
		var err1 error
		libfiles := []string{"/usr/lib/libffi.so"}
		for _, libfile := range libfiles {
			if !FileExist2(libfile) {continue}
			h,err := purego.Dlopen(libfile, purego.RTLD_NOW)
			err1 = err
			if err == nil {
				ffisoh = h
				break
			}
		}
		if ffisoh == 0 {
			log.Fatalln("Somerr load ffi", err1, libfiles)
		}
		ffilib_load_typeobjs()
		ffilib_load_funcs()
	})
}

type FfiCifOrigin struct {
	Abi int32
	Nargs uint32
	Arg_types *ffi_type
	Rtype ffi_type
	Bytes uint32
	Flags uint32
	Padding usize
}

// with ffilib_init first
func ffi_prep_cif(cif *FfiCifOrigin, abi int32, nargs uint32, rtype ffi_type, atypes *ffi_type) int32 {
	ffilib_init()
	return ffi_prep_cif_(cif, abi, nargs, rtype, atypes)
}
func ffi_call(cif *FfiCifOrigin, fun voidptr, rvalue voidptr, avalues *voidptr) {
	ffilib_init()
	ffi_call_(cif, fun, rvalue, avalues)
}


var ffi_prep_cif_ func(*FfiCifOrigin, int32, uint32, ffi_type, *ffi_type) int32
var ffi_call_ func(*FfiCifOrigin, voidptr, voidptr, *voidptr)

func ffilib_load_funcs() {
	purego.RegisterLibFunc(&ffi_prep_cif_, ffisoh, "ffi_prep_cif")

	purego.RegisterLibFunc(&ffi_call_, ffisoh, "ffi_call")
}

func ffilib_load_typeobj(symname string) voidptr {
		rv,err:= purego.Dlsym(ffisoh, symname)
		ErrPrint(err, symname)
		return voidptr(rv)
}


func ffilib_load_typeobjs() {

    ffi_type_void    = ffilib_load_typeobj("ffi_type_void")
    ffi_type_uint8   = ffilib_load_typeobj("ffi_type_uint8")
    ffi_type_sint8   = ffilib_load_typeobj("ffi_type_sint8")
    ffi_type_uint16  = ffilib_load_typeobj("ffi_type_uint16")
    ffi_type_sint16  = ffilib_load_typeobj("ffi_type_sint16")
    ffi_type_uint32  = ffilib_load_typeobj("ffi_type_uint32")
    ffi_type_sint32  = ffilib_load_typeobj("ffi_type_sint32")
    ffi_type_uint64  = ffilib_load_typeobj("ffi_type_uint64")
    ffi_type_sint64  = ffilib_load_typeobj("ffi_type_sint64")
   	ffi_type_float   = ffilib_load_typeobj("ffi_type_float")
    ffi_type_double  = ffilib_load_typeobj("ffi_type_double")
    ffi_type_pointer  = ffilib_load_typeobj("ffi_type_pointer")
    ffi_type_longdouble = ffilib_load_typeobj("ffi_type_longdouble")

}


type ffi_type  = voidptr

/* These are defined in types.c.  */
var  ffi_type_void   ffi_type
var  ffi_type_uint8  ffi_type
var  ffi_type_sint8  ffi_type
var  ffi_type_uint16  ffi_type
var  ffi_type_sint16  ffi_type
var  ffi_type_uint32  ffi_type
var  ffi_type_sint32  ffi_type
var  ffi_type_uint64  ffi_type
var  ffi_type_sint64  ffi_type
var  ffi_type_float  ffi_type
var  ffi_type_double  ffi_type
var  ffi_type_pointer  ffi_type
var  ffi_type_longdouble  ffi_type

// #ifdef FFI_TARGET_HAS_COMPLEX_TYPE
// FFI_EXTERN ffi_type ffi_type_complex_float;
// FFI_EXTERN ffi_type ffi_type_complex_double;
// FFI_EXTERN ffi_type ffi_type_complex_longdouble;
// #endif


/* If these change, update src/mips/ffitarget.h. */
const  FFI_TYPE_VOID     =  0
const  FFI_TYPE_INT      =  1
const  FFI_TYPE_FLOAT    =  2
const  FFI_TYPE_DOUBLE   =  3
// #if 1
const FFI_TYPE_LONGDOUBLE =  4
// #else
// #define FFI_TYPE_LONGDOUBLE FFI_TYPE_DOUBLE
// #endif
const  FFI_TYPE_UINT8    =   5
const  FFI_TYPE_SINT8    =   6
const  FFI_TYPE_UINT16   =   7
const  FFI_TYPE_SINT16   =   8
const  FFI_TYPE_UINT32   =   9
const  FFI_TYPE_SINT32   =   10
const  FFI_TYPE_UINT64   =   11
const  FFI_TYPE_SINT64   =   12
const  FFI_TYPE_STRUCT   =   13
const  FFI_TYPE_POINTER  =   14
const  FFI_TYPE_COMPLEX  =   15

/* This should always refer to the last type code (for sanity checks).  */
const FFI_TYPE_LAST      = FFI_TYPE_COMPLEX
