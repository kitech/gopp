////

//go:build usejni
// +build usejni

package cgopp

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/kitech/gopp"
	mobinit "github.com/kitech/gopp/internal/mobileinit"
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

// 这种方式隔离的更彻底，
// 准备在JNI_OnLoaded里安化，不知道这些函数指针不同线程是否一样的
type jnimemfnst struct {
	// JavaVM
	GetEnv                      voidptr
	AttachCurrentThread         voidptr
	DetachCurrentThread         voidptr
	AttachCurrentThreadAsDaemon voidptr
	DestroyJavaVM               voidptr

	// JNIEnv
	GetVersion              voidptr
	FindClass               voidptr
	GetStaticMethodID       voidptr
	CallStaticVoidMethod    voidptr
	CallStaticObjectMethod  voidptr
	CallStaticIntMethod     voidptr
	CallStaticLongMethod    voidptr
	CallStaticFloatMethod   voidptr
	CallStaticDoubleMethod  voidptr
	CallStaticBooleanMethod voidptr
	GetMethodID             voidptr
	CallObjectMethod        voidptr

	NewStringUTF          voidptr
	GetStringUTFChars     voidptr
	GetStringUTFLength    voidptr
	ReleaseStringUTFChars voidptr

	GetObjectClass voidptr
	GetFieldID     voidptr
	GetIntField    voidptr
	GetLongField   voidptr

	ExceptionCheck    voidptr
	ExceptionOccurred voidptr
	ExceptionDescribe voidptr
	ExceptionClear    voidptr

	NewGlobalRef    voidptr
	DeleteGlobalRef voidptr
	IsSameObject    voidptr
}

var jnimemfn = &jnimemfnst{}
var jmf = jnimemfn

func (j JNIEnv) Tocuptr() cuptr { return (cuptr)(j) }
func (j JNIEnv) Toptr() voidptr { return voidptr(j) }
func (j JNIEnv) String() string { return fmt.Sprintf("%v", voidptr(j)) }

func (je JNIEnv) GetVersion() int {
	var fnptr = je.fnGetVersion()
	rv := Litfficall(fnptr, je.Toptr())
	// log.Println(rv)
	return int(usize(rv))
}

// https://stackoverflow.com/questions/19113719/jni-findclass-function-returns-null
// 好像必须在jvm主线程中查找类
// 尝试L前缀，尝试把.替换/
// 为什么有时需要前缀L，有时不需要，像Main
// 使用 javap -s -p Main 查看函数签名信息
// cls: "Ljava/lang/String"
func (je JNIEnv) FindClass(cls string) usize {
	if strings.Count(cls, ".") > 0 {
		cls = strings.ReplaceAll(cls, ".", "/")
	}
	var cls4c = CStringgc(cls)
	log.Println(je.Tocuptr(), je, cls)

	// runtime: bad pointer in frame main.getJvmMemory.func1 at 0x400005fea8: 0x25
	v := FfiCall[usize](je.fnFindClass(), je, usize(cls4c))
	if v != 0 {
		return v
	}

	return 0
}

// s: ???
// argssig: "([Ljava/lang/String;)V"
// ()V
// 有的结尾有分号，有的没有
// 使用 javap -s -p Main 查看函数签名信息
func (je JNIEnv) GetStaticMethodID(clsid usize, mthname string, argssig string) usize {
	// log.Println(clsid, s, argssig)
	var s4c = CStringgc(mthname)
	var argssig4c = CStringgc(argssig)

	var fnptr = je.fnGetStaticMethodID()
	rv := FfiCall[usize](fnptr, je, clsid, s4c, argssig4c)
	je.ExceptionCheck()
	return rv
}

// support string/int/todo
// 这个应该使用的少
func (je JNIEnv) CallStaticVoidMethod(clsid, mthid usize, args ...any) {
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

	log.Println(args, Goargs2JvSignature(Void(0), args...))
	var fnptr = je.fnCallStaticVoidMethod()
	rv := FfiCall[int](fnptr, argv[:argc]...)
	gopp.GOUSED(rv)
}

// jni没有查看类型的函数！！！
// jvalue 类型?
type Jany uintptr

func (me Jany) Tostr() string {
	return ""
}

// go的方法不能带模板类型!!!
// 不支持float/double类型参数和返回值
// support args: string/int/todo
// support ret: string/int/todo
func JNIEnvCallStaticMethod[RTY any](je JNIEnv, clsid, mthid usize, args ...any) (rvx RTY) {

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

	log.Println(args, Goargs2JvSignature(rvx, args...))

	switch any(rvx).(type) {
	case string:
		// rvp := FfiCall[voidptr](je.fnCallStaticObjectMethod(), argv...)
		rvp := FfiCall[usize](jnimemfn.CallStaticObjectMethod, argv...)
		rv2 := je.GetStringUTFChars(rvp)
		rvx = any(rv2).(RTY)
	case usize:
		// rvp := FfiCall[voidptr](je.fnCallStaticObjectMethod(), argv...)
		rvx = FfiCall[RTY](jnimemfn.CallObjectMethod, argv...)
	case int, uint: // go的int是变长类型，要按照java的类型调用
		// log.Println(rvx, clsid, mthid, len(argv), argv)
		rvx = FfiCall[RTY](je.fnCallStaticIntMethod(), argv...)

	case int64, uint64:
		// log.Println(rvx, clsid, mthid, len(argv), argv)
		rvx = FfiCall[RTY](je.fnCallStaticLongMethod(), argv...)
		// log.Println(rvx, clsid, mthid, len(argv), argv)

	case float64:
		rvx = FfiCall[RTY](je.fnCallStaticDoubleMethod(), argv...)
	case Void:
		// log.Println(rvx, clsid, mthid, len(argv), argv)
		rvx = FfiCall[RTY](je.fnCallStaticVoidMethod(), argv...)
	default:
		log.Println("Nocat", reflect.TypeOf(any(rvx)))
	}

	return
}

