package cgopp

import (
	"strings"

	"github.com/kitech/gopp"
)

// this this public use

func DyldImagesOtool(file string) (rets []string) {
	myexe := file
	lines, err := gopp.RunCmd(".", "otool", "-L", myexe)
	gopp.ErrPrint(err)
	// log.Println(lines, len(lines))
	gopp.Mapdo(lines, func(vx any) any {
		libfile := gopp.FirstofGv(strings.Split(vx.(string), " "))
		libfile = strings.Trim(libfile, "\t :")
		// log.Println(libfile)
		rets = append(rets, libfile)
		return nil
	})
	// panic("not impl")
	return
}

// todo
// by ldd on linux
// by depends on windows

func DyldSymbolsNm(file string) (rets []string) {
	myexe := file
	lines, err := gopp.RunCmd(".", "nm", myexe)
	gopp.ErrPrint(err)
	// log.Println(lines, len(lines))
	gopp.Mapdo(lines, func(vx any) any {
		libfile := gopp.LastofGv(strings.Split(vx.(string), " "))
		libfile = strings.Trim(libfile, "\t :")
		// log.Println(libfile)
		rets = append(rets, libfile)
		return nil
	})
	return
}
