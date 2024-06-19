package gopp

import (
	"log"
	"testing"
)

func TestLastof1(t *testing.T) {
	{
		s := "12345abcde"
		v := LastofGs(s)
		log.Println(v, rune(v), string(rune(v)))
	}
	{
		c := []string{"1", "2", "3", "a"}
		// v := LastofG(c)
		v := LastofGv(c)
		log.Println(v)
	}

}
