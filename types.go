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
type f80 = [10]byte
type cfnadrty = *[0]byte
type Cfnadrty = *[0]byte

// type i128 = [16]uint8
// type u128 = [16]byte
type U128st struct{ H, L u64 }
type Fatf64 struct{ H, L f64 }

// exported
type Vptr = unsafe.Pointer
type Usize = uintptr
type Fatptr struct{ H, L usize }
type Quadptr struct{ H0, H1, L0, L1 usize }
type Fatany = U128st

func FatptrAs[T any](v Fatptr) (rv T) {
	p := (*T)(voidptr(&v))
	rv = *p
	return
}
func FatptrOf[T any](vx T) (rv Fatptr) {
	n := Copyx(&rv, &vx)
	// log.Println(vx, n)
	FalsePrint(usize(n) == unsafe.Sizeof(vx), "copy error", n, unsafe.Sizeof(vx))
	return
}

// 用于结构体和基本数据类型
// 拷贝长度按照 src 的长度计算，确保dst足够
func Copyx[DT any, ST any](dst *DT, src *ST) int {
	dlen := unsafe.Sizeof(*src)

	slc1 := GoSlice{(voidptr)(dst), int(dlen), int(dlen)}
	b1 := *(*[]byte)(unsafe.Pointer(&slc1))

	slc2 := GoSlice{(voidptr)(src), int(dlen), int(dlen)}
	b2 := *(*[]byte)(unsafe.Pointer(&slc2))

	return copy(b1, b2)
}

// todo
func (v Fatptr) Asany(ty reflect.Kind) (rv any) {
	switch ty {
	case reflect.Float64:
		rv = FatptrAs[float64](v)
	}
	return
}
func (v Fatptr) Ptr() (rv voidptr) {
	rv = FatptrAs[voidptr](v)
	return
}
func (v Fatptr) Int() (rv int) {
	rv = FatptrAs[int](v)
	return
}
func (v Fatptr) Int64() (rv int64) {
	rv = FatptrAs[int64](v)
	return
}
func (v Fatptr) Bool() (rv bool) {
	rv = FatptrAs[bool](v)
	return
}
func (v Fatptr) Float64() (rv float64) {
	rv = FatptrAs[float64](v)
	return
}

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
func (this Any) Bool() bool   { return this.I.(bool) }
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
var vInt16Ty int16
var vUint16Ty uint16
var vInt32Ty int32
var vUint32Ty uint32
var vInt64Ty int64
var vUint64Ty uint64
var vByteTy byte
var vFloat32Ty float32
var vFloat64Ty float64
var vBoolTy bool
var vStrTy string
var vUintptrTy uintptr
var vVoidptrTy unsafe.Pointer

var InvalidTy = reflect.TypeOf(nil)
var Int8Ty = reflect.TypeOf(vInt8Ty)
var Uint8Ty = reflect.TypeOf(vUint8Ty)
var IntTy = reflect.TypeOf(vIntTy)
var UintTy = reflect.TypeOf(vUintTy)
var Int16Ty = reflect.TypeOf(vInt16Ty)
var Uint16Ty = reflect.TypeOf(vUint16Ty)
var Int32Ty = reflect.TypeOf(vInt32Ty)
var Uint32Ty = reflect.TypeOf(vUint32Ty)
var Int64Ty = reflect.TypeOf(vInt64Ty)
var Uint64Ty = reflect.TypeOf(vUint64Ty)
var ByteTy = reflect.TypeOf(vByteTy)
var Float32Ty = reflect.TypeOf(vFloat32Ty)
var Float64Ty = reflect.TypeOf(vFloat64Ty)
var BoolTy = reflect.TypeOf(vBoolTy)
var StrTy = reflect.TypeOf(vStrTy)
var UsizeTy = reflect.TypeOf(vUintptrTy)
var VoidpTy = reflect.TypeOf(vVoidptrTy)

