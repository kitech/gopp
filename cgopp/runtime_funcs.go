package cgopp

import (
	"runtime"
	"time"

	"github.com/kitech/gopp"
)

// get go func by name, just like dlsym
//
//export Dlsymgo
func Dlsymgo(name string) usize {
	var fnptr usize
	var cnt = 0
	var btime = time.Now()
	GortWalkFuncs(func(fo *runtime.Func) bool {
		cnt++
		// log.Println(cnt, name, "??", fo.Name())
		if fo.Name() == name {
			fnptr = fo.Entry()
			return true
		}
		return false
	})
	gopp.ZeroPrint(fnptr, name, "not found after", cnt, time.Since(btime))
	return fnptr
}

func Gortmd0() *moduledata { return &firstmoduledata }

// cbfn return true stop
// see https://github.com/alangpierce/go-forceexport/blob/master/forceexport.go
func GortWalkFuncs(cbfn func(*runtime.Func) bool) {
	mod0 := &firstmoduledata
	for md := mod0; md != nil; md = md.next {
		for _, ftab := range md.ftab {
			f := (*runtime.Func)(voidptr(&md.pclntable[ftab.funcoff]))
			if ok := cbfn(f); ok {
				return
			}
		}
	}
}
