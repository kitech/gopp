////

//go:build usejni
// +build usejni

package cgopp

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

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

// 似乎这个不管用，是因为忽略了参数中的env，这个env是和当前线程绑定的
// 在不能保证线程是否为JVM线程时，使用参数中的JNIEnv变量
func RunOnJVM[FT func(vm, env, ctx uintptr) error |
	func(env usize) error | func(env usize)](fnx FT) error {
	switch fn := any(fnx).(type) {
	case func(vm, env, ctx uintptr) error:
		return mobinit.RunOnJVM(fn)
	case func(usize) error:
		var fn2 = func(vm, env, ctx uintptr) error { return fn(env) }
		return mobinit.RunOnJVM(fn2)
	case func(usize):
		var fn2 = func(vm, env, ctx uintptr) error { fn(env); return nil }
		return mobinit.RunOnJVM(fn2)
	default:
	}
	return nil
}

// 这个check是在go空间执行的不准确
func JNIThreadCheck(label ...any) bool {
	bv := JNIIsJvmth()
	if !bv {
		log.Println(label, "not on jvm thread:", jvmmttid, "but:", MyTid())
	}
	return bv
}

// todo WIP crashing
func JavaExe(clsname, funcname string, args ...string) int {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// var args2 = []string{"-Xsx32m", "-Xmx128m"}
	var args2 = args
	var argspp = StringSliceToCCharPP(args2)
	var args4c = (*charptr)(argspp)
	// var args4c = &C.tstargs[0]
	// rv := C.create_java_exe(C.CString("java/lang/String"), C.CString("main"), args4c)
	rv := C.create_java_exe(C.CString(clsname), C.CString(funcname), args4c)
	return int(rv)
}

// /// 这种方式封装可以分离部分依赖C的代码
func (jvm JavaVM) fnGetEnv() voidptr {
	return voidptr((*(*C.JavaVM)(voidptr(jvm))).GetEnv)
}
func (jvm JavaVM) fnAttachCurrentThread() voidptr {
	return voidptr((*(*C.JavaVM)(voidptr(jvm))).AttachCurrentThread)
}
func (jvm JavaVM) fnDetachCurrentThread() voidptr {
	return voidptr((*(*C.JavaVM)(voidptr(jvm))).DetachCurrentThread)
}
func (jvm JavaVM) fnAttachCurrentThreadAsDaemon() voidptr {
	return voidptr((*(*C.JavaVM)(voidptr(jvm))).AttachCurrentThreadAsDaemon)
}
func (jvm JavaVM) fnDestroyJavaVM() voidptr {
	return voidptr((*(*C.JavaVM)(voidptr(jvm))).DestroyJavaVM)
}

func (je JNIEnv) fnGetVersion() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetVersion)
}
func (je JNIEnv) fnFindClass() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).FindClass)
}
func (je JNIEnv) fnGetStaticMethodID() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetStaticMethodID)
}
func (je JNIEnv) fnCallStaticVoidMethod() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticVoidMethod)
}
func (je JNIEnv) fnCallStaticObjectMethod() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticObjectMethod)
}
func (je JNIEnv) fnCallStaticIntMethod() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticIntMethod)
}
func (je JNIEnv) fnCallStaticLongMethod() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticLongMethod)
}
func (je JNIEnv) fnCallStaticFloatMethod() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticFloatMethod)
}
func (je JNIEnv) fnCallStaticDoubleMethod() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticDoubleMethod)
}
func (je JNIEnv) fnCallStaticBooleanMethod() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticBooleanMethod)
}

func (je JNIEnv) fnNewStringUTF() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).NewStringUTF)
}
func (je JNIEnv) fnGetStringUTFLength() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetStringUTFLength)
}
func (je JNIEnv) fnGetStringUTFChars() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetStringUTFChars)
}
func (je JNIEnv) fnReleaseStringUTFChars() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).ReleaseStringUTFChars)
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

