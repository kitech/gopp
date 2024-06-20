// error with stack
package gopp

import (
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"

	_ "unsafe"

	_ "github.com/pkg/errors"
)

var (
	// 这可以导出包里的私有变量,函数

	//go:linkname ErrTimeout net.errTimeout
	ErrTimeout error
	//go:linkname CanceledError net.canceledError
	CanceledError error

	//go:linkname ErrClosed os.ErrClosed
	ErrClosed error

	// is type
	// go:linkname PlainError runtime.plainError
	// PlainError error
)

func ErrAny(errs ...error) bool {
	for _, err := range errs {
		if err != nil {
			return true
		}
	}
	return false
}
func ErrAll(errs ...error) bool {
	for _, err := range errs {
		if err == nil {
			return false
		}
	}
	return true
}

// Error with errno and stack info
type Error struct {
	errno  int
	errstr string
	stack  []uintptr
}

// eno 模擬可選參數
func NewError(estr string, eno ...int) Error {
	return newErrorN(estr, 2, eno...)
}
func NewErrora(ev interface{}, eno ...int) Error {
	return newErrorN(fmt.Sprintf("%v", ev), 2, eno...)
}
func NewErroraN(ev interface{}, skipn int, eno ...int) Error {
	return newErrorN(fmt.Sprintf("%v", ev), skipn+2, eno...)
}

func newErrorN(estr string, skipn int, eno ...int) Error {
	var pc = make([]uintptr, 0)
	n := runtime.Callers(skipn, nil)
	pc = make([]uintptr, n)
	runtime.Callers(skipn, pc)
	if eno != nil && len(eno) > 0 {
		return Error{errno: eno[0], errstr: estr, stack: pc}
	}
	return Error{errno: 0, errstr: estr, stack: pc}
}

func ErrorFrom(e error) Error {
	return NewError(e.Error())
}

func (this Error) Errno() int {
	return this.errno
}

func (this Error) Errstr() string {
	return this.errstr
}

func (this Error) Error() string {
	return this.String()
}

func (this Error) String() string {
	return fmt.Sprintf("Error: %d, %s", this.errno, this.errstr)
}

func (this Error) PrintStack() {
	fmt.Println(this.String())
	fmt.Println()
	for idx, pc := range this.stack {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		fmt.Printf("#%d, %s %s:%d\n", idx, fn.Name(), file, line)
	}
}

func (this Error) Show() {
	this.PrintStack()
}

func (this Error) Display() {
	this.PrintStack()
}

// /////// some conditional print utils
var justprintinfd = int(os.Stdin.Fd())
var justprinterrfd = int(os.Stderr.Fd()) // stderr/stdout
var justprintoutfd = int(os.Stdout.Fd())

func init() {
	Assert(justprintoutfd == 1, "why os.Stdout not 1", justprintoutfd)
	Assert(justprinterrfd == 2, "why os.Stderr not 2", justprinterrfd)
}

var justprintoutfn = func(s string) {}
var justprintoutfnasync bool = false

func SetLogPrintFunc(async bool, fn func(string)) (oldasync bool, oldfn func(string)) {
	oldasync = justprintoutfnasync
	oldfn = justprintoutfn

	if oldfn != nil {
		// wttodo
	}
	justprintoutfn = fn
	justprintoutfnasync = async
	return
}
func UnsetLogPrintFunc() (oldasync bool, oldfn func(string)) {
	oldasync = justprintoutfnasync
	oldfn = justprintoutfn

	justprintoutfn = nil
	justprintoutfnasync = false
	return
}

// =PackArg?
func printq(v any, args ...any) string {
	ci := GetCallerInfo(3)
	msg := fmt.Sprintf("%s %+v", ci, v)
	for _, argx := range args {
		switch arg := argx.(type) {
		case string:
			if len(arg) == 0 {
				argx = "\"\""
			}
		}
		msg += fmt.Sprintf(" %+v", argx)
	}
	return msg
}

