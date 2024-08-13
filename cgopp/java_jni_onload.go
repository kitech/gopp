//go:build usejni
// +build usejni

package cgopp

import (
	"log"

	"github.com/kitech/gopp"
	mobinit "github.com/kitech/gopp/internal/mobileinit"
	cmap "github.com/orcaman/concurrent-map/v2"
)

/*
// #include <jni.h>

// #include "java_jni.h"

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
	jnimemfninit(vm, jenvmt)
	jnimtclsesinit(jenvmt)

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

func jnimtclsesinit(je JNIEnv) {
	names := jnimtclses.Keys()

	classLoaderClass := je.FindClass("java/lang/ClassLoader")
	log.Println(classLoaderClass)

	for _, cls := range names {
		rv := je.FindClass(cls)
		if rv != 0 {
			// jnimtclses.Set(cls, rv)
			rvcls := je.GetObjectClass(rv)
			getClassLoaderMethod := je.GetMethodID(rvcls, "getClassLoader",
				"()Ljava/lang/ClassLoader;")
			gClassLoader := JNIEnvCallMethod[usize](je, rvcls, getClassLoaderMethod)
			gFindClassMethod := je.GetMethodID(classLoaderClass, "findClass",
				"(Ljava/lang/String;)Ljava/lang/Class;")
			log.Println(cls, gClassLoader, gFindClassMethod)

			gClassLoader2 := je.NewGlobalRef(gClassLoader)
			log.Println(cls, gClassLoader, gFindClassMethod)
			// gFindClassMethod2 := je.NewGlobalRef(gFindClassMethod)
			log.Println(cls, gClassLoader, gFindClassMethod)
			rv2 := je.NewGlobalRef(rv)
			log.Println(cls, gClassLoader2, gFindClassMethod, rv2)
			jnimtclses.Set(cls, []usize{gClassLoader2, gFindClassMethod, rv2})
			// 这种方式还是不行，jobject is an invalid local reference: 0x59 (deleted reference at index 5 in a table of size 1)

		}
		log.Println("mt.FindClass", cls, rv, jnimtclses.Count(), gopp.Retn(jnimtclses.Get(cls)))
	}

	//
}

// with cache version
func (je JNIEnv) FindClass2(cls string) usize {
	if meta, ok := jnimtclses.Get(cls); ok && meta != nil {
		log.Println(cls, meta, ok)
		// clsid := JNIEnvCallMethod[usize](je, meta[0], meta[1], cls)
		// log.Println(cls, meta, ok, clsid)
		// if clsid != 0 {
		// 	return clsid
		// }
		if meta[2] != 0 {
			return meta[2]
		}
	}
	// if clsid, ok := jnimtclses.Get(cls); ok && clsid != 0 {
	// 	return clsid
	// }
	log.Println("fallback origin FindClass", cls)
	clsid := je.FindClass(cls)
	return clsid
}

// 还是不行
// jclass is an invalid local reference: 0x25 (deleted reference at index 2 in a table of size 0)
// 正确缓存java object的方法？？？
// https://stackoverflow.com/questions/13263340/findclass-from-any-thread-in-android-jni
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
