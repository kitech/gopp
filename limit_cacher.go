package gopp

import (
	"log"
	"sync/atomic"
	"time"
)

// ex.
type LimitCacherInt LimitCacher[int]
type LimitCacherStr LimitCacher[string]

type LimitCacher[T any] struct {
	first int32
	ch    chan T
	timeo time.Duration
	tick  *time.Ticker

	// call when full or timeo
	// eventfn func(vals []T, full bool, timeo bool)
	// eventfn func(vals []T)
	eventfnx any

	Maxcnt int

	Stat struct {
		InrunCount int32
		TotalCount int
		FullCount  int
		TimeoCount int
	}
}

// 有 comparable，有没有 callable 关键字

// notify full or timeo
func LimitCacherNew[T any, FT func([]T) | func([]T, bool, bool)](max int, timeo time.Duration, notifyfn FT) *LimitCacher[T] {
	me := &LimitCacher[T]{}
	me.first = 1
	me.ch = make(chan T, max*3)
	me.Maxcnt = max
	me.eventfnx = notifyfn
	me.timeo = timeo
	me.tick = time.NewTicker(timeo)
	return me
}

func (me *LimitCacher[T]) Add(v T) {
	if atomic.CompareAndSwapInt32(&me.first, 1, 0) {
		go me.readtickproc()
	}
	me.ch <- v
	if len(me.ch) >= me.Maxcnt {
		me.notify(true, false)
	}
}
func (me *LimitCacher[T]) notify(full bool, timeo bool) {
	var chlen = len(me.ch)
	var vals []T
	for i := 0; i < me.Maxcnt && i < chlen; i++ {
		v := <-me.ch
		vals = append(vals, v)
	}
	// lftlen := len(me.ch)
	// log.Println("condok", "full", full, "timeo", timeo, "chlen", chlen, "lftlen", lftlen)
	if len(vals) > 0 {
		if me.eventfnx != nil {
			go func() {
				atomic.AddInt32(&me.Stat.InrunCount, 1)
				defer atomic.AddInt32(&me.Stat.InrunCount, -1)
				switch fn := me.eventfnx.(type) {
				case func([]T, bool, bool):
					fn(vals, full, timeo)
				case func([]T):
					fn(vals)
				}

			}()
		} else {
			log.Println("eventfn not set", me.eventfnx == nil)
		}
	}
	if timeo {
		chantrypop1(me.ch)
		//me.tick.Reset(me.timeo)
		go me.readtickproc()
	}
}
func (me *LimitCacher[T]) readtickproc() {
	select {
	case tm, ok := <-me.tick.C:
		// log.Println(tm, ok)
		if tm.IsZero() || !ok {
			break
		}
		me.notify(false, true)
	}
}
func (me *LimitCacher[T]) Flush() {
	me.notify(false, false)
}

func chantrypop1[T any](c <-chan T) (v T) {
	select {
	default:
	case v = <-c:
	}
	return
}
