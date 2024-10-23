package gopp

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"syscall"
	"testing"
	"time"
)

// TODO 要是侯选可以惰性求值就好了，否则在只能一个求值的场景则会有问题
// 简单的三元运算符模拟函数
func IfElse(q bool, tv interface{}, fv interface{}) interface{} {
	if q == true {
		return tv
	} else {
		return fv
	}
}

func IfElse2[T any](q bool, tv T, fv T) T {
	if q == true {
		return tv
	} else {
		return fv
	}
}

func IfElseInt(q bool, tv int, fv int) int {
	return IfElse2(q, tv, fv)
}

func IfElseStr(q bool, tv string, fv string) string {
	return IfElse2(q, tv, fv)
}

// 这个好像没有什么用
func IfThen2[T any](q bool, thens ...T) (v T) {
	if len(thens) > 0 {
		return thens[0]
	}
	return
}
func IfThen(q bool, thens ...interface{}) interface{} {
	if len(thens) > 0 {
		return thens[0]
	}
	return nil
}

// for number type, uintptr
func CmpAndSwapN(src interface{}, old interface{}, new interface{}) (swapped bool) {
	srcv := reflect.ValueOf(src)
	if srcv.Type().Kind() != reflect.Ptr {
		return
	}

	oldv := reflect.ValueOf(old)
	newv := reflect.ValueOf(new)

	if reflect.DeepEqual(srcv.Interface(), oldv) {
		srcv.Elem().Set(newv.Convert(srcv.Type()))
		swapped = true
	}

	return
}

// 把一个值转换为数组切片
// 如果本身即为数组切片，则显式转换为数组类型
// 如果本身不是数组切片，则把该值作为返回数组切片的第一个值。
func ToSlice(v interface{}, reverse bool) []interface{} {
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Slice {
		res := []interface{}{}
		vv := reflect.ValueOf(v)
		for i := 0; i < vv.Len(); i++ {
			idx := IfElse(reverse, vv.Len()-i-1, i).(int)
			res = append(res, vv.Index(idx).Interface())
		}
		return res
	} else {
		return []interface{}{v}
	}
}

// v express of bool, string, pointer
func Assertf(v interface{}, format string, args ...interface{}) {
	Assert(v, fmt.Sprintf(format, args...))
}

// v express of bool, string, pointer
func Assert(v interface{}, info string, args ...interface{}) {
	fmtv := fmt.Sprintf("%+v, %+v", v, info)
	for _, arg := range args {
		fmtv += fmt.Sprintf(", %+v", arg)
	}
	if v == nil {
		panic(fmtv)
	}

	tv := reflect.TypeOf(v)

	vv := reflect.ValueOf(v)
	switch tv.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uint8, reflect.Int8:
		if vv.Int() == 0 {
			panic(fmtv)
		}
	case reflect.String:
		if v.(string) == "" {
			panic(fmtv)
		}
	case reflect.Bool:
		if v.(bool) == false {
			panic(fmtv)
		}
	}
}

func AssertT() {

}

// 俩工具
// 直接忽略掉变量未使用编译提示
func GOUSED(vars ...interface{})   {}
func GOUNUSED(vars ...interface{}) {}
func G_USED(vars ...interface{})   {}
func G_UNUSED(vars ...interface{}) {}
func G_FATAL(err error) {
	if err != nil {
		panic(err)
	}
}

// 去掉返回值中的error
// 返回值个数不能是变长的
func NOE(v ...interface{}) interface{} {
	n := len(v)
	if n == 0 {
		return nil
	}
	last := v[n-1]
	lt := reflect.TypeOf(last)

	e := errors.New("dummy")

	if lt.Kind() == reflect.TypeOf(e).Kind() {
	}
	return v
}

func WAITIF(condfn func() bool, msec int) {
	for {
		if condfn() {
			break
		}
		time.Sleep(time.Duration(msec) * time.Microsecond)
	}
}

func FileExist(fname string) bool {
	if _, err := os.Stat(fname); err != nil {
		if err.(*os.PathError).Err == syscall.ENOENT {
			return false
		}
	}
	return true
}

// exists returns whether the given file or directory exists or not
func FileExist2(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Iface2Value(args []interface{}) []reflect.Value {
	if args == nil {
		return nil
	}

	vals := make([]reflect.Value, 0)
	for _, arg := range args {
		vals = append(vals, reflect.ValueOf(arg))
	}
	return vals
}

func Value2Iface(vals []reflect.Value) []interface{} {
	if vals == nil {
		return nil
	}

	rets := make([]interface{}, 0)
	for _, val := range vals {
		rets = append(rets, val.Interface())
	}
	return rets
}

func PackArgs(args ...any) string {
	sb := strings.Builder{}
	for i, arg := range args {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprintf("%+v", arg))
	}
	return sb.String()
}

func InTest() bool {
	exepath, err := os.Executable()
	if err != nil {
		// wtf
	}
	if strings.HasSuffix(exepath, ".test") {
		// seem good
	}
	return testing.Short()
}

// 好像 reflect.Value.IsZero等于这个函数要实现的功能，并且包括所有的类型包括结构体
// string, int*, float*, pointer, map, slice, chan
// struct == all empty
func Empty(vx any) (bv bool) {
	ty := reflect.TypeOf(vx)
	val := reflect.ValueOf(vx)

	if true {
		// 试试，不支持的类型会panic
		return vx == nil || reflect.ValueOf(vx).IsZero()
	}

	switch ty.Kind() {
	case reflect.Slice, reflect.Array:
		bv = val.IsNil() || val.Len() == 0
	case reflect.Map:
		bv = val.IsNil() || val.Len() == 0
	case reflect.Chan:
		bv = val.IsNil() || val.Len() == 0
	case reflect.String:
		bv = val.Len() == 0
	case reflect.Pointer, reflect.UnsafePointer:
		bv = val.IsNil()
	case reflect.Struct:
		// todo
		bv = val.IsZero()
	default:
		if Isnumtype(ty) {
			bv = val.Equal(reflect.Zero(ty))
		} else {

		}
	}
	return
}
