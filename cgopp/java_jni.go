package cgopp

/*
#include <jni.h>

#include "java_jni.h"

#cgo CFLAGS: -I/nix/store/6wrwg9jhabsx0mbq0j4ximym65zd0i0h-zulu-ca-jdk-22.0.0/include
*/
import "C"
import "log"

// see Running native Android code in a Fyne app

// extern /*jclass*/void* find_class(/*JNIEnv**/uintptr_t jni_env, const char* class_name);
// extern const char* getCString(uintptr_t jni_env, uintptr_t ctx, /*jstring*/ void* str);
// extern const char* androidName(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);

func FindClass() {
	v := C.find_class(0, C.CString("class/full/name"))
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
