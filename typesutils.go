package gopp

import (
	"log"
	"reflect"
	"unicode"
)

func FloatIsZero[T float32 | float64](v T) bool {
	return v > -0.000001 && v < 0.000001
}

// see unicode.IsNumber()
func Isnumtype(ty reflect.Type) bool {
	switch ty.Kind() {
	case reflect.Uintptr, reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64,
		reflect.Bool, reflect.Int32, reflect.Uint32, reflect.Int16, reflect.Uint16,
		reflect.Float32, reflect.Float64, reflect.Int8, reflect.Uint8:
		return true
	}
	return false
}

func CalcMemlen(v any) int {
	refval := reflect.ValueOf(v)
	refty := refval.Type()

	switch refty.Kind() {
	case reflect.Ptr:
		return CalcMemlen(refval.Elem().Interface())
	case reflect.Struct:
		len1 := int(refty.Size())
		for i := 0; i < refty.NumField(); i++ {
			fldname := refty.Field(i).Name
			fldty := refty.Field(i).Type
			if Isnumtype(fldty) {
			} else {
				if unicode.IsUpper(rune(fldname[0])) {
					len1 += CalcMemlen(refval.Field(i).Interface())
				} else {
					log.Printf("unexport %s.%s %s\n", refty.Name(), fldname, fldty.String())
				}
			}
		}
		return len1
	case reflect.Slice:
		len1 := 0
		for i := 0; i < refval.Len(); i++ {
			len1 += CalcMemlen(refval.Index(i).Interface())
		}
		return len1
	case reflect.Map:
		len1 := 0
		for _, kval := range refval.MapKeys() {
			len1 += CalcMemlen(kval.Interface())
			len1 += CalcMemlen(refval.MapIndex(kval).Interface())
		}
		return len1
	case reflect.Chan:
		len1 := 0
		len1 = refval.Cap() * CalcMemlen(refval.Elem().Interface())
		return len1
	case reflect.UnsafePointer:
		log.Println("oh raw UnsafePointer")
		return 0
	case reflect.String:
		return refval.Len()
	default:
		if Isnumtype(refty) {
			return refty.Bits() / 8
		} else {
			log.Println("todo", refty.Kind())
		}
	}
	return 0
}

func calcmemlen_test() {
	type Message struct {
		Msg   string
		Links []string
	}
	m := &Message{}
	m.Msg = ""
	m.Links = []string{"abc", "efg"}
	rv := CalcMemlen(m)
	log.Println(rv)
	log.Println(CalcMemlen(567.890))
}
