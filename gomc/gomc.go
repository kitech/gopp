// means go module cache offline tool
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
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
var dltmpdir = os.Getenv("HOME") + "/bprog/dmgoget"

var getpkg = "github.com/kitech/gopp"
var verbose bool
var showhelp bool

// cmd
// get pkgpath, // like go get, but offline first, then online
// list regexp, // only current workdir go.mod
// search regexp // in cachedir, downloaded
// dltmp pkgpath, // it's go get, but in our dltmpdir, not current dir
func main() {

	log.SetFlags(log.Flags() ^ log.Ldate ^ log.Ltime)
	defer func() { log.Println("Used", time.Since(gopp.StartTime)) }()
	flag.BoolVar(&verbose, "v", true, "verbose")
	flag.BoolVar(&showhelp, "h", false, "help")
	flagset := flag.NewFlagSet("gomc get", flag.ExitOnError)
	flag.Parse()

	subcmd := flag.Arg(0)
	curpkg := flag.Arg(1)
	log.Println("try run", subcmd, curpkg, "...")
	// flag.PrintDefaults()

	if showhelp || subcmd == "" {
		log.Println("go module cache offline tool")
		flag.Usage()
		flagset.Usage()
		log.Println("    subcmds: get, list, listcache, search, delete")
		return
	}

	switch subcmd {
	case "get", "ge", "g":
		if len(curpkg) == 0 {
			log.Println("must supply getpkg")
			break
		}
		mcget(curpkg)
	case "list", "lst", "l", "li":
		mclist(curpkg)
	case "listcache", "lc":
		mclistcache()
	case "search", "se":
		if len(curpkg) == 0 {
			log.Println("must supply word")
			break
		}
		mcsearch(curpkg)
	case "update", "up":
	case "updateall", "upall", "upa":
	case "delete", "del":
		if len(curpkg) == 0 {
			log.Println("must supply getpkg")
			break
		}
		mcdelete(curpkg)

		// dummy get a package in any directory, but not current dir
		// dont, change current dir's go.mod/go.sum
	case "dmget":
	case "dminst": // go install package@latest
	// todo 同步本地开发目录到 mod cache 目录的zip文件，完全离线
	case "dltmp":
		mcdltmpget(curpkg)
	case "locsync":
	case "fakelocalproxyserver": // todo
	default:
		log.Println("subcmd not found", subcmd)
	}
}

func mcdelete(pkgpath string) {
	res := mclistcacheall(pkgpath)
	if len(res) == 0 {
		// todo, string similarity
		// https://github.com/adrg/strutil
		log.Println("package not found in cache", pkgpath)
		return
	} else if len(res) > 1 {
		log.Println("too many packages found in cache", res)
		return
	}
	gopp.Mapdo(res, func(idx int, kx, vx any) []any {
		log.Printf("found %d/%d vc.%d %v %v\n", idx, len(res), gopp.Lenof(vx), kx, vx)
		return nil
	})

	bcc, err := os.ReadFile("go.mod")
	gopp.ErrPrint(err)
	mfo, err := modfile.Parse("", bcc, nil)
	gopp.ErrPrint(err)

	for modpath, modvers := range res {
		keys := gopp.MapKeys(modvers)
		slices.Sort(keys)
		modver := gopp.Lastof(keys).Str()

		err := mfo.DropRequire(modpath)
		gopp.ErrPrint(err)

		{
			bcc2, err := mfo.Format()
			gopp.ErrPrint(err)
			if string(bcc2) == string(bcc2) {
				log.Println("no change", "go.mod not have", modpath)
			} else {
				err = gopp.SafeWriteFile("go.mod", bcc, 0755)
				gopp.ErrPrint(err)
			}
		}

		bcc, err := os.ReadFile("go.sum")
		gopp.ErrPrint(err)

		lines := strings.Split(string(bcc), "\n")
		newline := fmt.Sprintf("%s %s %s", modpath, modver, modvers[modver])
		if idx := slices.Index(lines, newline); idx >= 0 {
			lines = slices.Delete(lines, idx, idx+1)

			scc := strings.Join(lines, "\n")
			err = gopp.SafeWriteFile("go.sum", []byte(scc), 0755)
			gopp.ErrPrint(err)
		} else {
			log.Println("no change", "go.sum not have", newline)
		}

		break
	}

}

