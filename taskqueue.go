package gopp

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/crc64"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// 使用最多一个runner goroutine
// runner不会常驻，只在需要的时候启动，执行完成退出
// 在没有任务的时候退出,有任务的时候启动,但不会同时启动多个
// var tq = NewTaskqez()

func NewTaskqez() *Taskqez {
	me := &Taskqez{}
	me.qc = make(chan *taskinfoez, 286)

	return me
}

type Taskqez struct {
	qc      chan *taskinfoez
	running int32
	runcnt  int32
	// curti   *taskinfoez
}

type taskinfoez struct {
	no     int32
	fn     func()
	name   string
	btime  time.Time
	rbtime time.Time
	etime  time.Time
}

func (me *Taskqez) Add(t func(), name ...string) {
	NilPanic(t)
	ti := &taskinfoez{}
	ti.fn = t
	ti.name = FirstofGv(name)
	ti.btime = time.Now()

	if len(me.qc) == cap(me.qc) {
		Warn("Queue full maybe block!!!", cap(me.qc))
	}
	me.qc <- ti
	if atomic.CompareAndSwapInt32(&me.running, 0, 1) {
		go me.taskproc()
	} else {
		// log.Println("Taskqez", "running???", me.running, len(me.qc))
	}
}

// 需要时创建，用完退出
func (me *Taskqez) taskproc() {
	btime := time.Now()

	qc := me.qc
	oldcnt := len(qc)
	var names []string
	var nos []int32
	for len(qc) > 0 {
		select {
		case ti, ok := <-qc:
			if !ok {
				goto endfor
			}
			names = append(names, ti.name)
			nos = append(nos, ti.no)
			ti.rbtime = time.Now()
			log.Println("task running", ti.name, any(ti.fn), "waitdur", ti.rbtime.Sub(ti.btime))
			ti.fn()
			ti.etime = time.Now()
			atomic.AddInt32(&me.runcnt, 1)
			log.Println("task done", ti.name, any(ti.fn), ti.rbtime.Sub(ti.btime))
			// goto endfor
		}
	}
endfor:
	log.Println("Taskproc done", me.runcnt, oldcnt, "=>", len(qc), names, "procalive", time.Since(btime))
	if atomic.CompareAndSwapInt32(&me.running, 1, 0) {
	} else {
		log.Println("cannot reset me.running to 0", me.running)
	}
}

// todo with max limit uncached task queue
// todo with max limit cached proc task queue

// /// Logdeduper
type Logdeduper struct {
	ts     time.Time
	crcval uint64
	ddcnt  int64
}

func NewLogdeduper() *Logdeduper {
	me := &Logdeduper{}
	return me
}

// 如果与最后一条重复，则返回true
func (me *Logdeduper) Check(args ...any) (string, bool) {
	logstr := PackArgs(args...)
	return me.Check2(logstr)
}
func (me *Logdeduper) Check2(logstr string) (string, bool) {
	crc := Crc64Str(logstr)
	if atomic.LoadUint64(&me.crcval) == crc {
		atomic.AddInt64(&me.ddcnt, 1)
		rv := fmt.Sprintf("%s (%d in %v)", logstr, me.ddcnt, Dur2hum(time.Since(me.ts)))
		return rv, true
	} else {
		me.ts = time.Now()
		atomic.StoreUint64(&me.crcval, crc)
		atomic.StoreInt64(&me.ddcnt, 0)
		return logstr, false
	}
}

// /// Deduper
type Deduper struct {
	bits     int
	hser     hash.Hash
	onlylast bool
	saveval  bool
	vals     sync.Map
}

type nophash struct {
	d []byte
}

var _ hash.Hash = (*nophash)(nil)

