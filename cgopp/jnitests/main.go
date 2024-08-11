package main

import (
	"log"
	"os"
	"testing"

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

//export Java_MainKt_goexpfn1
func Java_MainKt_goexpfn1() {
	log.Println("heree", C.int(0))
}

func main() {

}

func init() {
	log.Println("heree")
	// 没有正确安化testing不会执行的
	os.Args = append(os.Args, "-test.v=true")
	testing.Init()
	t := &testing.T{} // todo
	t.Error("hahhaha")

	// todo test cgopp.RunOnJvm
	cgopp.JNI_OnLoad_Callback = func() {
		log.Println("herere", cgopp.Jenv)
		je := cgopp.Jenv
		// cls4c := cgopp.CStringgc("LMain")
		ver := je.GetVersion()
		log.Println("jvm ver", ver, cgopp.JVMTid(), cgopp.MyTid())

		// tfn1 := func(t *testing.T) {
		// }
		// t.Run("tfn1", tfn1) // crash here

		x := je.FindClass("Main")
		log.Println("jvcls Main", x)
		gopp.NilPrint(x, "Err")
		if x == nil {
			t.Error(x)
		}

		{
			y := je.GetStaticMethodID(x, "jvexpfn1", "()V")
			log.Println("jvfn1", y)
			gopp.NilPrint(y, "Err")
			if y == nil {
				t.Error()
			}
			je.CallStaticVoidMethod(x, y)
		}
		{
			y := je.GetStaticMethodID(x, "jvexpfn2", "(Ljava/lang/String;)V")
			log.Println("jvfn2", y)
			gopp.NilPrint(y, "Err")
			if y == nil {
				t.Error()
			}

			cgopp.JNIEnvCallStaticMethod[int](je, x, y, "itisarg000")

		}
		{
			y := je.GetStaticMethodID(x, "jvexpfn22", "(Ljava/lang/String;I)V")
			log.Println("jvfn22", y)
			gopp.NilPrint(y, "Err")
			if y == nil {
				t.Error()
			}

			cgopp.JNIEnvCallStaticMethod[int](je, x, y, "itisarg000", 123)
		}
		{
			y := je.GetStaticMethodID(x, "jvexpfn3", "()Ljava/lang/String;")
			log.Println("jvfn3", y)
			gopp.NilPrint(y, "Err")
			if y == nil {
				t.Error()
			}

			rv := cgopp.JNIEnvCallStaticMethod[string](je, x, y)
			log.Println("rv", rv)
			if rv != "jvexpfn3retval" {
				t.Error()
			}
		}
		{
			y := je.GetStaticMethodID(x, "jvexpfn4", "()I")
			log.Println("jvfn4", y)
			gopp.NilPrint(y, "Err")
			if y == nil {
				t.Error()
			}

			rv := cgopp.JNIEnvCallStaticMethod[int](je, x, y)
			log.Println("rv", rv)
			gopp.TruePrint(rv != 444, "Err", "want", 444, "but", rv)
			if rv != 444 {
				t.Error("want", 444, "but", rv)
			}
		}

		// cgopp.RunOnJVM(func() error {
		// return nil
		// })
	}
}
