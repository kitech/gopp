package cgopp

/*
 */
import "C"

func MallocTrim() int {
	return int(usize(C.malloc(0)))
}
