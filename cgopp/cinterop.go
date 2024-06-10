package cgopp

/*

// passive mode

*/
import "C"
import "github.com/kitech/gopp"

// 两种模式，
// go作为exe入口，go是主调方
// go作为库使用，go是被调方
// 第三种方式，使用dlsym方式，这样不需要注册函数指针，只需要配置函数名字
// 只要链接到一起，不需要关心调用的顺序问题。
// \see github.com/ebitengine/purego

// 两种传参模式，
// 简单的JSON字段中
// 兼容的结构体模式

// passive mode

type GointeropFntype func(voidptr) voidptr

var irpasvcbfns = map[any]voidptr{}

// this must call nongo scope,
// such as c/rs/cxx
//
//export gointerop_set_passive_cbfn
func gointerop_set_passive_cbfn(name4c *C.char, fnptr voidptr) {
	var name = C.GoString(name4c)
	irpasvcbfns[name] = fnptr
	gopp.NilPrint(fnptr, "cannot set nil fnptr")
}

func GoirGetPasvCbfn[T string | int](name T) voidptr {
	if fnptr, ok := irpasvcbfns[name]; ok {
		return fnptr
	}
	return nil
}

// active mode

// dlsym mode