// 开头类型为error的，可以多个，之后任意类型
// 如果没有error类型的开头，则不打印直接返回
func ErrsPrint(errandargs ...any) {
	var errs []error
	var args []any
	for i, vx := range errandargs {
		switch v := vx.(type) {
		case error:
			if v != nil {
				errs = append(errs, v)
			}
		default:
			args = errandargs[i:]
			goto endfor
		}
	}
endfor:
	if len(errs) > 0 {
		err := fmt.Errorf("%v", errs)
		s := printq(err, args...)
		log.Output(2, s)
		if justprintoutfn != nil {
			if justprintoutfnasync {
				go justprintoutfn(s)
			} else {
				justprintoutfn(s)
			}

		}
	}
}

// ErrPrint 用的最多，其次是 NilPrint, ZeroPrint, TruePrint, FalsePrint
func ErrPrint(err error, args ...any) error {
	if err != nil {
		s := printq(err, args...)
		log.Output(2, s)
		if justprintoutfn != nil {
			if justprintoutfnasync {
				go justprintoutfn(s)
			} else {
				justprintoutfn(s)
			}

		}
	}
	return err
}
func ErrPrintExcept(err error, except error, args ...any) error {
	if err == except {
		return err
	}
	if err != nil {
		log.Output(2, printq(err, args...))
	}
	return err
}
func ErrPrintExcept2(err error, except string, args ...any) error {
	if ErrHave(err, except) {
		return err
	}
	if err != nil {
		log.Output(2, printq(err, args...))
	}
	return err
}

func ErrFatal(err error, args ...any) {
	if err != nil {
		log.Output(2, printq(err, args...))
		os.Exit(-1)
	}
}

func ErrPanic(err error, args ...any) {
	if err != nil {
		log.Output(2, printq(err, args...))
		log.Panicln(err)
	}
}

func ErrHave(err error, s string) bool {
	return err != nil && strings.Contains(err.Error(), s)
}
func ErrEqual(err error, s string) bool {
	return err != nil && err.Error() == s
}
func ErrPrefix(err error, s string) bool {
	return err != nil && strings.HasPrefix(err.Error(), s)
}
func ErrSuffix(err error, s string) bool {
	return err != nil && strings.HasSuffix(err.Error(), s)
}
func ErrBegin(err error, s string) bool {
	return ErrPrefix(err, s)
}
func ErrEnd(err error, s string) bool {
	return ErrSuffix(err, s)
}

