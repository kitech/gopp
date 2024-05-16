// means go module cache offline tool
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kitech/gopp"

	"golang.org/x/mod/modfile"
	// "golang.org/x/mod/module"
	// "golang.org/x/mod/semver"
)

// h1:xxx 的值，
// cat ~/go/pkg/mod/cache/download/github.com/kitech/gopp/\@v/v0.0.0-20240512142727-5384247fe530.ziphash
// go.sum line format: pkggetpath version ziphash

// $GOPATH/pkg/mod
var moddir = os.Getenv("HOME") + "/go/pkg/mod"
var moddldir = moddir + "/cache/download"

var getpkg = "github.com/kitech/gopp"
var verbose bool
var showhelp bool

// cmd
// get pkgpath, // like go get, but offline first, then online
// list regexp, // only current workdir go.mod
// search regexp // in cachedir, downloaded
func main() {
	defer func() { log.Println("Used", time.Since(gopp.StartTime)) }()
	flag.BoolVar(&verbose, "v", true, "verbose")
	flag.BoolVar(&showhelp, "h", false, "help")
	flag.Parse()

	subcmd := flag.Arg(0)
	curpkg := flag.Arg(1)
	log.Println(subcmd, curpkg)
	// flag.PrintDefaults()

	if showhelp || subcmd == "" {
		flag.Usage()
		return
	}

	switch subcmd {
	case "get", "ge", "g":
	case "list", "lst", "l", "li":
		mclist(curpkg)
	case "listcache", "lc":
		mclistcache()
	case "search", "se":
	case "update", "up":
	case "updateall", "upall", "upa":
	case "delete", "del":
	}
}

func mclistcache() {
	var zipdirs []string
	filepath.WalkDir(moddldir, func(path string, d fs.DirEntry, err error) error {
		// log.Println(path)
		ets, err := os.ReadDir(path)
		if gopp.ErrHave(err, "not a directory") {
			// wtf,
		} else {
			gopp.ErrPrint(err)
		}

		rv := gopp.Mapdo(ets, func(vx any) []any {
			v := vx.(fs.DirEntry)
			// log.Println(v)
			if strings.HasSuffix(v.Name(), ".ziphash") && !v.IsDir() {
				// return []any{path, v.Name()}
				return gopp.Retn(path, v.Name())
			}
			return nil
		})
		gopp.TruePrint(gopp.Lenof(rv) > 0 && false, gopp.Lenof(rv), gopp.Capof(rv), rv)
		if gopp.Lenof(rv) > 0 {
			sv := gopp.IV2Strings(rv.([]any))
			zipdirs = append(zipdirs, sv...)
		}

		return nil
	})

	modvers := map[string][]string{}
	//log.Println(len(zipdirs), zipdirs)
	gopp.Mapdo(zipdirs, func(idx int, vx any) []any {
		if idx%2 == 1 {
			return nil
		}
		fpath, modverx := zipdirs[idx], zipdirs[idx+1]
		// log.Println(fpath, modverx)
		mpath := fpath[len(moddldir)+1 : len(fpath)-3] // sfx: /@v
		modver := modverx[:len(modverx)-8]             // sfx: .ziphash
		// log.Println(idx/2, mpath, modver)
		if _, ok := modvers[mpath]; ok {
			modvers[mpath] = append(modvers[mpath], modver)
		} else {
			modvers[mpath] = []string{modver}
		}
		return nil
	})

	gopp.Mapdo(modvers, func(idx int, kx any, vx any) []any {
		log.Printf("%d/%d vc.%d %v %v\n", idx, len(modvers), gopp.Lenof(vx), kx, vx)
		return nil
	})
}

func mclist(wkdir string) {
	gomodf := "./go.mod"
	gomodf = gopp.IfElseStr(wkdir == "", gomodf, wkdir+"/go.mod")
	// log.Println(gomodf)

	bcc, err := os.ReadFile(gomodf)
	gopp.ErrPrint(err, gomodf)
	if err != nil {
		log.Println("go.mod file not found", wkdir)
		return
	}
	mfo, err := modfile.Parse("", bcc, nil)
	gopp.ErrPrint(err, len(bcc))

	// log.Println(mfo.Module.Syntax)
	if len(mfo.Replace) > 0 {
		log.Println("Replaces...", len(mfo.Replace))
	}
	for i, rep := range mfo.Replace {
		log.Println(i, rep.Old, rep.Old.Version, "=>", rep.New, rep.New.Version)
	}

	if len(mfo.Require) > 0 {
		log.Println(mfo.Module.Mod.Path, "Requires...", len(mfo.Require))
	}
	for i, rep := range mfo.Require {
		log.Println(i, rep.Mod.Path, rep.Mod.Version, rep.Indirect)
	}

}

func mainttt() {

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

	log.Println("======")
	fgomod := "./go.mod"
	bcc, err := os.ReadFile(fgomod)
	gopp.ErrPrint(err, fgomod)
	mfo, err := modfile.Parse("./go.mod", bcc, nil)
	gopp.ErrPrint(err)
	log.Println(mfo, mfo.Require)

	mfo.AddNewRequire("hehhe/hehhe222", "v0.0.0", true)
	err = mfo.AddRequire("hehhe/hehhe", "v0.0.0")
	gopp.ErrPrint(err)
	bcc, err = mfo.Format()
	gopp.ErrPrint(err)
	log.Println(string(bcc)) // ok
}
