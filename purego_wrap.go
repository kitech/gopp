package gopp

import (
	"log"
	"sync"

	"github.com/ebitengine/purego"
)

type Purego struct {
	mu sync.Once
	sofile string
	Dlh usize
}

// "" for self executable
// just libxxx.so is fine, not include path
func PuregoNew(sofile string) *Purego{
	rv := &Purego{}
	rv.sofile = sofile
	return rv
}
func (pgo *Purego) init() (reterr error) {
	if pgo.sofile == "" {
		return
	}
	pgo.mu.Do(func() {
		dlh , err := purego.Dlopen(pgo.sofile, purego.RTLD_NOW)
		reterr = err
		ErrPrint(err, pgo.sofile)
		pgo.Dlh = dlh
	})
	return
}
func (pgo *Purego) Close() error {
	return purego.Dlclose(pgo.Dlh)
}

func (pgo *Purego) Sym(sym string) (usize, error) {
	pgo.init()
	return purego.Dlsym(pgo.Dlh, sym)
}

// usage f1 := func() ; RegisterFunc(&f1, "name")
func (pgo *Purego) RegisterFunc(fptr any, sym string) error {
	pgo.init()
	adr, err := purego.Dlsym(pgo.Dlh, sym)
	if err != nil { return err }
	purego.RegisterFunc(fptr, adr)
	return err
}

func dummy() {
	// gopp.Keep()
	if false {
		purego.Dlclose(0)
		log.Println(12345)
	}
}
