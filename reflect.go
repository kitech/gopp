package gopp

import (
	"bytes"
	"encoding/gob"
	"log"
	"reflect"
	// _ "github.com/goccy/go-reflect"
	// "github.com/goccy/go-reflect"
	// _ "github.com/viant/xunsafe"
)

func Castto[T any, U any](from *T) *U {
	return (*U)(voidptr(from))
}

type GorefValue struct {
	// Typ  *abi.Type
	Typ  *Abitype
	Ptr  voidptr
	Flag GorefFlag
}

func (me *GorefValue) Org() *reflect.Value { return ((*reflect.Value)(voidptr(me))) }

// for unexport
func (me *GorefValue) UnsetRO() {
	me.Flag = me.Flag & (^flagRO)
}
func GorefValueUnsetRO(vv ...*reflect.Value) {
	for i := 0; i < len(vv); i++ {
		v2 := (*GorefValue)(voidptr(vv[i]))
		v2.UnsetRO()
	}
}

type GorefFlag uintptr

const (
	flagKindWidth             = 5 // there are 27 kinds
	flagKindMask    GorefFlag = 1<<flagKindWidth - 1
	flagStickyRO    GorefFlag = 1 << 5
	flagEmbedRO     GorefFlag = 1 << 6
	flagIndir       GorefFlag = 1 << 7
	flagAddr        GorefFlag = 1 << 8
	flagMethod      GorefFlag = 1 << 9
	flagMethodShift           = 10
	flagRO          GorefFlag = flagStickyRO | flagEmbedRO
)

// todo 可能出现 fatal error: concurrent map read and map write
// how deal unexport field??? UnsetRO
// 需要一个更强大的reflect库
func DeepSizeof(vx any, depth int) (rv int) {
	valx := reflect.ValueOf(vx)
	GorefValueUnsetRO(&valx)
	vty := reflect.TypeOf(vx)
	if vty == nil {
		// todo???
		// log.Println(vx, depth)
		return 0
	}

	// log.Println(depth, vty.String(), vty.Size())
	switch vty.Kind() {
	case reflect.Pointer:
		if valx.IsNil() {
			rv += int(UintptrTySz)
			break
		}

		e := valx.Elem().Interface()
		rv += DeepSizeof(e, depth+1)

	case reflect.Chan: // todo 不准确的
		elemtysz := valx.Elem().Type().Size()
		elemcnt := valx.Len()
		rv += elemcnt * int(elemtysz)

	case reflect.Struct:
		for i := 0; i < valx.NumField(); i++ {
			fvxy := valx.Field(i)
			if fvxy.CanInterface() {
				fvx := fvxy.Interface()
				rv += DeepSizeof(fvx, depth+1)
			} else {
				// grv := (*GorefValue)((voidptr)(&valx))
				log.Println("Cannot iface, unexported?", vty.String(), vty.Field(i).Name)
			}

		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < valx.Len(); i++ {
			fvx := valx.Index(i).Interface()
			rv += DeepSizeof(fvx, depth+1)
		}
	case reflect.Map:
		keysx := valx.MapKeys()
		for _, keyx := range keysx {
			mvx := valx.MapIndex(keyx)
			rv += DeepSizeof(keyx.Interface(), depth+1)
			rv += DeepSizeof(mvx.Interface(), depth+1)
		}
	case reflect.String:
		rv += len(vx.(string))
	default:
		// rv += int(unsafe.Sizeof(vx))
	}
	rv += int(vty.Size())
	return
}

// todo ob: type Foo has no exported fields Bar

func DeepSizeBygob(vx any) int {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(vx)
	ErrPrint(err, reflect.TypeOf(vx))
	return b.Len()
}

func IsnilNotyped(v any) bool {
	return false
}

func IsnilTyped(v any) bool {
	return false
}