// will write file go.mod, go.sum
func mcget(pkgpath string) {
	res := mclistcacheall(pkgpath)
	if len(res) == 0 {
		// todo, string similarity
		// https://github.com/adrg/strutil
		log.Println("package not found in cache", pkgpath)
		return
	} else if len(res) > 1 {
		// exact equal
		res2 := map[string]map[string]string{}
		gopp.Mapdo(res, func(i int, kx, vx any) map[any]any {
			if kx.(string) == pkgpath {
				res2[kx.(string)] = vx.(map[string]string)
				// return map[any]any{kx.(string): vx.(map[string]string)}
				return nil
			}
			return nil
		})
		log.Println(res2)
		if gopp.Lenof(res2) > 1 {
			log.Println("too many packages found in cache", res2)
			return
		}
		res = res2
	}
	gopp.Mapdo(res, func(idx int, kx, vx any) []any {
		log.Printf("found %d/%d vc.%d %v %v\n", idx, len(res), gopp.Lenof(vx), kx, vx)
		return nil
	})

	bcc, err := os.ReadFile("go.mod")
	gopp.ErrPrint(err)
	mfo, err := modfile.Parse("", bcc, nil)
	gopp.ErrPrint(err)

	for modpath, modvers := range res {
		keys := gopp.MapKeys(modvers)
		slices.Sort(keys)
		modver := gopp.Lastof(keys).Str()
		err := mfo.AddRequire(modpath, modver)
		gopp.ErrPrint(err)

		{
			bcc, err := mfo.Format()
			gopp.ErrPrint(err)
			err = gopp.SafeWriteFile("go.mod", bcc, 0755)
			gopp.ErrPrint(err)
		}

		bcc, err := os.ReadFile("go.sum")
		gopp.ErrPrint(err)

		lines := strings.Split(string(bcc), "\n")
		newline := fmt.Sprintf("%s %s %s", modpath, modver, modvers[modver])
		if !slices.Contains(lines, newline) {
			lines = append(lines, newline)

			scc := strings.Join(lines, "\n")
			err = gopp.SafeWriteFile("go.sum", []byte(scc), 0755)
			gopp.ErrPrint(err)
		}

		break
	}
}

func mcsearch(word string) {
	pkgverhashs := mclistcacheall(word)
	res := map[string]map[string]string{}
	gopp.Mapdo(pkgverhashs, func(idx int, kx, vx any) []any {
		if !strings.Contains(kx.(string), word) {
			return nil
		}
		res[kx.(string)] = vx.(map[string]string)
		return nil
	})
	gopp.Mapdo(res, func(idx int, kx, vx any) []any {
		log.Printf("%d/%d vc.%d %v %v\n", idx, len(res), gopp.Lenof(vx), kx, vx)
		return nil
	})
}
func mclistcache() {
	res := mclistcacheall("")
	gopp.Mapdo(res, func(idx int, kx, vx any) []any {
		log.Printf("%d/%d vc.%d %v %v\n", idx, len(res), gopp.Lenof(vx), kx, vx)
		return nil
	})
}

// contains match, not exact match
func mclistcacheall(word string) (pkgvers map[string]map[string]string) {
	log.Println("walking", moddldir, "...")
	var zipdirs []string
	filepath.WalkDir(moddldir, func(path string, d fs.DirEntry, err error) error {
		// log.Println(path)
		tpath := path[len(moddldir):]
		if len(word) > 0 && !gopp.StrHaveNocase(tpath, word) {
			return nil
		}

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

	modvers := map[string]map[string]string{}
	//log.Println(len(zipdirs), zipdirs)
	gopp.Mapdo(zipdirs, func(idx int, vx any) []any {
		if idx%2 == 1 {
			return nil
		}
		fpath, modverx := zipdirs[idx], zipdirs[idx+1]
		bcc, err := os.ReadFile(fpath + "/" + modverx)
		gopp.ErrPrint(err, fpath, modverx)
		// log.Println(fpath, modverx)
		mpath := fpath[len(moddldir)+1 : len(fpath)-3] // sfx: /@v
		modver := modverx[:len(modverx)-8]             // sfx: .ziphash
		// log.Println(idx/2, mpath, modver)
		if _, ok := modvers[mpath]; ok {
			modvers[mpath][modver] = string(bcc)
		} else {
			modvers[mpath] = map[string]string{modver: string(bcc)}
		}
		return nil
	})

	pkgvers = modvers

	if true {
		return
	}
	gopp.Mapdo(modvers, func(idx int, kx any, vx any) []any {
		log.Printf("%d/%d vc.%d %v %v\n", idx, len(modvers), gopp.Lenof(vx), kx, vx)
		return nil
	})

	return
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

func mcdltmpget(curpkg string) {
	log.Println("Workdir:", dltmpdir)
	err := gopp.RunCmdSout(nil, dltmpdir, "go get -v", curpkg)
	gopp.ErrPrint(err, curpkg)
}