func (je JNIEnv) GetVersion() int {
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = (voidptr((*e2).GetVersion))
	rv := Litfficall(fnptr, je.Toptr())
	// log.Println(rv)
	return int(usize(rv))
}

// 尝试L前缀，尝试把.替换/
// 为什么有时需要前缀L，有时不需要，像Main
// 使用 javap -s -p Main 查看函数签名信息
// cls: "Ljava/lang/String"
func (je JNIEnv) FindClass(cls string) voidptr {
	if strings.Count(cls, ".") > 0 {
		cls = strings.ReplaceAll(cls, ".", "/")
	}
	var cls4c = CStringgc(cls)
	// log.Println(je.Tocuptr(), je, rc)

	var e2 = (*C.JNIEnv)(voidptr(je))
	// v := Litfficallg((*e2).FindClass, je.Toptr(), cls4c)
	v := FfiCall[voidptr]((*e2).FindClass, je.Toptr(), cls4c)
	if v != nil {
		return v
	}

	return nil
}

// s: ???
// argssig: "([Ljava/lang/String;)V"
// ()V
// 有的结尾有分号，有的没有
// 使用 javap -s -p Main 查看函数签名信息
func (je JNIEnv) GetStaticMethodID(clsid voidptr, mthname string, argssig string) voidptr {
	// log.Println(clsid, s, argssig)
	var s4c = CStringgc(mthname)
	var argssig4c = CStringgc(argssig)

	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).GetStaticMethodID)
	// rv := Litfficall(fnptr, je.Toptr(), clsid, s4c, argssig4c)
	rv := FfiCall[voidptr](fnptr, je.Toptr(), clsid, s4c, argssig4c)
	je.ExceptionCheck()
	return rv
}

// support string/int/todo
// 这个应该使用的少
func (je JNIEnv) CallStaticVoidMethod(clsid, mthid voidptr, args ...any) {
	var argc = len(args) + 3
	var argv [9]any
	argv[0] = je.Toptr()
	argv[1] = clsid
	argv[2] = mthid

	for i, argx := range args {
		switch arg := argx.(type) {
		case string:
			tv := je.NewStringUTF(arg)
			argv[i+3] = tv
		case int:
			argv[i+3] = arg
		default:
			log.Println("Nocat", reflect.TypeOf(argx), argx)
		}
	}

	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = (*e2).CallStaticVoidMethod
	// rv := Litfficallg(fnptr, je.Toptr(), clsid, mthid, a04c)
	rv := FfiCall[int](fnptr, argv[:argc]...)
	gopp.GOUSED(rv)
}

// jni没有查看类型的函数！！！
type Jany uintptr

func (me Jany) Tostr() string {
	return ""
}

// go的方法不能带模板类型!!!
// 不支持float/double类型参数和返回值
// support args: string/int/todo
// support ret: string/int/todo
func JNIEnvCallStaticMethod[RTY any](je JNIEnv, clsid, mthid voidptr, args ...any) (rvx RTY) {

	var argvarr [9]any
	var argc = len(args) + 3
	var argv = argvarr[:argc]
	argv[0] = je.Toptr()
	argv[1] = clsid
	argv[2] = mthid

	for i, argx := range args {
		switch arg := argx.(type) {
		case string:
			tv := je.NewStringUTF(arg)
			argv[i+3] = tv
		case int:
			tv := arg
			// argv[i+3] = voidptr(usize(tv))
			argv[i+3] = tv // fixsome gostack pointer check
		default:
			log.Println("Nocat", reflect.TypeOf(argx), argx)
		}
	}

	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr *[0]byte
	switch any(rvx).(type) {
	case string:
		fnptr = (*e2).CallStaticObjectMethod
		rvp := FfiCall[voidptr](fnptr, argv...)
		rv2 := je.GetStringUTFChars(rvp)
		rvx = any(rv2).(RTY)
	case int:
		log.Println(rvx, clsid, mthid, len(argv), argv)
		fnptr = (*e2).CallStaticIntMethod
		rvx = FfiCall[RTY](fnptr, argv...)
	case float64:
		fnptr = (*e2).CallStaticDoubleMethod
		rvx = FfiCall[RTY](fnptr, argv...)
	case Void:
		// log.Println(rvx, clsid, mthid, len(argv), argv)
		fnptr = (*e2).CallStaticVoidMethod
		FfiCall[RTY](fnptr, argv...)
	default:
		log.Println("Nocat", reflect.TypeOf(any(rvx)))
	}
	gopp.GOUSED(e2, fnptr)

	return
}

