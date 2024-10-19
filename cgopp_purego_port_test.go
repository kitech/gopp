package gopp

import "testing"

func TestMem1(t *testing.T) {
	if Cmalloc == nil {
		t.Error("Cmalloc should not nil", Cmalloc != nil)
	}

	if v := Mallocgc(0); v == nil {
		t.Error("Mallocgc(0) not nil", v)
	}
	if v := Mallocgc(-1); v != nil {
		t.Error("Mallocgc(-1) not nil", v)
	}
	{
		p := CString("abc")
		if x := GoString(p); x != "abc" {
			t.Error("abc but", x)
		}

	}
	{
		p := CStringgc("abc")
		if x := GoString(p); x != "abc" {
			t.Error(x)
		}
		if n := cstrlen(p); n != 3 {
			t.Error(3, n)
		}
	}
	{
		p, sz := CStringRef("abc")
		if x := GoStringN(p, sz); x != "abc" {
			t.Error(x)
		}
		r := CmemdupAsstr(p, sz)
		if n := cstrlen(r); n != sz {
			t.Error(sz, n)
		}
		r2 := cstrdupgc(r)
		if n := cstrlen(r2); n != sz {
			t.Error(sz, n)
		}

		if x := GoString(r2); x != "abc" {
			t.Error(x)
		}
	}
}
