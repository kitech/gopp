package cgopp

import (
	"log"
	"runtime"
	"testing"
)

func TestGetg0(t *testing.T) {
	// now works???
	// Undefined symbols for architecture x86_64: "_runtime.getg"
	// g := getg()
	// log.Println(g)

	// works
	x := acquirem()
	defer releasem(x)
	log.Println(x)

}

func rtmoddatadump(mod *moduledata) {
	log.Println(voidptr(mod), mod.hasmain, mod.pluginpath, GoStringN(voidptr(mod.text), 26))
	log.Println(voidptr(mod), mod.minpc, mod.maxpc, mod.maxpc-mod.minpc)

}

func TestRtfilefunc(t *testing.T) {
	pcs := make([]usize, 32)
	n := runtime.Callers(0, pcs)
	// log.Println(n, len(pcs))
	pcs = pcs[:n]
	frms := runtime.CallersFrames(pcs)
	log.Println(frms, len(pcs))

	mods := map[voidptr]int{}
	for i, pc := range pcs {
		fi := Rtfindfunc(pc)
		log.Println(i, pc, fi)
		mod := (*moduledata)(fi.MD)
		if _, ok := mods[voidptr(mod)]; !ok {
			mods[voidptr(mod)] = 1
			rtmoddatadump(mod)
		}
		log.Println(i, Rtfuncpkgpath(fi), Rtfuncname(fi), Rtfuncfile(fi, 0))
	}
}

func TestMD0(t *testing.T) {
	prtmod := func(mod *moduledata) {
		for i := 0; mod != nil; mod = mod.next {
			log.Println(i, mod.modulename)
			i++
		}
	}
	{
		pcs := make([]usize, 32)
		n := runtime.Callers(0, pcs)
		log.Println(n, len(pcs))
		pcs = pcs[:n]
		frms := runtime.CallersFrames(pcs)
		log.Println(frms, len(pcs))
		for i, pc := range pcs {
			fi := Rtfindfunc(pc)
			log.Println(i, pc, fi)
			mod := (*moduledata)(fi.MD)
			prtmod(mod)
			log.Println(Rtfuncfile(fi, 0))
			log.Println(Rtfuncname(fi))
		}
	}
	log.Println(firstmoduledata)
	mod := firstmoduledata
	prtmod(&mod)
}