// safe
func Errtostr(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func ErrHuman(err error) string {
	if err == nil {
		return "OK"
	}
	return err.Error()
}
func ErrHumanShort(err error) string {
	if err == nil {
		return "OK"
	}
	return "Failed"
}

func FalsePrint(ok bool, args ...any) bool {
	if !ok {
		s := printq("CondFalse", args...)
		log.Output(2, s)
		if justprintoutfn != nil {
			if justprintoutfnasync {
				go justprintoutfn(s)
			} else {
				justprintoutfn(s)
			}

		}
	}
	return ok
}

func TruePrint(ok bool, args ...any) bool {
	if ok {
		s := printq("CondTrue", args...)
		log.Output(2, s)
		if justprintoutfn != nil {
			if justprintoutfnasync {
				go justprintoutfn(s)
			} else {
				justprintoutfn(s)
			}

		}
	}
	return ok
}

func LevelPrint(lvl string, args ...any) {
	s := printq(lvl, args...)
	log.Output(2, s)
	if justprintoutfn != nil {
		if justprintoutfnasync {
			go justprintoutfn(s)
		} else {
			justprintoutfn(s)
		}

	}
	return
}
func Infop(args ...any) {
	s := printq("Info", args...)
	log.Output(2, s)
	if justprintoutfn != nil {
		if justprintoutfnasync {
			go justprintoutfn(s)
		} else {
			justprintoutfn(s)
		}

	}
	return
}
func Warnp(args ...any) {
	s := printq("Warn", args...)
	log.Output(2, s)
	if justprintoutfn != nil {
		if justprintoutfnasync {
			go justprintoutfn(s)
		} else {
			justprintoutfn(s)
		}

	}
	return
}
func Debugp(args ...any) {
	s := printq("Debug", args...)
	log.Output(2, s)
	if justprintoutfn != nil {
		if justprintoutfnasync {
			go justprintoutfn(s)
		} else {
			justprintoutfn(s)
		}

	}
	return
}

// BUG: panic: reflect: call of reflect.Value.IsNil on uint64 Value
func NilPrint(v interface{}, args ...any) any {
	if v == nil {
		s := printq("CondNil", args...)
		log.Output(2, s)
		if justprintoutfn != nil {
			if justprintoutfnasync {
				go justprintoutfn(s)
			} else {
				justprintoutfn(s)
			}
		}
	}
	return v
}

func NilFatal(v any, args ...any) {
	if v == nil {
		log.Fatalln(printq("CondNil", args...))
	}
}

// supported: number,string,pointer
func ZeroPrint(v any, args ...any) any {
	if reflect.Zero(reflect.TypeOf(v)).Interface() == v {
		s := printq("CondZero", args...)
		log.Output(2, s)
		if justprintoutfn != nil {
			if justprintoutfnasync {
				go justprintoutfn(s)
			} else {
				justprintoutfn(s)
			}
		}
	}
	return v
}

// NOT mean: error != nil or bool == false or int == 0 or pointer == nil or string == "" other what?
func NotPrint(v any, args ...any) any {
	switch rv := v.(type) {
	case error:
		ErrPrint(rv, args...)
	case bool:
		FalsePrint(rv, args...)
	case string:
		if len(rv) == 0 {
			ZeroPrint(rv, args...)
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		ZeroPrint(rv, args...)
	default:
		vty := reflect.TypeOf(v)
		switch vty.Kind() {
		case reflect.Ptr, reflect.Func, reflect.Interface,
			reflect.Array, reflect.Chan, reflect.Map, reflect.Slice,
			reflect.UnsafePointer:
			NilPrint(v, args...)
		default:
			log.Println("unknown type:", vty.String())
		}
	}
	return v
}

// seperate by commba
func CommaPrintln(args ...any) {
	nargs := []interface{}{}
	for _, arg := range args {
		nargs = append(nargs, arg, ", ")
	}
	log.Output(2, printq("", nargs...))
}

func init() {
	if false {
		f1 := func() error {
			return NewError("hehe")
		}
		if f1 != nil {
		}
	}
}

// usage: defer Panicp()
func Panicp() {
	if err := recover(); err != nil {
		bs := debug.Stack()
		log.Println("error:", err, ", stack:", string(bs))
	}
}

// usage: defer func(){Panicp(recover())}()
// do use like this: defer gopp.Panicp(recover())
// because need lazy call recover()
// need macro in golang
func Panicp2(err interface{}) {
	if err != nil {
		bs := debug.Stack()
		log.Println("error:", err, ", stack:", string(bs))
	}
}

// ///
func Trace(args ...any) {
	s := printq("[TRACE] ", args...)
	log.Output(2, s)
}
func Debug(args ...any) {
	s := printq("[DEBUG] ", args...)
	log.Output(2, s)
}
func Info(args ...any) {
	s := printq("[INFO] ", args...)
	log.Output(2, s)
}
func Warn(args ...any) {
	s := printq("[WARN] ", args...)
	log.Output(2, s)
}
func Fatal(args ...any) {
	s := printq("[FATAL] ", args...)
	log.Output(2, s)
}

// 一般 GetCallerInfo(2)
// return ` funcName, fileName, lineNo`
// funcName like pkg1/barfun
// fileName only basename
func GetCallerInfo(skip int) (info string) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	if strings.Count(funcName, "/") > 1 {
		pos := strings.LastIndex(funcName, "/")
		pos = strings.LastIndex(funcName[:pos], "/")
		funcName = funcName[pos+1:]
	}
	fileName := path.Base(file) // The Base function returns the last element of the path

	if !strings.HasSuffix(funcName, ")") {
		funcName += "()"
	}
	if log.Flags()&log.Llongfile != 0 || log.Flags()&log.Lshortfile != 0 {
		return fmt.Sprintf("%s:", funcName)
	}
	// FuncName,FileLine
	return fmt.Sprintf("%s, %s:%d ", funcName, fileName, lineNo)
}
