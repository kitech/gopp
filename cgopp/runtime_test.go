package cgopp

import (
	"log"
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
