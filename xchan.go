package gopp

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

// 通用的可自动回收关闭的go channel封装
// 这个类型的封装不行啊，根本行不通，限制太多了。
var gwg sync.WaitGroup

type xchan struct {
	uchan  chan interface{}
	wg     sync.WaitGroup
	pdchan chan struct{}
	cdchan chan struct{}
	// 使用标识的话，就怕在不同goroutine使用时有问题。
}

func NewXChan(c int) *xchan {
	uchan := make(chan interface{}, c)
	pdchan := make(chan struct{}, 1)
	cdchan := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	gwg.Add(1)
	return &xchan{uchan, wg, pdchan, cdchan}
}

// Producer done
func (this *xchan) PDone() {
	if len(this.pdchan) > 0 {
		return
	}
	this.pdchan <- struct{}{}

	go func() {
		// 写入数据完成，读取数据完成
		this.wg.Wait()
		close(this.pdchan)
		close(this.cdchan)
		close(this.uchan)
		gwg.Done()
	}()
}

// Consumer done
func (this *xchan) CDone() {
	if len(this.cdchan) > 0 {
		return
	}
	this.cdchan <- struct{}{}

	this.wg.Done()
	// this.wg.Done()  // 怎么就不能检测一下呢
}

func (this *xchan) Write(v interface{}) {
	this.uchan <- v
}

func (this *xchan) Read() interface{} {
	v := <-this.uchan
	return v
}

func xchan_wait() {
	gwg.Wait()
}

//////

// 安全地寫入channel，避免panic，避免阻塞
func SafeTrySend(c interface{}, v interface{}) (err error) {
	cv := reflect.ValueOf(c)
	vv := reflect.ValueOf(v)
	defer func() {
		if x := recover(); x != nil {
			// runtime.plainError
			err = errors.New(fmt.Sprintf("%v", x))

		}
	}()
	ok := cv.TrySend(vv)
	if !ok {
		err = errors.New("will block")
	}
	return
}

// 安全地寫入channel，避免panic，避免阻塞x
func ChanTrySend(c any, v any) error {
	return SafeTrySend(c, v)
}
func ChanSendTry[T any](c chan T, v T) bool {
	select {
	default:
		return false
	case c <- v:
		return true
	}
}

func _chanTrySend1[T any](c chan T, v T) error {
	var c1 = make(chan int)
	return _chanTrySend1(c1, 8)
}

// 帶超時的寫入channel
func ChanSendTimeo[T any](c chan T, v T, d time.Duration) (err error) {
	select {
	case c <- v:
	case <-time.After(d):
		err = ErrTimeout
	}

	return
}

// 帶超時的寫入channel
func ChanSendCtx[T any](c chan T, v T, ctx context.Context) (err error) {
	select {
	case c <- v:
	case <-ctx.Done():
		err = CanceledError
	}
	return
}

func ChanRecvTry[T any](c chan T) (v T, ok bool) {
	select {
	default:
	case v, ok = <-c:
	}
	return
}

// 帶超時的寫入channel
func ChanRecvCtx[T any](c chan T, ctx context.Context) (v T, err error) {
	select {
	case v = <-c:
	case <-ctx.Done():
		err = CanceledError
	}

	return
}

// 帶超時的寫入channel
func ChanRecvTimeo[T any](c chan T, d time.Duration) (v T, err error) {
	select {
	case v = <-c:
	case <-time.After(d):
		err = ErrTimeout
	}
	return
}

// 同时select一组不同类型的channel
func Select(chs []interface{}, timeout time.Duration) interface{} {
	sltCases := []reflect.SelectCase{}
	if timeout.Nanoseconds() != 0 {
		ch := time.After(timeout)
		sltCase := reflect.SelectCase{}
		sltCase.Chan = reflect.ValueOf(ch)
		sltCases = append(sltCases, sltCase)
	}

	for _, chx := range chs {
		sltCase := reflect.SelectCase{}
		sltCase.Chan = reflect.ValueOf(chx)
		if sltCase.Chan.Type().Kind() != reflect.Chan {
			return nil
		}
		sltCase.Dir = reflect.SelectRecv
		sltCases = append(sltCases, sltCase)
	}

	chosen, rval, ok := reflect.Select(sltCases)
	if !ok {
		_ = chosen
		return nil
	}
	return rval.Interface()
}
