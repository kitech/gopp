package cgopp

import (
	"fmt"
	"log"
	"time"
	"unsafe"

	mobinit "github.com/kitech/gopp/internal/mobileinit"
)

/*
#include <jni.h>

#include "java_jni.h"

#cgo CFLAGS: -I/nix/store/rflj4qrjp5km8kqfwh2s70s64y4d904v-zulu-ca-jdk-17.0.10/include
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
	log.Println(je.Tocuptr(), je)
	v := C.find_class(je.Tocuptr(), cls2)
	return v
}
func RunOnJVM(fn func(vm, env, ctx uintptr) error) error {
	return mobinit.RunOnJVM(fn)
}
func RunOnJVM2(fn func() error) error {
	var fn2 = func(vm, env, ctx uintptr) error { return fn() }
	return mobinit.RunOnJVM(fn2)
}

// cls format: L
func FindClass(cls string) {
	log.Println("hehehe", cls)
	time.Sleep(3 * time.Second)
	// v := C.find_class(0, C.CString(""))
	v := C.find_class(0, C.CString(cls))
	log.Println(v)
}
func GetCString() {
	v := C.getCString(0, 0, nil)
	log.Println(v)
}
func AndroidName() {
	v := C.androidName(0, 0, 0)
	log.Println(v)
}
