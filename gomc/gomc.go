package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kitech/gopp"
)

// h1:xxx 的值，
// cat ~/go/pkg/mod/cache/download/github.com/kitech/gopp/\@v/v0.0.0-20240512142727-5384247fe530.ziphash
// go.sum line format: pkggetpath version ziphash

// $GOPATH/pkg/mod
var moddir = os.Getenv("HOME") + "/go/pkg/mod"
var moddldir = moddir + "/cache/download"

var getpkg = "github.com/kitech/gopp"

func main() {
	if len(os.Args) > 1 {
		getpkg = gopp.ArrayLast(os.Args)
	}

	getpkgbase := filepath.Base(getpkg)
	getmoddir := moddir + "/" + filepath.Dir(getpkg)
	log.Println(getmoddir, getpkgbase)

	ets, err := os.ReadDir(getmoddir)
	gopp.ErrPrint(err)
	log.Println(ets)

	outx := gopp.Mapdo(ets, func(idx int, vx any) []any {
		log.Println(idx, vx)
		v := vx.(fs.DirEntry)
		if v.Name() == getpkgbase {
			// more subpkg
			return nil
		}
		if !v.IsDir() {
			log.Println("wt", v.Name(), v.Type())
			return nil
		}
		if !strings.Contains(v.Name(), "@v") {
			log.Println("unknown format", v.Name())
			return nil
		}
		pkgver := strings.Split(v.Name(), "@")[1]

		pkgdlzip := moddldir + "/" + getpkg + "/@v/" + pkgver + ".ziphash"
		bcc, err := os.ReadFile(pkgdlzip)
		gopp.ErrPrint(err, pkgdlzip)
		pkghash := string(bcc)
		log.Println(string(bcc))

		return []any{pkgver, pkghash}
	})
	log.Println("versions of", getpkg, ":", outx)

	outx2 := gopp.Mapdo(outx, func(idx int, vx any) []any {
		if idx%2 == 1 {
			return nil
		}
		ziphash := outx.([]any)[idx+1]
		gosumline := fmt.Sprintf("%v %v %v", getpkg, vx, ziphash)
		log.Println(idx, "go.sum line,", gosumline)
		return []any{gosumline}
	})
	log.Println(outx2)
}
