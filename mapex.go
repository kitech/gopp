package gopp

import (
	"math/rand"
	"reflect"
)

// map的更多的操作

func MapKeys[KT comparable, VT any](m map[KT]VT) []KT {
	mt := reflect.ValueOf(m)
	if mt.Kind() != reflect.Map {
		return nil
	}

	outs := make([]KT, len(m))
	mv := reflect.ValueOf(m)
	for i, vk := range mv.MapKeys() {
		// vv := mv.MapIndex(vk).Interface()
		// outs = append(outs, vk.Interface().(KT))
		outs[i] = vk.Interface().(KT)
	}
	return outs
}

func MapValues[KT comparable, VT any](m map[KT]VT) []VT {
	mt := reflect.ValueOf(m)
	if mt.Kind() != reflect.Map {
		return nil
	}

	outs := make([]VT, len(m))
	mv := reflect.ValueOf(m)
	for i, vk := range mv.MapKeys() {
		vv := mv.MapIndex(vk).Interface()
		// outs = append(outs, vv.(VT))
		outs[i] = vv.(VT)
	}
	return outs
}

func MapFlat[KT comparable, VT any](m map[KT]VT) []any {
	mt := reflect.ValueOf(m)
	if mt.Kind() != reflect.Map {
		return nil
	}

	outs := make([]any, len(m)*2)
	mv := reflect.ValueOf(m)
	for i, vk := range mv.MapKeys() {
		vv := mv.MapIndex(vk).Interface()

		outs[i*2] = vk.Interface()
		outs[i*2+1] = vv
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

func ArrayRand[T any](a []T) T {
	if len(a) == 0 {
		var v T
		return v
	}
	var i = (int(rand.Uint32()) + len(a)) % len(a)
	return a[i]
}
func ArrayLast[T any](a []T) T {
	if len(a) == 0 {
		var v T
		return v
	}
	var i = len(a) - 1
	return a[i]
}

func MapRand[KT comparable, VT any](m map[KT]VT) (KT, VT) {
	refval := reflect.ValueOf(m)
	iter := refval.MapRange()
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		return key.Interface().(KT), val.Interface().(VT)
	}
	var k KT
	var v VT
	return k, v
}
func MapRandKey[KT comparable, VT any](m map[KT]VT) KT {
	k, _ := MapRand(m)
	return k
}
func MapRandVal[KT comparable, VT any](m map[KT]VT) VT {
	_, v := MapRand(m)
	return v
}
