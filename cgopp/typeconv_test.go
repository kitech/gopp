package cgopp

import (
	"log"
	"testing"
)

func strdup(s string) string {
	r := s + "/"
	return r
}
func TestStrAltChar(t *testing.T) {
	var s0 = "abcdefg12345"
	{
		var s = strdup(s0)
		StrAltChar(&s, 0, 'x')
		log.Println(s)
	}
	{
		var s = strdup(s0)
		StrAltChar(&s, 1, 'x')
		log.Println(s)
	}
	{
		var s = strdup(s0)
		StrNilTail(&s)
		log.Println(s)
	}
}
