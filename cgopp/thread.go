package cgopp

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

/*

#include <unistd.h>
#include <sys/syscall.h>
#include <stdint.h>
#include <pthread.h>

static uintptr_t MyTid() { return (uintptr_t)pthread_self(); }
// macos warning depcreated syscall
// static uintptr_t MyTid2() { return syscall(sizeof(void*)==4?224:186); }
// macos
// static uintptr_t MyTid3() { return kdebug_signpost(SYS_kdebug_trace); }
*/
import "C"

// TODO unix/linux only
func MyTid() usize {
	return usize(C.MyTid())
}

// func MyTid2() uint64 {
// 	return uint64(C.MyTid2())
// }

const PtrSize = 32 << uintptr(^uintptr(0)>>63)
const IntSize = strconv.IntSize
const CIntSize = C.sizeof_int

var archs = map[int]uintptr{
	32: 224, 64: 186,
}

// todo macos not work
func MyTid3() usize {
	r1, r2, err := syscall.Syscall(archs[PtrSize], 0, 0, 0)
	if err != 0 && runtime.GOOS == "darwin" {
		panic("notimpl " + runtime.GOOS + " " + err.Error())
	}
	if false {
		log.Println(r1, r2, err)
	}
	return usize(r1)
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
