package cgopp

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ebitengine/purego"
	"github.com/kitech/gopp"
)

func TestFC1(t *testing.T) {
	// TestLitffi2callz()
	// TestLitffi3callz()
}

func TestFCBM2(t *testing.T) {
	// BMLitffi2callz()
}

func TestFC3BM3(t *testing.T) {
	// dlh := testfc3_setup()
	// defer purego.Dlclose(dlh)

	// BMLitffi3callz(dlh)
	// BMLitffi3callz2(dlh)
}

func TestFC3aaa(t *testing.T) {
	dlh := testfc3_setup()
	defer purego.Dlclose(dlh)

	TestLitffi3callz(dlh)
	// bench
	BMLitffi3callz(dlh)
	BMLitffi3callz2(dlh)
}

var litffi3_test_code = `
//
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>

// int
double
litffi3_test1(double a, void*b, int64_t c) {
	// printf("%s: %f, %d, %p=%ld, %lld, \n", __FUNCTION__, a, (int)a, b, (uintptr_t)b, c);
	return a;
    return (int)(a);
}

float
litffi3_test2(float a) {
	printf("%s: %f\n", __FUNCTION__, a);
	return a+1;
}
`

func testfc3_setup() usize {
	log.Println("wkdir:", gopp.Retn(os.Getwd()))
	log.Println("exefile", gopp.Retn(os.Executable()))
	exefile, err := os.Executable()
	gopp.ErrPrint(err)
	exedir := filepath.Dir(exefile)
	tmpname := gopp.RandStrHex(9)
	tmpcfile := exedir + "/" + tmpname + ".c"
	tmplibfile := exedir + "/lib" + tmpname + ".dylib"
	log.Println("exedir:", exedir)

	gopp.SafeWriteFile(tmpcfile, []byte(litffi3_test_code), os.ModePerm)
	defer os.Remove(tmpcfile)
	log.Println("tmpcfile:", gopp.FileExist2(tmpcfile), tmpcfile)

	cccmd := []string{"cc", "-g -O0 -shared -o", tmplibfile, tmpcfile}
	err = gopp.RunCmdSout(nil, exedir, cccmd...)
	gopp.ErrPrint(err)
	defer os.Remove(tmplibfile)
	log.Println("libfile:", gopp.FileExist2(tmplibfile), tmplibfile)

	cccmd = []string{"nm", tmplibfile}
	err = gopp.RunCmdSout(nil, exedir, cccmd...)
	gopp.ErrPrint(err)
	err = gopp.RunCmdSout(nil, exedir, "file", tmplibfile)
	gopp.ErrPrint(err)

	log.Println("Setup done, run test now...")
	dlh, err := purego.Dlopen(tmplibfile, purego.RTLD_LAZY)
	gopp.ErrPrint(err)
	log.Println("dlopen:", dlh)

	return dlh
}
