package gopp

import (
	"log"
	"testing"
)

func TestHC0(t *testing.T) {
	hc := NewHttpClient()
	hc.Get("http://")
	log.Println(hc, hc.Req.Method, hc.Req.URL)
	resp, err := hc.Do()
	ErrPrint(err)
	log.Println(resp, hc.Getopts())
}
