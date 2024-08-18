package cgopp

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/kitech/gopp"
)

// not need usejni tag

var JNI_OnLoad_Callback = func() {}

// ///
// jobject, jclass, jstring, jarray... = voidptr
type Void int
type Jint = int32 // int32_t
type Jsize = int32
type Jlong = int64 // int64_t
type Jbool = uint8
type Jbyte = int8
type Jshort = int16
type Jclass = usize
type Jobject = usize
type Jstring = usize
type Jarray = usize
type Jthrowable = usize
type Jweak = usize
type JfieldID = usize
type JmethodID = usize

// jni没有查看类型的函数！！！
// jvalue 类型?
type Jany usize

func (me Jany) Tostr() string {
	return ""
}

func JNICallstm[RTY any](je JNIEnv, cls, mth string, sigt string, args ...any) (rv RTY) {
	rv = jnicallmthany[RTY](je, cls, mth, sigt, false, 0, args...)
	return
}

func JNICalldym[RTY any](je JNIEnv, cls, mth string, sigt string, thiso usize, args ...any) (rv RTY) {
	rv = jnicallmthany[RTY](je, cls, mth, sigt, true, thiso, args...)
	return
}

func jnicallmthany[RTY any](je JNIEnv, cls, mth string, sigt string, isstatic bool, thiso usize, args ...any) (rv RTY) {
	clsid := je.FindClass2(cls)
	gopp.ZeroPrint(clsid, "cannot find class", cls)
	if clsid == 0 {
		return
	}
	var mthid usize
	if isstatic {
		mthid = je.GetStaticMethodID(clsid, mth, sigt)
	} else {
		mthid = je.GetMethodID(clsid, mth, sigt)
	}
	gopp.ZeroPrint(mthid, "cannot find method", cls, mth, sigt)
	if mthid == 0 {
		return
	}

	if isstatic {
		rv = JNIEnvCallStaticMethod[RTY](je, clsid, mthid, args...)
	} else {
		gopp.ZeroPrint(thiso, "this is nil to crash")
		nargs := append(gopp.Sliceof(any(thiso)), args...)
		rv = JNIEnvCallMethod[RTY](je, clsid, mthid, nargs...)
	}

	return
}

// note: 只支持基础类型和String
// todo 还有些分号没有处理好，结尾有的有分号有的没分号
func Goargs2JvSignature(rv any, args ...any) string {
	var sb = strings.Builder{}
	sb.WriteRune('(')

	// args = append(args, rv)
	for i, argx := range args {
		sb.WriteString(Goarg2Jvtype(false, false, argx))

		if i < len(args)-1 {
			sb.WriteRune(';')
		}
	}

	sb.WriteRune(')')
	if _, ok := rv.(Void); ok {
		sb.WriteRune('V')
	} else {
		sb.WriteString(Goarg2Jvtype(false, true, rv))
	}

	return sb.String()
}

// full: true, java type name
// full: false, java sig name
func Goarg2Jvtype(full bool, ret bool, argx any) (rv string) {
	rv = "???"

	ty := reflect.TypeOf(argx)
	switch ty.Kind() {
	case reflect.String:
		if full {
			rv = "java.lang.String"
		} else {
			if ret {
				rv = "Ljava/lang/String;"
			} else {
				rv = "Ljava/lang/String"
			}
		}
	case reflect.Int: // => jlong, J, or jint I
		if unsafe.Sizeof(int(0)) == unsafe.Sizeof(int64(0)) {
			if full {
				rv = "long"
			} else {
				rv = "J"
			}
		} else if unsafe.Sizeof(int32(0)) == unsafe.Sizeof(int32(0)) {
			if full {
				rv = "int"
			} else {
				rv = "I"
			}
		} else {
			panic("wtf")
		}
	case reflect.Int32, reflect.Uint32: // => jint, I
		if full {
			rv = "int"
		} else {
			rv = "I"
		}
	case reflect.Int64, reflect.Uint64:
		if full {
			rv = "long"
		} else {
			rv = "J"
		}
	case reflect.Float32:
		if full {
			rv = "float"
		} else {
			rv = "f"
		}
	case reflect.Float64:
		if full {
			rv = "double"
		} else {
			rv = "d"
		}
	case reflect.Bool:
		if full {
			rv = "boolean"
		} else {
			rv = "b"
		}
	default:
	}

	return
}
