package main

import (
	"log"
	"os"
	"strings"

	"github.com/kitech/gopp"
	"github.com/kitech/gopp/cgopp"
	_ "github.com/kitech/gopp/cgopp"
)

/*

// #cgo CFLAGS: -I/nix/store/na82nfg5k2nw2kz6bfrsv447fwnhwh8z-zulu-ca-jdk-17.0.10/include

#include <stdlib.h>
#include <stdio.h>

*/
import "C"

//export Java_Main_goexpfn1
func Java_Main_goexpfn1() {
	log.Println("heree", C.int(0))
}

//export Java_Main_goexploop
func Java_Main_goexploop() {
	for {
		gopp.SleepSec(5)
		// log.Println("heree", C.int(0))
	}
}

//export Java_MainKt_goexpfn1
func Java_MainKt_goexpfn1() {
	log.Println("heree", C.int(0))
}

func main() {
	log.Println("runexe", runexe)
	cgopp.JavaExe("Main", "main")
}

var runexe = false // 0 lib, 1 exe
func init() {

	exe, _ := os.Executable()
	runexe = !strings.HasSuffix(exe, "/java")
	log.Println("heree", "runexe", true)

	// todo test cgopp.RunOnJvm
	// cgopp.JNI_OnLoad_Callback = func() {
	// 	log.Println(gopp.MyFuncName(), "tid", cgopp.MyTid())
	// 	// testfunc()
	// }
	if !runexe {
		go init2()
	}
}
func init2() {

	log.Println("looping...")
	for i := 0; ; i++ {
		gopp.SleepSec(1)
		go cgopp.RunOnJVM(func() {
			testfunc()
		})
		gopp.SleepSec(2)
	}
}

func testfunc() {
	log.Println("herere", cgopp.Jenv, "tid", cgopp.MyTid())
	cgopp.JNIThreadCheck()
	defer log.Println("herere", cgopp.Jenv, "eee=== top -pid", os.Getpid())
	je := cgopp.Jenv
	// cls4c := cgopp.CStringgc("LMain")
	ver := je.GetVersion()
	log.Println("jvm ver", ver, cgopp.JVMTid(), cgopp.MyTid())

	x := je.FindClass("Main")
	log.Println("jvcls Main", x)
	gopp.NilPrint(x, "Err")

	{
		y := je.GetStaticMethodID(x, "jvexpfn1", "()V")
		log.Println("jvfn1", y)
		gopp.NilPrint(y, "Err")

		je.CallStaticVoidMethod(x, y)
	}

	{
		y := je.GetStaticMethodID(x, "jvexpfn2", "(Ljava/lang/String;)V")
		log.Println("jvfn2", y)
		gopp.NilPrint(y, "Err")

		// cgopp.JNIEnvCallStaticMethod[cgopp.Void](je, x, y, "itisarg000")
		je.CallStaticVoidMethod(x, y, "itisarg000")
	}

	{
		y := je.GetStaticMethodID(x, "jvexpfn22", "(Ljava/lang/String;I)V")
		log.Println("jvfn22", y)
		gopp.NilPrint(y, "Err")

		cgopp.JNIEnvCallStaticMethod[int](je, x, y, "itisarg000", 123)
	}
	{
		y := je.GetStaticMethodID(x, "jvexpfn3", "()Ljava/lang/String;")
		log.Println("jvfn3", y, x)
		gopp.NilPrint(y, "Err")

		rv := cgopp.JNIEnvCallStaticMethod[string](je, x, y)
		log.Println("rv", rv)

	}
	{
		y := je.GetStaticMethodID(x, "jvexpfn4", "()I")
		log.Println("jvfn4", y)
		gopp.NilPrint(y, "Err")

		rv := cgopp.JNIEnvCallStaticMethod[int](je, x, y)
		log.Println("rv", rv)
		gopp.TruePrint(rv != 444, "Err", "want", 444, "but", rv)

	}

	// cgopp.RunOnJVM(func() error {
	// return nil
	// })
}
