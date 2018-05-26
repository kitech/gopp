package gopp

/*
#include <stdint.h>
*/
import "C"
import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

func BytesDup(src []byte) []byte {
	r := make([]byte, len(src))
	n := copy(r, src)
	if n != len(src) {
		panic("wtf")
	}
	return r
}

func DeepCopy(from interface{}, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, to)
}

// deepcopy的一種實現，使用json作爲中轉
// github.com/getlantern/deepcopy
// 還有一種使用reflect遞歸copy所有元素
// https://github.com/mohae/deepcopy

// if integer, number, must use var's addr, like , a := 5; &a
// 一般用于数字类型的操作
// TODO 考虑类型的存储大小，防止丢失精度
func OpAssign(tovalp, fromvaluep interface{}) {
	toval := reflect.ValueOf(tovalp)
	log.Println(toval.CanAddr(), toval.Type().String())
	toty := toval.Type()
	fromvalue := reflect.ValueOf(fromvaluep)
	fromty := fromvalue.Type()
	if fromty.Elem().AssignableTo(toty.Elem()) {
		toval.Elem().Set(fromvalue.Elem())
	} else if fromty.Elem().ConvertibleTo(toty.Elem()) {
		toval.Elem().Set(fromvalue.Elem().Convert(toty.Elem()))
	} else {
		log.Panicln("Connot assign.", toty.String(), fromty.String())
	}
}

func OpEqual(left, right interface{}) bool {
	return false
}

func OpGreatThan(left, right interface{}) bool {
	return false
}
func OpGreatOrEqual(left, right interface{}) bool {
	return false
}
func OpLessThan(left, right interface{}) bool {
	return false
}
func OpLessOrEqual(left, right interface{}) bool {
	return false
}

func _TestAssign1(t *testing.T) {
	var to C.uint32_t = 123
	var from int = 567
	OpAssign(&to, &from)
	if to != C.uint32_t(from) {
		t.Fail()
	}
}
