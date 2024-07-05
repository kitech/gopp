package gopp

import (
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
func Retn(args ...interface{}) []interface{} { return args }

// 对于没有返回值的函数，不能作为函数参数传递，所以无法用Progn函数调用。

type Retval struct {
	Ret   interface{}
	Err   error
	Extra interface{}
}

func MyFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	if pos := strings.LastIndex(name, "/"); pos != -1 {
		name = name[pos+1:]
	}

	return name
}
