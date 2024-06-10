////

//go:build usejni
// +build usejni

package cgopp

import (
	"fmt"
	"log"

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
// #cgo LDFLAGS: -ljli

const char*tstargs[] = {
	"-Xmx128m", "-Xms16m",
NULL,
};

*/
import "C"

// see Running native Android code in a Fyne app

const (
	JNI_VERSION_1_1 = C.JNI_VERSION_1_1
	JNI_VERSION_1_2 = C.JNI_VERSION_1_2
	JNI_VERSION_1_4 = C.JNI_VERSION_1_4
	JNI_VERSION_1_6 = C.JNI_VERSION_1_6
	// below android not have
	// JNI_VERSION_1_8 = C.JNI_VERSION_1_8
	// JNI_VERSION_9   = C.JNI_VERSION_9
	// JNI_VERSION_10  = C.JNI_VERSION_10
)

func RunOnJVM[FT func(vm, env, ctx uintptr) error |
	func() error | func()](fnx FT) error {
	switch fn := any(fnx).(type) {
	case func(vm, env, ctx uintptr) error:
		return mobinit.RunOnJVM(fn)
	case func() error:
		var fn2 = func(vm, env, ctx uintptr) error { return fn() }
		return mobinit.RunOnJVM(fn2)
	case func():
		var fn2 = func(vm, env, ctx uintptr) error { fn(); return nil }
		return mobinit.RunOnJVM(fn2)
	}
	return nil
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
func (j JNIEnv) Toc() *C.JNIEnv { return (*C.JNIEnv)(voidptr(j)) }
func (j JNIEnv) Tocuptr() cuptr { return (cuptr)(j) }
func (j JNIEnv) Toptr() voidptr { return voidptr(j) }
func (j JNIEnv) String() string { return fmt.Sprintf("%v", voidptr(j)) }

// cls: "Ljava/lang/String"
func (je JNIEnv) FindClass(cls string) voidptr {
	var cls2 = C.CString(cls)
	defer Cfree(cls2)
	// log.Println(je.Tocuptr(), je, cls)

	var e2 = (*C.JNIEnv)(voidptr(je))
	v := Litfficallg((*e2).FindClass, je, cls2)
	return v
}

func JniFindClass(cls string) voidptr { return jenv.FindClass(cls) }
func JniFindClassTS(cls string) voidptr {
	var rv voidptr
	err := RunOnJVM(func() error {
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
	rv := Litfficall(voidptr((*e2).NewStringUTF), je.Toptr(), voidptr(s4c))
	// rv := C.litfficall(voidptr((*e2).NewStringUTF), cint(2), je.Toptr(), voidptr(s4c), nil, nil, nil)
	// cannot call non-function (*e2).NewStringUTF (variable of type *[0]byte)
	// rv := (*e2).NewStringUTF(je.Toptr(), voidptr(s4c))
	if false { // compile ok
		Litfficallg((*e2).NewStringUTF, je, s4c)
	}
	// log.Println("fly333", rv, s4c)
	return rv
}

// argssig: "([Ljava/lang/String;)V"
func (je JNIEnv) GetStaticMethodID(clsid voidptr, s string, argssig string) voidptr {
	var s4c = C.CString(s)
	defer Cfree(s4c)
	var argssig4c = C.CString(argssig)
	defer Cfree(argssig4c)

	var e2 = (*C.JNIEnv)(voidptr(je))
	rv := Litfficall((voidptr((*e2).GetStaticMethodID)), je.Toptr(), clsid, voidptr(s4c), voidptr(&argssig4c))
	// rv := C.litfficall(voidptr((*e2).GetStaticMethodID), cint(4), je.Toptr(), clsid, voidptr(s4c), voidptr(&argssig4c), nil)
	// cannot call non-function (*e2).NewStringUTF (variable of type *[0]byte)
	// rv := (*e2).NewStringUTF(je.Toptr(), voidptr(s4c))
	// log.Println("fly333", rv, s4c)
	return rv
}
func (je JNIEnv) CallStaticVoidMethod(clsid, mthid voidptr) {
	var e2 = (*C.JNIEnv)(voidptr(je))
	rv := Litfficall(voidptr((*e2).CallStaticVoidMethod), je.Toptr(), clsid, mthid)
	// rv := C.litfficall(voidptr((*e2).CallStaticVoidMethod), cint(3), je.Toptr(), clsid, mthid, nil, nil)
	// cannot call non-function (*e2).NewStringUTF (variable of type *[0]byte)
	// rv := (*e2).NewStringUTF(je.Toptr(), voidptr(s4c))
	// log.Println("fly333", rv)
	gopp.GOUSED(rv)
}
func (je JNIEnv) ExceptionClear() {
	var e2 = (*C.JNIEnv)(voidptr(je))
	rv := Litfficallg((*e2).ExceptionClear, je)
	gopp.GOUSED(rv)
}

func GetCStringddd() {
	v := C.getCStringddd(0, 0, nil)
	log.Println(v)
}
func AndroidNameddd() {
	v := C.androidNameddd(0, 0, 0)
	log.Println(v)
}
