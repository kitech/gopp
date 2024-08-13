//go:build usejni
// +build usejni

package cgopp

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
import (
	"log"
	"runtime"
)

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

func jnimemfninit(jvm JavaVM, je JNIEnv) {
	mf := jnimemfn

	mf.GetEnv = voidptr((*(*C.JavaVM)(voidptr(jvm))).GetEnv)
	mf.AttachCurrentThread = voidptr((*(*C.JavaVM)(voidptr(jvm))).AttachCurrentThread)
	mf.DetachCurrentThread = voidptr((*(*C.JavaVM)(voidptr(jvm))).DetachCurrentThread)
	mf.AttachCurrentThreadAsDaemon = voidptr((*(*C.JavaVM)(voidptr(jvm))).AttachCurrentThreadAsDaemon)
	mf.DestroyJavaVM = voidptr((*(*C.JavaVM)(voidptr(jvm))).DestroyJavaVM)

	mf.FindClass = voidptr((*(*C.JNIEnv)(voidptr(je))).FindClass)
	mf.GetStaticMethodID = voidptr((*(*C.JNIEnv)(voidptr(je))).GetStaticMethodID)
	mf.CallStaticObjectMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticObjectMethod)
	mf.CallStaticVoidMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticVoidMethod)
	mf.CallStaticIntMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticIntMethod)
	mf.CallStaticLongMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticLongMethod)
	mf.CallStaticFloatMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticFloatMethod)
	mf.CallStaticDoubleMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticDoubleMethod)
	mf.CallStaticBooleanMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallStaticBooleanMethod)
	mf.GetMethodID = voidptr((*(*C.JNIEnv)(voidptr(je))).GetMethodID)
	mf.CallObjectMethod = voidptr((*(*C.JNIEnv)(voidptr(je))).CallObjectMethod)

	mf.NewStringUTF = voidptr((*(*C.JNIEnv)(voidptr(je))).NewStringUTF)
	mf.GetStringUTFLength = voidptr((*(*C.JNIEnv)(voidptr(je))).GetStringUTFLength)
	mf.GetStringUTFChars = voidptr((*(*C.JNIEnv)(voidptr(je))).GetStringUTFChars)
	mf.ReleaseStringUTFChars = voidptr((*(*C.JNIEnv)(voidptr(je))).ReleaseStringUTFChars)

	mf.GetObjectClass = voidptr((*(*C.JNIEnv)(voidptr(je))).GetObjectClass)
	mf.GetFieldID = voidptr((*(*C.JNIEnv)(voidptr(je))).GetFieldID)
	mf.GetIntField = voidptr((*(*C.JNIEnv)(voidptr(je))).GetIntField)
	mf.GetLongField = voidptr((*(*C.JNIEnv)(voidptr(je))).GetLongField)

	mf.ExceptionCheck = voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionCheck)
	mf.ExceptionOccurred = voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionOccurred)
	mf.ExceptionDescribe = voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionDescribe)
	mf.ExceptionClear = voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionClear)

	mf.NewGlobalRef = voidptr((*(*C.JNIEnv)(voidptr(je))).NewGlobalRef)
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

func (je JNIEnv) fnGetObjectClass() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetObjectClass)
}

func (je JNIEnv) fnGetFieldID() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetFieldID)
}
func (je JNIEnv) fnGetIntField() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetIntField)
}
func (je JNIEnv) fnGetLongField() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).GetLongField)
}

func (je JNIEnv) fnExceptionCheck() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionCheck)
}
func (je JNIEnv) fnExceptionOccurred() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionOccurred)
}
func (je JNIEnv) fnExceptionDescribe() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionDescribe)
}
func (je JNIEnv) fnExceptionClear() voidptr {
	return voidptr((*(*C.JNIEnv)(voidptr(je))).ExceptionClear)
}

////

func (j JavaVM) Env() JNIEnv {
	var env JNIEnv
	envx := C.getjavaenvbyjavavm(voidptr(j))
	env = JNIEnv(envx)
	return env
}
func (j JNIEnv) Toc() *C.JNIEnv { return (*C.JNIEnv)(voidptr(j)) }

func GetCStringddd() {
	v := C.getCStringddd(0, 0, nil)
	log.Println(v)
}
func AndroidNameddd() {
	v := C.androidNameddd(0, 0, 0)
	log.Println(v)
}
