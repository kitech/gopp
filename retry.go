package gopp

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	RETRY_GET = iota // get a next retry time
	RETRY_FN
	RETRY_FN_WITH_NO
	RETRY_MODE_MAX
)

const (
	// BO_POWER2
	BO_FIBONACCI = iota
	BO_EXPONENTIAL
	BO_NATRURAL
	BO_FIXED
	BO_RANDOM
)

// 1.0
type FixedBackOff struct {
	ntimes int
	name   string
}

func NewFixedBackOff() *FixedBackOff {
	this := &FixedBackOff{}
	this.name = "Fixed"
	return this
}
func (this *FixedBackOff) Reset() {}
func (this *FixedBackOff) Next() (int, time.Duration) {
	this.ntimes++
	return this.ntimes, 100 * time.Millisecond
}
func (this *FixedBackOff) Count() int { return this.ntimes }

// 1.1
type NaturalBackOff struct {
	ntimes int
	name   string
}

func NewNaturalBackOff() *NaturalBackOff {
	this := &NaturalBackOff{}
	this.name = "Natural"
	return this
}
func (this *NaturalBackOff) Reset() {}

func (this *NaturalBackOff) Next() (int, time.Duration) {
	this.ntimes++
	return this.ntimes, time.Duration(this.ntimes*100) * time.Millisecond
}
func (this *NaturalBackOff) Count() int { return this.ntimes }

// 1.5
type ExponentialBackOff struct {
	initialInterval int     // = 100;//初始间隔
	maxInterval     int     // = 5 * 1000L;//最大间隔
	maxElapsedTime  int     // = 50 * 1000L;//最大时间间隔
	multiplier      float32 //= 1.5;//递增倍数（即下次间隔是上次的多少倍）
	ntimes          int
	name            string
}

func NewExponentialBackOff() *ExponentialBackOff {
	this := &ExponentialBackOff{}
	this.name = "Exponential"
	this.Reset()
	return this
}
func (this *ExponentialBackOff) Reset() {
	// this.ntimes = 0
	this.initialInterval = 100
	this.maxInterval = this.initialInterval * 50
	this.maxElapsedTime = this.initialInterval * 500
	this.multiplier = 1.5
}

func (this *ExponentialBackOff) Next() (int, time.Duration) {
	// 首次不要等待
	if this.ntimes == 0 {
		n := this.ntimes
		this.ntimes++
		return n, 0
	}
	v := int(float32(this.initialInterval) * this.multiplier)
	this.initialInterval = v
	n := this.ntimes
	this.ntimes++
	return n, time.Duration(v) * time.Millisecond
}
func (this *ExponentialBackOff) Count() int { return this.ntimes }

// backoff too quick
// 1.9
type FibonacciBackOff struct {
	no1    int
	no2    int
	ntimes int
	name   string
}

func NewFibonacciBackOff() *FibonacciBackOff {
	this := &FibonacciBackOff{}
	this.name = "Fibonacci"
	this.Reset()
	return this
}
func (this *FibonacciBackOff) Reset() {
	this.no1, this.no2 = 0, 100
}

func (this *FibonacciBackOff) Next() (int, time.Duration) {
	no1, no2 := this.no1, this.no2
	this.no1, this.no2 = no2, no1+no2
	this.ntimes += 1
	return this.ntimes, time.Duration(no1+no2) * time.Millisecond // *100 for uniform unit
}
func (this *FibonacciBackOff) Count() int { return this.ntimes }

type RandomBackOff struct {
	mindur time.Duration
	maxdur time.Duration
	ntimes int
}

func NewRandomBackOff(mindur, maxdur time.Duration) *RandomBackOff {
	this := &RandomBackOff{}
	this.Reset()
	this.mindur = mindur
	this.maxdur = maxdur
	return this
}
func (this *RandomBackOff) Reset() {
	// this.mindur = time.Millisecond
	// this.maxdur = 30 * time.Second
}
func (this *RandomBackOff) Next() (int, time.Duration) {
	nxtdur := time.Duration(rand.Int63n(int64(this.maxdur-this.mindur))) + this.mindur
	this.ntimes += 1
	return this.ntimes, nxtdur
}
func (this *RandomBackOff) Count() int { return this.ntimes }

type RetryBackOff interface {
	Next() (int, time.Duration)
	Reset()
	Count() int
}

func backoffByType(botype int) (bkoff RetryBackOff) {
	switch botype {
	case BO_EXPONENTIAL:
		bkoff = NewExponentialBackOff()
	case BO_FIBONACCI:
		bkoff = NewFibonacciBackOff()
	case BO_NATRURAL:
		bkoff = NewNaturalBackOff()
	case BO_FIXED:
		fallthrough
	default:
		bkoff = NewFibonacciBackOff()
	}
	return
}

type Retryer struct {
	ratio  int //
	mode   int
	boff   RetryBackOff
	dofnno func(int) error
	dofn   func() error
}

func (this *Retryer) setBackOff(boty int) {
	this.boff = backoffByType(boty)
}