var RefKindTys = map[reflect.Kind]reflect.Type{
	reflect.Invalid: InvalidTy,
	reflect.Int8:    Int8Ty, reflect.Uint8: Uint8Ty,
	reflect.Int: IntTy, reflect.Uint: UintTy,
	reflect.Int16: Int16Ty, reflect.Uint16: Uint16Ty,
	reflect.Int32: Int32Ty, reflect.Uint32: Uint32Ty,
	reflect.Int64: Int64Ty, reflect.Uint64: Uint64Ty,
	reflect.Float32: Float32Ty, reflect.Float64: Float64Ty,
	reflect.Bool: BoolTy, reflect.String: StrTy,
	reflect.Uintptr:       UintptrTy,
	reflect.UnsafePointer: VoidpTy,
}
var RefKindVals = map[reflect.Kind]reflect.Value{
	reflect.Invalid:       reflect.ValueOf(nil),
	reflect.Int8:          reflect.ValueOf(vInt8Ty),
	reflect.Uint8:         reflect.ValueOf(vUintTy),
	reflect.Int:           reflect.ValueOf(vIntTy),
	reflect.Uint:          reflect.ValueOf(vUintTy),
	reflect.Int16:         reflect.ValueOf(vInt16Ty),
	reflect.Uint16:        reflect.ValueOf(vUint16Ty),
	reflect.Int32:         reflect.ValueOf(Int32Ty),
	reflect.Uint32:        reflect.ValueOf(Uint32Ty),
	reflect.Int64:         reflect.ValueOf(Int64Ty),
	reflect.Uint64:        reflect.ValueOf(Uint64Ty),
	reflect.Float32:       reflect.ValueOf(Float32Ty),
	reflect.Float64:       reflect.ValueOf(Float64Ty),
	reflect.Bool:          reflect.ValueOf(vBoolTy),
	reflect.String:        reflect.ValueOf(ZeroStr),
	reflect.Uintptr:       reflect.ValueOf(vUintptrTy),
	reflect.UnsafePointer: reflect.ValueOf(vVoidptrTy),
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
var UintptrTy = reflect.TypeOf(&vUintptrTy)

const ByteTySz = unsafe.Sizeof(byte(0))
const Int8TySz = strconv.IntSize / 4
const Int16TySz = Int8TySz * 2
const Int32TySz = Int8TySz * 4
const Int64TySz = Int8TySz * 8
const Float64TySz = Int8TySz * 8
const Float32TySz = Int8TySz * 4
const UintptrTySz = unsafe.Sizeof(uintptr(0))
const UsizeTySz = unsafe.Sizeof(uintptr(0))
const BoolTySz = unsafe.Sizeof(true)
const IntTySz = unsafe.Sizeof(int(0))
const RuneTySz = unsafe.Sizeof(rune(0))
const EmptyStructTySz = unsafe.Sizeof(struct{}{})

// const NilSz = unsafe.Sizeof(nil)

func IsMap(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func IsArray(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Array
}

func IsSlice(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsChan(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Chan
}

func IsFunc(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func IsStruct(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func IsPtr(v any) bool {
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

// safe mode
func ValueAt[T any](vx []T, idx int) T {
	var rv T
	if idx < len(vx) {
		rv = vx[idx]
	}
	return rv
}

// 可以避免nil check
// 还是喜欢这种写法的！
// Lastof(vx).Str()
func Lastof(vx any) (rv Any) {
	return lastorfirstof(false, vx)
}

// 可以避免nil check
// 还是喜欢这种写法的！
// Firstof(vx).Str()
func Firstof(vx any) (rv Any) {
	return lastorfirstof(true, vx)
}

// 还是喜欢这种写法的！
// Lastof(vx).Str()
func lastorfirstof(first bool, vx any) (rv Any) {
	if vx == nil || Lenof(vx) <= 0 {
		return ToAny(0)
	}
	tv := reflect.ValueOf(vx)
	ty := tv.Type()

	idx := IfElse2(first, 0, Lenof(vx)-1)
	switch ty.Kind() {
	case reflect.Slice, reflect.Array:
		ev := tv.Index(idx)
		rv = ToAny(ev.Interface())

	case reflect.String:
		ev := tv.Index(idx)
		rv = ToAny(ev.Interface())

	case reflect.Map:
		ek := tv.MapKeys()[idx]
		ev := tv.MapIndex(ek)
		rv = ToAny(ev.Interface())
	}
	return
}

// todo, LastofG(vx)，这种写法go推导不出来！！！
// Usage: LastofG[string](vx)
func LastofG[VT any, T map[any]VT | []VT | string](vx T) (rv VT) {
	// var s string = LastofG([]string{}) // cannot infer VT
	var vxty = reflect.TypeOf(vx)
	switch vxty.Kind() {
	case reflect.Slice:
		n := Lenof(vx)
		if n > 0 {
			rv = reflect.ValueOf(vx).Index(n - 1).Interface().(VT)
		}
	case reflect.Map:
		n := Lenof(vx)
		if n > 0 {
			vo := reflect.ValueOf(vx)
			keys := vo.MapKeys()
			rv = vo.MapIndex(keys[len(keys)-1]).Interface().(VT)
		}
	case reflect.String:
		n := Lenof(vx)
		if n > 0 {

		}
	}
	return
}
func LastofGv[VT any](vx []VT) (rv VT) {
	if len(vx) > 0 {
		rv = vx[len(vx)-1]
	}
	return
}
func FirstofGv[VT any](vx []VT) (rv VT) {
	if len(vx) > 0 {
		rv = vx[0]
	}
	return
}
func RandofGv[VT any](vx []VT) (rv VT) {
	if len(vx) > 0 {
		rv = vx[mrand.Int()%len(vx)]
	}
	return
}
func LastofGm[T any](vx map[any]T) (rv T) {
	if len(vx) > 0 {
		for _, v := range vx {
			rv = v
			return
		}
	}
	return
}
func FirstofGm[T any](vx map[any]T) (rv T) {
	if len(vx) > 0 {
		for _, v := range vx {
			rv = v
			return
		}
	}
	return
}
func LastofGs(vx string) (rv byte) {
	if len(vx) > 0 {
		rv = vx[len(vx)-1]
	}
	return
}
func FirstofGs(vx string) (rv byte) {
	if len(vx) > 0 {
		rv = vx[0]
	}
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
