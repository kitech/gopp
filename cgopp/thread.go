package cgopp

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

/*

#define _GNU_SOURCE
#include <pthread.h>
#include <unistd.h>
#include <sys/syscall.h>
#include <stdint.h>
#include <assert.h>

static uint64_t gettidwp(int rdnum) {
	uint64_t tid = -1;
#if defined(__APPLE__)
	#include <sys/_pthread/_pthread_t.h>
	#include <sys/_pthread/_pthread_types.h>
	//  warning: 'syscall' is deprecated: first deprecated in macOS 10.12 - syscall(2) is unsupported; please switch to a supported interface. For SYS_kdebug_trace use kdebug_signpost().
	// int rv = (rdnum % 2)==1 ?
		// syscall(((int)(SYS_thread_selfid))) : pthread_threadid_np(0, &tid);
	assert(SYS_thread_selfid>0);
	int rv = pthread_threadid_np(0, &tid);
	assert(rv == 0);
#else
	int rv = syscall(((int)(SYS_gettid)));
	tid = rv;
#endif
	return tid;
}

static uintptr_t MyTid() { return (uintptr_t)pthread_self(); }
// macos warning depcreated syscall
// static uintptr_t MyTid2() { return syscall(sizeof(void*)==4?224:186); }
// macos
// static uintptr_t MyTid3() { return kdebug_signpost(SYS_kdebug_trace); }

typedef struct CgoppThreadNoid {
	uint64_t Handle;
	uint32_t No;
	uint32_t Seq;
	uintptr_t Stksize;
	void* Stkaddr;
} CgoppThreadNoid;
static CgoppThreadNoid CgoppThreadNoidInst;
// first call in c, with np==NULL
const CgoppThreadNoid* getCgoppThreadNoid(CgoppThreadNoid* np) {
	CgoppThreadNoid* p0 = &CgoppThreadNoidInst;
	pthread_attr_t attr;
	size_t stksize = 0;
	void* stkaddr = 0;
	pthread_getattr_np(pthread_self(), &attr);
	pthread_attr_getstacksize(&attr, &stksize);
	// pthread_attr_getstackaddr(&attr, &stkaddr);
	pthread_attr_getstack(&attr, &stkaddr, &stksize);

	if (np != 0) {
	np->Handle = (uint64_t)pthread_self();
	np->No = (uint32_t)gettidwp(0);
	np->Seq = np->No;
	np->Stksize = stksize;
	np->Stkaddr = stkaddr;
	return np;
	}else if (p0->Handle == 0){
	p0->Handle = (uint64_t)pthread_self();
	p0->No = (uint32_t)gettidwp(0);
	p0->Seq = p0->No;
	p0->Stksize = stksize;
	p0->Stkaddr = stkaddr;
	}
	return p0;
}
*/
import "C"

// TODO unix/linux/mingw only
// todo 还有一种数字小些的线程号，和pid对应的那种。
//
// Deprecated: Use Gettid instead.
func MyTid() usize {
	return usize(C.MyTid())
}

// final version,. support linux/macos
func Gettid() uint64 { return uint64(C.gettidwp(cint(rand.Int31()))) }

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
// Deprecated: Use Gettid instead.
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
