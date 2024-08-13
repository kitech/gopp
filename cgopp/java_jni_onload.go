//go:build usejni
// +build usejni

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

var jvm JavaVM    // C.JavaVM*
var jenvmt JNIEnv // C.JNIEnv*
var Jvm JavaVM

// var Jenvmt JNIEnv // main jvm thread env
var jvmmttid uintptr

func JNIIsLoad() bool   { return jvm != 0 }
func JVMmtTid() uintptr { return jvmmttid }
func JNIIsJvmth() bool  { return jvmmttid == MyTid() }
func JNIEnvmt() JNIEnv  { return jenvmt }

// see jni.h JNI_OnLoad
//
//export JNI_OnLoad
func JNI_OnLoad(vm JavaVM, reserved uintptr) int {
	// log.Println("hello", vm, reserved)
	mobinit.SetCurrentContext(voidptr(vm), 0)

	jvmmttid = MyTid()
	log.Printf("cgopp.JNI_OnLoad %v, %v\n", voidptr(vm), MyTid())
	jvm, Jvm = vm, vm
	jenvmt = jvm.Env()

	log.Println("jvm", voidptr(jvm), "jenvmt", voidptr(jenvmt), "tid", jvmmttid)
	gopp.NilPrint(jenvmt, "some error occus", voidptr(vm))

	JNI_OnLoad_Callback()
	return JNI_VERSION_1_6
}

// 一般执行不到这个回调的

//export JNI_OnUnload
func JNI_OnUnload(vm JavaVM, reserved uintptr) {
	log.Printf("cgopp.JNI_OnUnload %v\n", voidptr(vm))
}