// 指数递增模式，默认需要自己Sleep()
// 默认初始100ms，等待间隔按照1.5倍递增，
// 即第一次100ms，第二次150ms，第三次225ms
// todo 需要添加个倍率参数
// usage:
//
//	rter := NewRetry() || Retry(1)
//	time.Sleep(rter.NextWaitOnly())
func NewRetry(ratio ...int) *Retryer {
	this := &Retryer{}
	this.ratio = Firstof(ratio).Int()
	this.ratio = IfElse2(this.ratio == 0, 1, this.ratio)
	this.mode = RETRY_GET
	this.setBackOff(BO_EXPONENTIAL)
	return this
}

func (this *Retryer) NextWait() (ntimes int, nwait time.Duration) {
	ntimes, nwait = this.boff.Next()
	// log.Println(ntimes, nwait)
	nwait *= time.Duration(this.ratio)
	// log.Println(ntimes, nwait)
	return
}

func (this *Retryer) NextWaitOnly() time.Duration {
	_, nwait := this.NextWait()
	return nwait
}

// 指数递增模式，指定回调函数
func NewRetryFn(f func(ntimes int) error, ratio ...int) *Retryer {
	this := NewRetry(ratio...)
	this.mode = RETRY_FN_WITH_NO
	this.dofnno = f
	return this
}
func NewRetryFnOnly(f func() error, ratio ...int) *Retryer {
	this := NewRetry(ratio...)
	this.mode = RETRY_FN
	this.dofn = f
	return this
}

// 指数递增模式，指定回调函数
// 回调函数返回error!=nil表示重试，nil表示结束重试
// 对于有返回值的情况，需要自己处理
// usage:
//
//	rter := NewRetryFng(fn)
//	rter.Do(...)
func NewRetryFng[FT func(ntimes int) error | func() error](fx FT, ratio ...int) *Retryer {

	var ff = fx // works
	_ = ff

	switch f := any(fx).(type) {
	case func(int) error:
		return NewRetryFn(f, ratio...)
	case func() error:
		return NewRetryFnOnly(f, ratio...)
	default:
		Panicp2(fx)
	}

	// return this
	return nil
}

// should block
// unit base: 1*time.Millisecond, so if want 2s, unit=time.Duration(2000)
func (this *Retryer) Do(ntimes ...int) error {
	return this.do(this.mode == RETRY_FN_WITH_NO, ntimes...)
}

func (this *Retryer) do(withno bool, ntimes ...int) (err error) {
	innern := 0
	for {
		if withno {
			err = this.dofnno(innern)
		} else {
			err = this.dofn()
		}
		if err == nil {
			break
		} else {
			if len(ntimes) > 0 && innern+1 >= ntimes[0] {
				err = fmt.Errorf("Exceed max ntimes: %d, %v", ntimes, err)
				break
			} else {
				n, v := this.NextWait()
				innern = n
				// waitdur := time.Duration(this.ratio) * v
				waitdur := v
				// log.Println(innern, waitdur, ntimes, this.ratio)
				if waitdur != 0 {
					time.Sleep(waitdur) // /100 for uniform unit
				}
			}
		}
	}
	return
}

func DoTimes(n int, f func(n int)) {
	for i := 0; i < n; i++ {
		f(i)
	}
}
func DoTimesOnly(n int, f func()) {
	for i := 0; i < n; i++ {
		f()
	}
}
func DoTimesg[FT func(int) | func()](n int, fg FT) {
	switch ff := any(fg).(type) {
	case func(int):
		DoTimes(n, ff)
	case func():
		DoTimesOnly(n, ff)
	default:
		Panicp2(fg)
	}
}

// / only calc next time, not do run
// avoid retry times
type Retryer2 struct {
	minwait  time.Duration
	maxwait  time.Duration
	stepwait time.Duration
	steptype int
	cycle    bool
	reverse  bool // TODO
	bkoff    RetryBackOff
}

// ...
func NewRetryer2(minwait, maxwait, stepwait time.Duration, steptype int, cycle bool) *Retryer2 {
	r2 := &Retryer2{}
	r2.minwait = minwait
	r2.maxwait = maxwait
	r2.stepwait = stepwait
	r2.steptype = steptype
	r2.cycle = cycle

	if steptype == BO_RANDOM {
		r2.bkoff = NewRandomBackOff(minwait, maxwait)
	} else {
		r2.bkoff = backoffByType(steptype)
	}

	return r2
}

func (this *Retryer2) Next() time.Duration {
	if this.steptype == BO_RANDOM {
		_, dur := this.bkoff.Next()
		return dur
	}

	cnter, _ := this.bkoff.Next()
	absdur := time.Duration(cnter-1)*this.stepwait + this.minwait
	if this.cycle && absdur > this.maxwait {
		this.bkoff.Reset()
	}
	return absdur
}

func (this *Retryer2) SleepNext()    { time.Sleep(this.Next()) }
func (this *Retryer2) TryCount() int { return this.bkoff.Count() }

func Forever() {
	for {
		time.Sleep(1 * time.Second)
	}
}