func (me *nophash) Reset() {
	if me.d != nil {
		me.d = me.d[:0]
	}
}
func (me *nophash) Size() int      { return len(me.d) }
func (me *nophash) BlockSize() int { return 128000 }
func (me *nophash) Sum(b []byte) []byte {
	if b == nil {
		return me.d
	}
	return append(me.d, b...)
}
func (me *nophash) Write(b []byte) (int, error) {
	if b == nil {
		return 0, nil
	}
	me.d = append(me.d, b...)
	return len(b), nil
}

// bits 0, 40(sha1), 64(crc64), 128(md5), 256(sha256), 512(sha512)
func NewDeduper(bits int) *Deduper {
	me := &Deduper{}
	me.bits = bits

	switch bits {
	case 512:
		me.hser = sha512.New()
	case 256:
		me.hser = sha256.New()
	case 128:
		me.hser = md5.New()
	case 64:
		me.hser = crc64.New(crc64htiso)
	default:
		me.hser = &nophash{}
	}

	return me
}

func (me *Deduper) Isdup(vx any) bool {
	var hser = me.hser
	defer hser.Reset()

	switch v := vx.(type) {
	case string:
		hser.Write([]byte(v))
	case []byte:
		hser.Write(v)
	default:
		hser.Write([]byte(ToStr(vx)))
	}
	hsval := string(hser.Sum(nil))

	_, ok := me.vals.Load(hsval)
	if !ok {
		me.vals.Store(hsval, true)
	}
	return ok
}

// //// CallLater. CallMerger
type CallMerger struct {
	// q1 cmap.ConcurrentMap[string, []*cmrunnerinfo]
	q2 map[string][]*cmrunnerinfo
	mu sync.RWMutex
}
type cmrunnerinfo struct {
	fn    func()
	key   string
	delay time.Duration
	ts    time.Time
	async bool
}

func NewCallMerger() *CallMerger {
	me := &CallMerger{}
	// me.q1 = cmap.New[[]*cmrunnerinfo]()
	me.q2 = map[string][]*cmrunnerinfo{}
	return me
}

func (me *CallMerger) AddAsync(key string, delay time.Duration, f func()) {
	me.implAdd(true, key, delay, f)
}
func (me *CallMerger) Add(key string, delay time.Duration, f func()) {
	me.implAdd(false, key, delay, f)
}
func (me *CallMerger) implAdd(async bool, key string, delay time.Duration, f func()) {

	me.mu.Lock()
	defer me.mu.Unlock()

	ri := &cmrunnerinfo{}
	ri.key = key
	ri.delay = delay
	ri.fn = f
	ri.async = async
	ri.ts = time.Now().Add(delay)

	// tk := time.NewTicker(maxdelay)
	// tmer := time.NewTimer(maxdelay)

	var cnt int
	_, ok := me.q2[key]
	if !ok {
		me.q2[key] = Sliceof(ri)
	} else {
		me.q2[key] = append(me.q2[key], ri)
	}
	cnt = len(me.q2[key]) - 1

	// todo cancel old timer when new coming
	time.AfterFunc(delay, func() { me.runit(key, cnt) })
}

// 如果在delay时间没有新加入的，则执行
func (me *CallMerger) runit(key string, no int) {
	var okri *cmrunnerinfo
	var totcnt int
	me.mu.Lock()
	olds, ok := me.q2[key]
	if ok {
		mm := NewMinmaxer[int64]()
		for _, ri := range olds {
			mm.Input(ri.ts.UnixNano())
		}
		nowt := time.Now()
		if nowt.UnixNano() > mm.Max {
			delete(me.q2, key)
			totcnt = len(olds)
			okri = olds[no]
			// log.Println("canrun", key, no, totcnt)
		} else {
			// log.Println("notrun", key, no, totcnt, "left", time.Duration(mm.Max-nowt.UnixNano()))
		}
	} else {
		// maybe already runned
	}

	me.mu.Unlock()

	if okri != nil {
		log.Println("runmrg ...", key, no, totcnt, "=>", 1)
		if okri.async {
			go okri.fn()
		} else {
			okri.fn()
		}
	}
}
