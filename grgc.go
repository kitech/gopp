package gopp

// 也许可以做一个goroutine回收监测的库
// 记录启动标识，位置，时间，parent，便于调试
// 在任意位置输出还存活的goroutine
// 这需要依赖一个custom的goroutine，类似pthread_create创建

import (
	// "log"
	"math"
	"reflect"
	"sync/atomic"
	"time"
)

type Prthread struct {
	tid     uint64
	tfn     func()
	ctime   time.Time
	mtime   time.Time
	etime   time.Time
	retime  time.Time // real end time
	stopped int32
}

var prthread_id uint64

func NewThread(tfn func()) *Prthread {

	var tid uint64

	atomic.CompareAndSwapUint64(&prthread_id, math.MaxUint64, 0)
	tid = atomic.AddUint64(&prthread_id, 1)

	return &Prthread{tid: tid, ctime: time.Now(), tfn: tfn}
}

func (this *Prthread) Start() bool {
	this.mtime = time.Now()
	if this.tfn != nil {
		go func() {
			donech := make(chan struct{}, 0)
			go func() {
				this.tfn()
				atomic.CompareAndSwapInt32(&this.stopped, 0, 1)
				donech <- struct{}{}
			}()

			<-donech
			this.retime = time.Now()
		}()

		return true
	}

	return false
}

func (this *Prthread) Terminate() {
	if atomic.CompareAndSwapInt32(&this.stopped, 0, 1) {
		// ok
		this.etime = time.Now()
	} else {
		// already stopped
	}
}

func (this *Prthread) IsStopped() bool {
	return atomic.LoadInt32(&this.stopped) == 1
}

type grpool struct {
}

// /
var fakerefcnt = 0

func FakeRef(x any) int { return FakeSymref(x) }

// we want some func in binary even not used when build
// avoid go link omit unused func he think
func FakeSymref(x any) int {
	if Absfalse() {
		k := reflect.TypeOf(x).Kind()
		return int(k)
	}
	fakerefcnt++
	return fakerefcnt
}

func Reftimed(v any, dur time.Duration) {
	time.AfterFunc(dur, func() {
		if v != nil {
		}
	})
}
func ReftimedSec(v any, sec int) {
	Reftimed(v, time.Duration(sec)*time.Second)
}

type ThreadNoid struct {
	Handle  uint64 // pthread_self
	No      uint32 // SYS_gettid
	Seq     uint32
	Stksize usize
	Stkaddr voidptr
}
