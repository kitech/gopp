package cgopp

import (
	"reflect"
	"strings"
	"unsafe"
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
type Jclass = voidptr
type Jobject = voidptr
type Jstring = voidptr
type Jarray = voidptr
type Jthrowable = voidptr
type Jweak = voidptr
type JfieldID = voidptr
type JmethodID = voidptr

// note: 只支持基础类型和String
func Goargs2JvSignature(rv any, args ...any) string {
	var sb = strings.Builder{}
	sb.WriteRune('(')

	// args = append(args, rv)
	for i, argx := range args {
		sb.WriteString(Goarg2Jvtype(false, argx))

		if i < len(args)-1 {
			sb.WriteRune(',')
		}
	}

	sb.WriteRune(')')
	if _, ok := rv.(Void); ok {
		sb.WriteRune('V')
	} else {
		sb.WriteString(Goarg2Jvtype(false, rv))
	}

	return sb.String()
}

// full: true, java type name
// full: false, java sig name
func Goarg2Jvtype(full bool, argx any) (rv string) {
	rv = "???"

	ty := reflect.TypeOf(argx)
	switch ty.Kind() {
	case reflect.String:
		if full {
			rv = "java.lang.String"
		} else {
			rv = "Ljava/lang/String"
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