func (je JNIEnv) GetMethodID(clsid usize, mthname string, argssig string) usize {
	// log.Println(clsid, s, argssig)
	var s4c = CStringgc(mthname)
	var argssig4c = CStringgc(argssig)

	var fnptr = jmf.GetMethodID
	rv := FfiCall[usize](fnptr, je, clsid, s4c, argssig4c)
	je.ExceptionCheck()
	return rv
}

func JNIEnvCallMethod[RTY any](je JNIEnv, clsid, mthid usize, args ...any) (rvx RTY) {

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

	log.Println(args, Goargs2JvSignature(rvx, args...))

	switch any(rvx).(type) {
	case string:
		// rvp := FfiCall[voidptr](je.fnCallStaticObjectMethod(), argv...)
		rvp := FfiCall[usize](jnimemfn.CallObjectMethod, argv...)
		rv2 := je.GetStringUTFChars(rvp)
		rvx = any(rv2).(RTY)
	case usize:
		// rvp := FfiCall[voidptr](je.fnCallStaticObjectMethod(), argv...)
		rvx = FfiCall[RTY](jnimemfn.CallObjectMethod, argv...)

	// case int, uint: // go的int是变长类型，要按照java的类型调用
	// 	// log.Println(rvx, clsid, mthid, len(argv), argv)
	// 	rvx = FfiCall[RTY](je.fnCallIntMethod(), argv...)

	// case int64, uint64:
	// 	// log.Println(rvx, clsid, mthid, len(argv), argv)
	// 	rvx = FfiCall[RTY](je.fnCallLongMethod(), argv...)
	// 	// log.Println(rvx, clsid, mthid, len(argv), argv)

	// case float64:
	// 	rvx = FfiCall[RTY](je.fnCallDoubleMethod(), argv...)
	// case Void:
	// 	// log.Println(rvx, clsid, mthid, len(argv), argv)
	// 	rvx = FfiCall[RTY](je.fnCallVoidMethod(), argv...)
	default:
		log.Println("Nocat", reflect.TypeOf(any(rvx)))
	}

	return
}

// https://stackoverflow.com/questions/40004522/how-to-get-values-from-jobject-in-c-using-jni

func (je JNIEnv) GetObjectClass(obj usize) usize {
	var fnptr = je.fnGetObjectClass()
	rv := Litfficall(fnptr, je.Toptr(), voidptr(obj))
	return usize(rv)
}
func (je JNIEnv) GetFieldID(clsobj usize, a0, a1 string) usize {
	a04c := CStringgc(a0)
	a14c := CStringgc(a1)

	var fnptr = je.fnGetFieldID()
	rv := Litfficall(fnptr, je.Toptr(), voidptr(clsobj), a04c, a14c)
	return usize(rv)
}
func (je JNIEnv) GetIntField(clsobj usize, fidobj usize) int {
	var fnptr = je.fnGetIntField()
	rv := Litfficall(fnptr, je.Toptr(), voidptr(fidobj))
	return int(usize(rv))
}

func (je JNIEnv) ExceptionClear() {
	var fnptr = je.fnExceptionClear()
	rv := Litfficallg(fnptr, je.Toptr())
	gopp.GOUSED(rv)
}
func (je JNIEnv) ExceptionCheck() bool {
	var fnptr = je.fnExceptionCheck()
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
	var fnptr = je.fnExceptionDescribe()
	rv := Litfficallg(fnptr, je.Toptr())
	if rv != nil {
		log.Println("Some error", rv, usize(rv))
	}
}

func (je JNIEnv) NewGlobalRef(obj usize) usize {
	var fnptr = jmf.NewGlobalRef
	rv := FfiCall[usize](fnptr, je, obj)
	if rv == 0 {
		log.Println("Some error", obj, rv)
	}
	return rv
}
func (je JNIEnv) DeleteGlobalRef(obj usize) {
	var fnptr = jmf.DeleteGlobalRef
	FfiCall[Void](fnptr, je, obj)
}
func (je JNIEnv) IsSameObject(obj0, obj1 usize) bool {
	var fnptr = jmf.IsSameObject
	rv := FfiCall[uint8](fnptr, je, obj0, obj1)
	return rv != 0
}

// 这个生成的object不需要自己释放
func (je JNIEnv) NewStringUTF(s string) usize {
	s4c := CStringgc(s)

	var fnptr = je.fnNewStringUTF()
	rv := FfiCall[usize](fnptr, je.Toptr(), s4c)
	return rv
}

func (je JNIEnv) ReleaseStringUTFChars(strx usize, utfx usize) {
	// JNI DETECTED ERROR IN APPLICATION: non-nullable argument was NULL
	var fnptr = je.fnReleaseStringUTFChars()
	rv := FfiCall[Void](fnptr, je, strx, utfx) // 最后一个参数很奇怪
	gopp.GOUSED(rv)
}

func (je JNIEnv) GetStringUTFChars(sx usize) string {
	var copyed uint8
	var fnptr = je.fnGetStringUTFChars()
	rv := FfiCall[voidptr](fnptr, je.Toptr(), sx, usize(voidptr((&copyed))))
	gopp.GOUSED(rv, copyed)
	defer je.ReleaseStringUTFChars(sx, usize(rv))

	return GoString(rv)
}
func (je JNIEnv) GetStringUTFLength(sx usize) int {
	var fnptr = je.fnGetStringUTFLength()
	rv := FfiCall[int32](fnptr, je, sx)
	gopp.GOUSED(rv)
	return int(rv)
}