// https://stackoverflow.com/questions/40004522/how-to-get-values-from-jobject-in-c-using-jni

func (je JNIEnv) GetObjectClass(obj voidptr) voidptr {
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).GetObjectClass)
	rv := Litfficall(fnptr, je.Toptr(), obj)
	return rv
}
func (je JNIEnv) GetFieldID(clsobj voidptr, a0, a1 string) voidptr {
	a04c := CStringgc(a0)
	a14c := CStringgc(a1)

	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).GetFieldID)
	rv := Litfficall(fnptr, je.Toptr(), clsobj, a04c, a14c)
	return rv
}
func (je JNIEnv) GetIntField(clsobj voidptr, fidobj voidptr) int {
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).GetIntField)
	rv := Litfficall(fnptr, je.Toptr(), fidobj)
	return int(usize(rv))
}

func (je JNIEnv) ExceptionClear() {
	var e2 = (*C.JNIEnv)(voidptr(je))
	rv := Litfficallg((*e2).ExceptionClear, je)
	gopp.GOUSED(rv)
}
func (je JNIEnv) ExceptionCheck() bool {
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = (voidptr((*e2).ExceptionCheck))
	rv := Litfficallg(fnptr, je.Toptr())
	if rv != nil {
		log.Println("Some error", rv, usize(rv), MyTid())
		je.ExceptionDescribe()
		JNIThreadCheck()
		panic(rv)
	}
	return rv != nil
}
func (je JNIEnv) ExceptionDescribe() {
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = (voidptr((*e2).ExceptionDescribe))
	rv := Litfficallg(fnptr, je.Toptr())
	if rv != nil {
		log.Println("Some error", rv, usize(rv))
	}
}

// 这个生成的object不需要自己释放
func (je JNIEnv) NewStringUTF(s string) voidptr {
	s4c := CStringgc(s)

	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).NewStringUTF)
	rv := Litfficallg(fnptr, je.Toptr(), s4c)
	return rv
}

func (je JNIEnv) ReleaseStringUTFChars(strx voidptr, utfx voidptr) {
	// JNI DETECTED ERROR IN APPLICATION: non-nullable argument was NULL
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).ReleaseStringUTFChars)
	rv := Litfficallg(fnptr, je.Toptr(), strx, utfx) // 最后一个参数很奇怪
	gopp.GOUSED(rv)
}

func (je JNIEnv) GetStringUTFChars(sx voidptr) string {
	var copyed uint8
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).GetStringUTFChars)
	rv := Litfficallg(fnptr, je.Toptr(), sx, voidptr(&copyed))
	gopp.GOUSED(rv, copyed)
	defer je.ReleaseStringUTFChars(sx, rv)

	return GoString(rv)
}
func (je JNIEnv) GetStringUTFLength(sx voidptr) int {
	var e2 = (*C.JNIEnv)(voidptr(je))
	var fnptr = voidptr((*e2).GetStringUTFLength)
	rv := Litfficallg(fnptr, je.Toptr(), sx)
	gopp.GOUSED(rv)
	return int(usize(rv))
}

func GetCStringddd() {
	v := C.getCStringddd(0, 0, nil)
	log.Println(v)
}
func AndroidNameddd() {
	v := C.androidNameddd(0, 0, 0)
	log.Println(v)
}
