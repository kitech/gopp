package gopp

import (
	"encoding/json"
	"fmt"
	"log"
	mrand "math/rand"
	"reflect"
	"strconv"
	"unsafe"
)

// want these as buildin type name
type i8 = int8
type u8 = uint8
type i16 = int16
type u16 = uint16
type i32 = int32
type u32 = uint32
type i64 = int64
type u64 = uint64
type f32 = float32
type f64 = float64
type f80 = [10]byte
type i128 = [16]uint8
type u128 = [16]byte
type vptr = unsafe.Pointer // void pointer

// TODO how add methods for Any type
// TODO how add methods for primity string type
type IAny any

/*
invalid receiver type *Any (Any is an interface type)
func (this *Any) Hehe() {
}
*/

type Any struct {
	I any
	// v T
	// v *reflect.Value
}

func AnyNew(i any) Any {
	// v := reflect.ValueOf(i)
	return Any{i}
}
func AnyOf(i any) Any {
	// v := reflect.ValueOf(i)
	return Any{i}
}
func ToAny(i any) Any {
	// v := reflect.ValueOf(i)
	return Any{i}
}
func (this Any) Raw() any     { return this.I }
func (this Any) IsNil() bool  { return this.I == nil }
func (this Any) Int() int     { return this.I.(int) }
func (this Any) Uint() uint   { return this.I.(uint) }
func (this Any) I8() int8     { return this.I.(int8) }
func (this Any) U8() uint8    { return this.I.(uint8) }
func (this Any) I16() int16   { return this.I.(int16) }
func (this Any) U16() uint16  { return this.I.(uint16) }
func (this Any) I32() int32   { return this.I.(int32) }
func (this Any) U32() uint32  { return this.I.(uint32) }
func (this Any) I64() int64   { return this.I.(int64) }
func (this Any) U64() uint64  { return this.I.(uint64) }
func (this Any) F32() float32 { return this.I.(float32) }
func (this Any) F64() float64 { return this.I.(float64) }
func (this Any) Str() string  { return this.I.(string) }
func (this Any) Iterable() bool {
	switch reflect.TypeOf(this.I).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map,
		reflect.String, reflect.Struct:
		return true
	}
	return false
}
func (this Any) CanMKey() bool {
	if true {
		return !this.Iterable()
	}
	return true // 是否能作为map的key
}

// todo any => Any
func (this Any) Mapdo(fx any) (out any) {
	out = Mapdo(this.I, fx)
	return
}

// /// as means type convert if need
func (this Any) AsStr() string  { return fmt.Sprintf("%v", this.I) }
func (this Any) AsStrp() string { return fmt.Sprintf("+%v", this.I) }
func (this Any) AsJson() []byte {
	bcc, err := json.Marshal(this.I)
	ErrPrint(err)
	return bcc
}
func (this Any) AsJsons() string { return string(this.AsJson()) }

// for json
func (this Any) String() string { return fmt.Sprintf("%v", this.I) }

func AnyAsg[T any](this Any) (rv T) {
	ety := reflect.TypeOf(this.I)
	toty := reflect.TypeOf(rv)

	rvx, ok := vecmapconvertvalue(this.I, ety, toty)
	if !ok {
		log.Println("cannot", ety.String(), "=>", toty.String())
		return
	}
	rv = rvx.(T)
	return
}

func (this Any) AsInt() (rv int)     { return AnyAsg[int](this) }
func (this Any) AsUint() (rv uint)   { return AnyAsg[uint](this) }
func (this Any) AsI64() (rv int64)   { return AnyAsg[int64](this) }
func (this Any) AsU64() (rv uint64)  { return AnyAsg[uint64](this) }
func (this Any) AsF64() (rv float64) { return AnyAsg[float64](this) }
func (this Any) AsF32() (rv float32) { return AnyAsg[float32](this) }

//////////////

// maybe can use Once for lazy
var vInt8Ty int8
var vUint8Ty uint8
var vIntTy int
var vUintTy uint
var vInt32Ty int32
var vUint32Ty uint32
var vInt64Ty int64
var vUint64Ty uint64
var vByteTy byte
var vFloat32Ty float32
var vFloat64Ty float64
var vBoolTy bool
var vStrTy string

var Int8Ty = reflect.TypeOf(vInt8Ty)
var Uint8Ty = reflect.TypeOf(vUint8Ty)
var IntTy = reflect.TypeOf(vIntTy)
var UintTy = reflect.TypeOf(vUintTy)
var Int32Ty = reflect.TypeOf(vInt32Ty)
var Uint32Ty = reflect.TypeOf(vUint32Ty)
var Int64Ty = reflect.TypeOf(vInt64Ty)
var Uint64Ty = reflect.TypeOf(vUint64Ty)
var ByteTy = reflect.TypeOf(vByteTy)
var Float32Ty = reflect.TypeOf(vFloat32Ty)
var Float64Ty = reflect.TypeOf(vFloat64Ty)
var BoolTy = reflect.TypeOf(vBoolTy)
var StrTy = reflect.TypeOf(vStrTy)

