// //goddd:build usejni
// // +-buildddd usejni

package cgopp

import (
	"log"

	"github.com/kitech/gopp"
	cmap "github.com/orcaman/concurrent-map/v2"
)

/*
// #include <jni.h>

// #include "java_jni.h"

// #cgo CFLAGS: -I/nix/store/rflj4qrjp5km8kqfwh2s70s64y4d904v-zulu-ca-jdk-17.0.10/include
*/
import "C"

type JavaVM usize
type JNIEnv usize

var jvm JavaVM    // C.JavaVM*
var jenvmt JNIEnv // C.JNIEnv*
var Jvm JavaVM

// var Jenvmt JNIEnv // main jvm thread env
var jvmmttid usize

func JNIIsLoad() bool  { return jvm != 0 }
func JVMmtTid() usize  { return jvmmttid }
func JNIIsJvmth() bool { return jvmmttid == MyTid() }
func JNIEnvmt() JNIEnv { return jenvmt }

// see jni.h JNI_OnLoad
//
//export JNI_OnLoad
func JNI_OnLoad(vm JavaVM, reserved usize) int {
	// log.Println("hello", vm, reserved)
	mobinitSetCurrentContext(voidptr(vm), 0)

	jvmmttid = MyTid()
	log.Printf("cgopp.JNI_OnLoad %v, %v\n", voidptr(vm), MyTid())
	jvm, Jvm = vm, vm
	jenvmt = jvm.Env()
	jnimemfninit(vm, jenvmt)
	jnimtclsesinit(jenvmt)

	log.Println("jvm", voidptr(jvm), "jenvmt", voidptr(jenvmt), "tid", jvmmttid)
	gopp.NilPrint(jenvmt, "some error occus", voidptr(vm))

	JNI_OnLoad_Callback()
	return JNI_VERSION_1_6
}

// 一般执行不到这个回调的

//export JNI_OnUnload
func JNI_OnUnload(vm JavaVM, reserved usize) {
	log.Printf("cgopp.JNI_OnUnload %v\n", voidptr(vm))
}

// set in java_jnicw.go
var jnimemfninit func(jvm JavaVM, je JNIEnv)
var mobinitSetCurrentContext func(voidptr, usize)
var mobinitRunOnJVM func(func(vm, env, ctx usize) error) error
var JNI_VERSION_1_6 int // dont change
var getJavaVMEnvByc func(jvmx usize) usize

func jnimtclsesinit(je JNIEnv) {
	names := jnimtclses.Keys()

	for _, cls := range names {
		rv := je.FindClass(cls)
		if rv != 0 {
			rv2 := je.NewGlobalRef(rv)
			jnimtclses.Set(cls, []usize{rv2, rv2, rv2})
		} else {
			log.Println("java class not found", cls, rv)
		}
	}

	//
}

// with cache version
func (je JNIEnv) FindClass2(cls string) usize {
	if meta, ok := jnimtclses.Get(cls); ok && meta != nil {
		if meta[2] != 0 {
			return meta[2]
		}
	}
	log.Println("fallback origin FindClass", cls)
	clsid := je.FindClass(cls)
	return clsid
}

// 还是不行
// jclass is an invalid local reference: 0x25 (deleted reference at index 2 in a table of size 0)
var jnimtclses = cmap.New[[]usize]()

// var classloaderobj

func JNIRegistFindClasses(clses ...string) {
	mtcs := jnimtclses
	for _, cls := range clses {
		if _, ok := mtcs.Get(cls); !ok {
			mtcs.Set(cls, nil)
		}
	}
}
