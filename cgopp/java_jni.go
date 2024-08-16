// //goddd:build usejni
// // +buildddd usejni

package cgopp

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/kitech/gopp"
)

// 似乎这个不管用，是因为忽略了参数中的env，这个env是和当前线程绑定的
// 在不能保证线程是否为JVM线程时，使用参数中的JNIEnv变量
func RunOnJVM[FT func(vm, env, ctx usize) error |
	func(env usize) error | func(env usize)](fnx FT) error {
	switch fn := any(fnx).(type) {
	case func(vm, env, ctx usize) error:
		return mobinitRunOnJVM(fn)
	case func(usize) error:
		var fn2 = func(vm, env, ctx usize) error { return fn(env) }
		return mobinitRunOnJVM(fn2)
	case func(usize):
		var fn2 = func(vm, env, ctx usize) error { fn(env); return nil }
		return mobinitRunOnJVM(fn2)
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

func (jvm JavaVM) Env() JNIEnv {
	if jmf.GetEnv == nil {
		// maybe not inited
		var envx = getJavaVMEnvByc(usize(jvm))
		var env = JNIEnv(envx)
		return env
	} else {
		var iop voidptr
		rv := FfiCall[int32](jmf.GetEnv, jvm, voidptr(&iop), JNI_VERSION_1_6)
		if rv != 0 {
			log.Println("some err", jvm)
		}
		return JNIEnv(rv)
	}
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
	GetObjectField voidptr
	GetIntField    voidptr
	GetLongField   voidptr
	GetDoubleField voidptr

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
	rv := Litfficall(jmf.GetVersion, je.Toptr())
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
	v := FfiCall[usize](jmf.FindClass, je, usize(cls4c))
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

	rv := FfiCall[usize](jmf.GetStaticMethodID, je, clsid, s4c, argssig4c)
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
	rv := FfiCall[int](jmf.CallStaticVoidMethod, argv[:argc]...)
	gopp.GOUSED(rv)
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
		rvp := FfiCall[usize](jmf.CallStaticObjectMethod, argv...)
		rv2 := je.GetStringUTFChars(rvp)
		rvx = any(rv2).(RTY)
	case usize:
		// rvp := FfiCall[voidptr](je.fnCallStaticObjectMethod(), argv...)
		rvx = FfiCall[RTY](jmf.CallObjectMethod, argv...)
	case int, uint: // go的int是变长类型，要按照java的类型调用
		// log.Println(rvx, clsid, mthid, len(argv), argv)
		rvx = FfiCall[RTY](jmf.CallStaticIntMethod, argv...)

	case int64, uint64:
		// log.Println(rvx, clsid, mthid, len(argv), argv)
		rvx = FfiCall[RTY](jmf.CallStaticLongMethod, argv...)
		// log.Println(rvx, clsid, mthid, len(argv), argv)

	case float64:
		rvx = FfiCall[RTY](jmf.CallStaticDoubleMethod, argv...)
	case Void:
		// log.Println(rvx, clsid, mthid, len(argv), argv)
		rvx = FfiCall[RTY](jmf.CallStaticVoidMethod, argv...)
	default:
		log.Println("Nocat", reflect.TypeOf(any(rvx)))
	}

	return
}

func (je JNIEnv) GetMethodID(clsid usize, mthname string, argssig string) usize {
	// log.Println(clsid, s, argssig)
	var s4c = CStringgc(mthname)
	var argssig4c = CStringgc(argssig)

	rv := FfiCall[usize](jmf.GetMethodID, je, clsid, s4c, argssig4c)
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
		rvp := FfiCall[usize](jmf.CallObjectMethod, argv...)
		rv2 := je.GetStringUTFChars(rvp)
		rvx = any(rv2).(RTY)
	case usize:
		// rvp := FfiCall[voidptr](je.fnCallStaticObjectMethod(), argv...)
		rvx = FfiCall[RTY](jmf.CallObjectMethod, argv...)

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
	rv := Litfficall(jmf.GetObjectClass, je.Toptr(), voidptr(obj))
	return usize(rv)
}
func (je JNIEnv) GetFieldID(clsobj usize, a0, a1 string) usize {
	a04c := CStringgc(a0)
	a14c := CStringgc(a1)

	rv := Litfficall(jmf.GetFieldID, je.Toptr(), voidptr(clsobj), a04c, a14c)
	return usize(rv)
}

func (je JNIEnv) GetObjectField(clsobj usize, fidobj usize) usize {
	rv := Litfficall(jmf.GetObjectField, je.Toptr(), voidptr(fidobj))
	return usize(rv)
}
func (je JNIEnv) GetIntField(clsobj usize, fidobj usize) int {
	rv := Litfficall(jmf.GetIntField, je.Toptr(), voidptr(fidobj))
	return int(usize(rv))
}
func (je JNIEnv) GetLongField(clsobj usize, fidobj usize) int64 {
	rv := Litfficall(jmf.GetLongField, je.Toptr(), voidptr(fidobj))
	return int64(usize(rv))
}
func (je JNIEnv) GetDoubleField(clsobj usize, fidobj usize) float64 {
	// rv := Litfficall(jmf.GetDoubleField, je.Toptr(), voidptr(fidobj))
	// return int64(usize(rv))
	rv := FfiCall[float64](jmf.GetDoubleField, je, fidobj)
	return rv
}

func (je JNIEnv) ExceptionClear() {
	rv := Litfficallg(jmf.ExceptionClear, je.Toptr())
	gopp.GOUSED(rv)
}
func (je JNIEnv) ExceptionCheck() bool {
	rv := Litfficallg(jmf.ExceptionCheck, je.Toptr())
	if rv != nil {
		log.Println("Some error", rv, usize(rv), MyTid())
		je.ExceptionDescribe()
		JNIThreadCheck()
		panic(rv)
	}
	return rv != nil
}
func (je JNIEnv) ExceptionDescribe() {
	rv := Litfficallg(jmf.ExceptionDescribe, je.Toptr())
	if rv != nil {
		log.Println("Some error", rv, usize(rv))
	}
}

func (je JNIEnv) NewGlobalRef(obj usize) usize {
	rv := FfiCall[usize](jmf.NewGlobalRef, je, obj)
	if rv == 0 {
		log.Println("Some error", obj, rv)
	}
	return rv
}
func (je JNIEnv) DeleteGlobalRef(obj usize) {
	FfiCall[Void](jmf.DeleteGlobalRef, je, obj)
}
func (je JNIEnv) IsSameObject(obj0, obj1 usize) bool {
	rv := FfiCall[uint8](jmf.IsSameObject, je, obj0, obj1)
	return rv != 0
}

// 这个生成的object不需要自己释放
func (je JNIEnv) NewStringUTF(s string) usize {
	s4c := CStringgc(s)

	rv := FfiCall[usize](jmf.NewStringUTF, je.Toptr(), s4c)
	return rv
}

func (je JNIEnv) ReleaseStringUTFChars(strx usize, utfx usize) {
	// JNI DETECTED ERROR IN APPLICATION: non-nullable argument was NULL
	rv := FfiCall[Void](jmf.ReleaseStringUTFChars, je, strx, utfx) // 最后一个参数很奇怪
	gopp.GOUSED(rv)
}

func (je JNIEnv) GetStringUTFChars(sx usize) string {
	var copyed uint8
	rv := FfiCall[voidptr](jmf.GetStringUTFChars, je.Toptr(), sx, usize(voidptr((&copyed))))
	gopp.GOUSED(rv, copyed)
	defer je.ReleaseStringUTFChars(sx, usize(rv))

	return GoString(rv)
}
func (je JNIEnv) GetStringUTFLength(sx usize) int {
	rv := FfiCall[int32](jmf.GetStringUTFLength, je, sx)
	gopp.GOUSED(rv)
	return int(rv)
}