var RefKindTys = map[reflect.Kind]reflect.Type{
	reflect.Int8: Int8Ty, reflect.Uint8: Uint8Ty,
	reflect.Int: IntTy, reflect.Uint: UintTy,
	reflect.Int32: Int32Ty, reflect.Uint32: Uint32Ty,
	reflect.Int64: Int64Ty, reflect.Uint64: Uint64Ty,
	reflect.Float32: Float32Ty, reflect.Float64: Float64Ty,
	reflect.Bool: BoolTy, reflect.String: StrTy,
}

var Int8PtrTy = reflect.TypeOf(&vInt8Ty)
var Uint8PtrTy = reflect.TypeOf(&vUint8Ty)
var IntPtrTy = reflect.TypeOf(&vIntTy)
var UintPtrTy = reflect.TypeOf(&vUintTy)
var Int32PtrTy = reflect.TypeOf(&vInt32Ty)
var Uint32PtrTy = reflect.TypeOf(&vUint32Ty)
var Int64PtrTy = reflect.TypeOf(&vInt64Ty)
var Uint64PtrTy = reflect.TypeOf(&vUint64Ty)
var BytePtrTy = reflect.TypeOf(&vByteTy)
var Float32PtrTy = reflect.TypeOf(&vFloat32Ty)
var Float64PtrTy = reflect.TypeOf(&vFloat64Ty)
var BoolPtrTy = reflect.TypeOf(&vBoolTy)
var StrPtrTy = reflect.TypeOf(&vStrTy)

const ByteTySz = unsafe.Sizeof(byte(0))
const Int8TySz = strconv.IntSize / 4
const Int16TySz = Int8TySz * 2
const Int32TySz = Int8TySz * 4
const Int64TySz = Int8TySz * 8
const Float64TySz = Int8TySz * 8
const Float32TySz = Int8TySz * 4
const UintptrTySz = unsafe.Sizeof(uintptr(0))
const BoolTySz = unsafe.Sizeof(true)
const IntTySz = unsafe.Sizeof(int(0))
const RuneTySz = unsafe.Sizeof(rune(0))
const EmptyStructTySz = unsafe.Sizeof(struct{}{})

// const NilSz = unsafe.Sizeof(nil)

func IsMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func IsArray(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Array
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsChan(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Chan
}

func IsFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func IsPtr(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}

func Lenof(vx any) int {
	ty := reflect.TypeOf(vx)
	vv := reflect.ValueOf(vx)
	switch ty.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return vv.Len()
	}
	// return len(vx)
	return -1
}
func Capof(vx any) int {
	ty := reflect.TypeOf(vx)
	vv := reflect.ValueOf(vx)
	switch ty.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return vv.Cap()
	}
	// return len(vx)
	return -1
}

// 还是喜欢这种写法的！
// Lastof(vx).Str()
func Lastof(vx any) (rv Any) {
	if Lenof(vx) <= 0 {
		return
	}
	tv := reflect.ValueOf(vx)
	ty := tv.Type()

	switch ty.Kind() {
	case reflect.Slice, reflect.Array:
		ev := tv.Index(tv.Len() - 1)
		rv = ToAny(ev.Interface())

	case reflect.String:
		ev := tv.Index(tv.Len() - 1)
		rv = ToAny(ev.Interface())

	case reflect.Map:
		ek := tv.MapKeys()[0]
		ev := tv.MapIndex(ek)
		rv = ToAny(ev.Interface())
	}
	return
}

// todo
// LastofG[string](vx)
func LastofG[VT any, T map[any]VT | []VT | string](vx T) (rv VT) {
	// var s string = LastofG([]string{}) // cannot infer VT
	return
}

func Randof(vx any) (rv Any) {
	if Lenof(vx) <= 0 {
		return
	}
	tv := reflect.ValueOf(vx)
	ty := tv.Type()

	switch ty.Kind() {
	case reflect.Slice, reflect.Array:
		ev := tv.Index(mrand.Int() % tv.Len())
		rv = ToAny(ev.Interface())

	case reflect.String:
		ev := tv.Index(mrand.Int() % tv.Len())
		rv = ToAny(ev.Interface())

	case reflect.Map:
		ek := tv.MapKeys()[mrand.Int()%tv.Len()]
		ev := tv.MapIndex(ek)
		rv = ToAny(ev.Interface())
	}
	return

}
