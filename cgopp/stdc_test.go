package cgopp

import (
	"log"
	"testing"
)

func TestCStringgc(t *testing.T) {
	var s = "foo"
	var cv = CStringgc(s)
	var s2 = GoString(cv)
	// log.Println(cv, s2)
	if s2 != s {
		t.Error("str2c err", s, s2)
	}

}
func TestGostrdup1(t *testing.T) {
	var s = "foo"
	var s2 = Gostrdup(s)
	log.Println(s2, len(s2))
	if s2 != s {
		t.Error("str2c err", s, s2)
	}

}

// go test -v -bench CStringaf1 -run CStringaf1

// 14.04 ns/op
func BenchmarkCStringgc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = "foo"
		var cv = CStringgc(s)
		_ = cv
	}
}

// 118.6 ns/op
func BenchmarkCStringgc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = "foo"
		var cv = CStringgc2(s)
		_ = cv
	}
}

// 442.3 ns/op
func BenchmarkCStringaf1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = "foo"
		var cv = CStringaf(s)
		_ = cv
	}
}
