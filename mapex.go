package gopp

import (
	"reflect"
)

// map的更多的操作

func MapKeys(m interface{}) []interface{} {
	mt := reflect.ValueOf(m)
	if mt.Kind() != reflect.Map {
		return nil
	}

	mv := reflect.ValueOf(m)
	return Value2Iface(mv.MapKeys())
}

func MapValues(m interface{}) []interface{} {
	mt := reflect.ValueOf(m)
	if mt.Kind() != reflect.Map {
		return nil
	}

	outs := make([]interface{}, 0)
	mv := reflect.ValueOf(m)
	for _, vk := range mv.MapKeys() {
		vv := mv.MapIndex(vk).Interface()
		outs = append(outs, vv)
	}
	return outs
}

type Map struct {
	m interface{}
}

func NewMap(m interface{}) *Map { return &Map{m} }

type Array struct {
	a Any
}

func NewArray(a interface{}) *Array {
	this := &Array{}
	this.a = ToAny(a)
	return this
}

func (this *Array) Contains(i interface{}) bool {
	return false
}

func (this *Array) Length() int {
	return 0
}

func MapFirst(m any) any {
	refval := reflect.ValueOf(m)
	iter := refval.MapRange()
	for iter.Next() {
		val := iter.Value()
		return val
	}
	return nil
}
func MapFirstStr(m any) string {
	val := MapFirst(m)
	if val != nil {
		return val.(string)
	}
	return ""
}
