package cgopp

import (
	"log"
	"testing"
)

func TestDllist0(t *testing.T) {
	c := DyldImageCount()
	log.Println(c)
	for i := 0; i < c; i++ {
		log.Println(i, DyldImageName(i))
		DyldListSymbols(i)
		if i > 2 {
			break
		}
	}
}

func TestDllist1(t *testing.T) {
	vals := DyldImagesInc()
	log.Println(len(vals), vals, len(vals))
}
