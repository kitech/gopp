package gopp

import (
	"reflect"
	"runtime"
	"strings"
)

// no use???
// usage: Progn(Retn(f1(...)), Retn(f2(...)), Retn(f3(...)))
func Progn(args ...interface{}) (rets [][]interface{}) {
	for _, arg := range args {
		rets = append(rets, arg.([]interface{}))
	}
	return
}

// 支持>=1个参数的函数
func Retn(args ...any) []any { return args }

// 最后一个参数为 error 类型，则检测err的值
// 返回除最后一个 error 参数外的其他参数，顺序不变
// Usage: Mustify(foo())
func Mustify(args ...any) (retx []Any) {
	for _, argx := range args {
		if arg, ok := argx.(error); ok {
			ErrPrint(arg)
		} else {
			retx = append(retx, AnyOf(argx))
		}
	}
	return
}

// 除了error之外，还有0个返回值
func Mustify0(err error) { Mustify(err) }

// 除了error之外，还有1个返回值
func Mustify1[T any](arg T, err error) T {
	return Mustify(arg, err)[0].I.(T)
}

// 除了error之外，还有2个返回值
func Mustify2[T1 any, T2 any](arg1 T1, arg2 T2, err error) (T1, T2) {
	retx := Mustify(arg1, arg2, err)
	return retx[0].I.(T1), retx[1].I.(T2)
}

// 对于没有返回值的函数，不能作为函数参数传递，所以无法用Progn函数调用。

type Retval struct {
	Ret   interface{}
	Err   error
	Extra interface{}
}

func CallFuncx(fx any, args ...any) {
	fv := reflect.ValueOf(fx)
	argv := make([]reflect.Value, len(args))
	for i, argx := range args {
		argv[i] = reflect.ValueOf(argx)
	}
	outs := fv.Call(argv)
	GOUSED(outs)
}

func MyFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	if pos := strings.LastIndex(name, "/"); pos != -1 {
		name = name[pos+1:]
	}

	return name
}
