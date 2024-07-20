package cgopp

import (
	"os"

	"github.com/kitech/gopp"
)

func DyldImagesSelf() (rets []string) {
	myexe := gopp.Mustify1(os.Executable())
	rets = DyldImages(myexe)
	// panic("not impl")
	return
}

func DyldSymbolsSelf() (rets []string) {
	myexe := gopp.Mustify1(os.Executable())
	rets = DyldSymbols(myexe)

	return
}

func DyldImages(file string) (rets []string) {
	switch dyldfuncs_whichimpl {
	case 1:
		rets = DyldImagesOtool(file)
	case 2:
		rets = DyldImagesInc(file)
	}
	return
}

const dyldfuncs_whichimpl = 1 // 1, cmd, 2, in c

func DyldSymbols(file string) (rets []string) {
	switch dyldfuncs_whichimpl {
	case 1:
		rets = DyldSymbolsNm(file)
	case 2:
		rets = DyldSymbolsInc(file)
	}
	return
}
