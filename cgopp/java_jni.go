package cgopp

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/kitech/gopp"
	mobinit "github.com/kitech/gopp/internal/mobileinit"
)

/*
#include <jni.h>
#include <stdlib.h>
#include <string.h>

#include "java_jni.h"

// or set CGO_CFLAGS
// #cgo CFLAGS: -I/path/to/jdkhome/include

// or set CGO_LDFLAGS
// #cgo LDFLAGS: -ljava -ljvm
#cgo LDFLAGS: -ljli

const char*tstargs[] = {
	"-Xmx128m", "-Xsx16m",
NULL,
};

void* litfficall(void* fnptr, int argc, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4) {
	printf("fnptr=%p,argc=%d,arg0=%p\n", fnptr, argc, arg0);
	typedef void* (*fnargc0)() ;
	fnargc0 fn0 = fnptr;
	typedef void* (*fnargc1)(void*) ;
	fnargc1 fn1 = fnptr;
	typedef void* (*fnargc2)(void*, void*) ;
	fnargc2 fn2 = fnptr;
	typedef void* (*fnargc3)(void*, void*, void*) ;
	fnargc3 fn3 = fnptr;
	typedef void* (*fnargc4)(void*, void*, void*, void*) ;
	fnargc4 fn4 = fnptr;

	void* retptr = 0;
	switch (argc) {
	case 0:
		retptr = fn0();
		break;
	case 1:
		retptr = fn1(arg0);
		break;
	case 2:
		retptr = fn2(arg0, arg1);
		break;
	case 3:
		retptr = fn3(arg0, arg1, arg2);
		break;
	case 4:
		retptr = fn4(arg0, arg1, arg2, arg3);
		break;
	}

	return retptr;
}

*/
import "C"

// see Running native Android code in a Fyne app

// extern /*jclass*/void* find_class(/*JNIEnv**/uintptr_t jni_env, const char* class_name);
// extern const char* getCString(uintptr_t jni_env, uintptr_t ctx, /*jstring*/ void* str);
// extern const char* androidName(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);

const (
	JNI_VERSION_1_1 = C.JNI_VERSION_1_1
	JNI_VERSION_1_2 = C.JNI_VERSION_1_2
	JNI_VERSION_1_4 = C.JNI_VERSION_1_4
	JNI_VERSION_1_6 = C.JNI_VERSION_1_6
	JNI_VERSION_1_8 = C.JNI_VERSION_1_8
	JNI_VERSION_9   = C.JNI_VERSION_9
	JNI_VERSION_10  = C.JNI_VERSION_10
)

func RunOnJVM(fn func(vm, env, ctx uintptr) error) error {
	return mobinit.RunOnJVM(fn)
}
func RunOnJVM2(fn func() error) error {
	var fn2 = func(vm, env, ctx uintptr) error { return fn() }
	return mobinit.RunOnJVM(fn2)
}

// todo WIP crashing
func JavaExe(clsname, funcname string, args ...string) int {
	// var args2 = []string{"-Xsx32m", "-Xmx128m"}
	var args2 = args
	var argspp = StringSliceToCCharPP(args2)
	var args4c = (**C.char)(argspp)
	// var args4c = &C.tstargs[0]
	// rv := C.create_java_exe(C.CString("java/lang/String"), C.CString("main"), args4c)
	rv := C.create_java_exe(C.CString(clsname), C.CString(funcname), args4c)
	return int(rv)
}

////

func (j JavaVM) Env() JNIEnv {
	var env JNIEnv
	envx := C.getjavaenvbyjavavm(voidptr(j))
	env = JNIEnv(envx)
	return env
}
func (j JNIEnv) Toc() *C.JNIEnv         { return (*C.JNIEnv)(voidptr(j)) }
func (j JNIEnv) Tocuptr() C.uintptr_t   { return (C.uintptr_t)(j) }
func (j JNIEnv) Toptr() voidptr         { return voidptr(j) }
func (j JNIEnv) Toptr2() unsafe.Pointer { return voidptr(j) }
func (j JNIEnv) String() string         { return fmt.Sprintf("%v", voidptr(j)) }

// cls: "Ljava/lang/String"
func (je JNIEnv) FindClass(cls string) voidptr {
	var cls2 = C.CString(cls)
	defer Cfree(cls2)
	log.Println(je.Tocuptr(), je, cls)
	v := C.find_class(je.Tocuptr(), cls2)
	return v
}

func JniFindClass(cls string) voidptr { return jenv.FindClass(cls) }
func JniFindClassTS(cls string) voidptr {
	var rv voidptr
	err := RunOnJVM2(func() error {
		rv = JniFindClass(cls)
		return nil
	})
	gopp.ErrPrint(err, cls)
	return rv
}

func (je JNIEnv) NewStrUTF(s string) voidptr {
	var s4c = C.CString(s)
	defer Cfree(s4c)

	var e2 = (*C.JNIEnv)(voidptr(je))
	rv := C.litfficall(voidptr((*e2).NewStringUTF), cint(2), je.Toptr(), voidptr(s4c), nil, nil, nil)
	// cannot call non-function (*e2).NewStringUTF (variable of type *[0]byte)
	// rv := (*e2).NewStringUTF(je.Toptr(), voidptr(s4c))
	log.Println("fly333", rv, s4c)
	return rv
}

// argssig: "([Ljava/lang/String;)V"
func (je JNIEnv) GetStaticMethodID(clsid voidptr, s string, argssig string) voidptr {
	var s4c = C.CString(s)
	defer Cfree(s4c)
	var argssig4c = C.CString(argssig)
	defer Cfree(argssig4c)

	var e2 = (*C.JNIEnv)(voidptr(je))
	rv := C.litfficall(voidptr((*e2).GetStaticMethodID), cint(4), je.Toptr(), clsid, voidptr(s4c), voidptr(&argssig4c), nil)
	// cannot call non-function (*e2).NewStringUTF (variable of type *[0]byte)
	// rv := (*e2).NewStringUTF(je.Toptr(), voidptr(s4c))
	log.Println("fly333", rv, s4c)
	return rv
}
func (je JNIEnv) CallStaticVoidMethod(clsid, mthid voidptr) {
	var e2 = (*C.JNIEnv)(voidptr(je))
	rv := C.litfficall(voidptr((*e2).CallStaticVoidMethod), cint(3), je.Toptr(), clsid, mthid, nil, nil)
	// cannot call non-function (*e2).NewStringUTF (variable of type *[0]byte)
	// rv := (*e2).NewStringUTF(je.Toptr(), voidptr(s4c))
	log.Println("fly333", rv)
}

func GetCString() {
	v := C.getCString(0, 0, nil)
	log.Println(v)
}
func AndroidName() {
	v := C.androidName(0, 0, 0)
	log.Println(v)
}
