package cgopp

import (
	"log"

	"github.com/kitech/gopp"
	mobinit "github.com/kitech/gopp/internal/mobileinit"
)

/*
// #include <jni.h>

#include "java_jni.h"

// #cgo CFLAGS: -I/nix/store/rflj4qrjp5km8kqfwh2s70s64y4d904v-zulu-ca-jdk-17.0.10/include
*/
import "C"

type JavaVM uintptr
type JNIEnv uintptr

var jvm JavaVM  // C.JavaVM*
var jenv JNIEnv // C.JNIEnv*
var Jvm JavaVM
var Jenv JNIEnv
var jvmtid uint64

func JNIIsLoad() bool { return jvm != 0 }

// see jni.h JNI_OnLoad

//export JNI_OnLoad
func JNI_OnLoad(vm JavaVM, x uintptr) int {
	// log.Println("hello", vm, x)
	mobinit.SetCurrentContext(voidptr(vm), 0)

	jvmtid = MyTid()
	log.Printf("cgopp.JNI_OnLoad %v, %v\n", voidptr(vm), MyTid())
	jvm = vm
	Jvm = vm
	jenv = jvm.Env()
	Jenv = jenv
	log.Println("jvm", voidptr(jvm), "jenv", voidptr(jenv))
	gopp.NilPrint(jenv, "some error occus", voidptr(vm))

	// Jenv.FindClass("java/lang/String")
	return JNI_VERSION_1_8
}

// 一般执行不到这个回调的

//export JNI_OnUnload
func JNI_OnUnload(vm JavaVM, x uintptr) {
	log.Printf("cgopp.JNI_OnUnload %v\n", voidptr(vm))
}
