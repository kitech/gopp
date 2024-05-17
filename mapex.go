package gopp

import (
	"math/rand"
	"reflect"
)

// gopp.MapSS // 这个还是更简短些
// map[string]string
type MapSS = map[string]string // 应该这个用的最多吧
type MapSA = map[string]any
type MapIA = map[int]any
type MapSI = map[string]int
type MapIS = map[int]string
type MapSL = map[string]int64 // means LongLong
type MapLS = map[int64]string

// map的更多的操作

func Keysof[KT comparable, VT any](m map[KT]VT) []KT {
	return MapKeys(m)
}
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

func Valuesof[KT comparable, VT any](m map[KT]VT) []VT {
	return MapValues(m)
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

// 这个可以用 Mapdo 实现，不过这个是带类型的，好点
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
	m any
}

func NewMap(m any) *Map { return &Map{m} }

type Array struct {
	a Any
}

func NewArray(a any) *Array {
	this := &Array{}
	this.a = ToAny(a)
	return this
}

func (this *Array) Contains(i any) bool {
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
