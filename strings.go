package gopp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	_ "github.com/huandu/xstrings"
)

// 安全提取子字符串。支持负值，表示从后面
func SubStr(s string, n int) string {
	if n < 0 {
		absn := AbsI32(n)
		if absn > len(s) {
			return s
		}
		return s[len(s)-absn:]
	} else {
		if n >= len(s) {
			return s
		}
		return s[:n]
	}
}

func StrSuf(s string, n int) string {
	r := SubStr(s, n)
	if len(r) < len(s) {
		return r + "..."
	}
	return r
}

// for ui, ascii = 1, else 2
func StrSuf4ui(s string, n int, center ...int) string {
	r := ""
	rlen := 0
	for _, c := range s {
		rlen += IfElseInt(c < 128, 1, 2)
		r += IfElseStr(rlen > n, IfElseStr(len(center) > 0, " …", "..."), string(c))
		if rlen > n {
			break
		}
	}
	return r
}

func SubBytes(p []byte, n int) []byte {
	if n >= len(p) {
		return p
	}
	return p[:n]
}

// 按长度切割字符串
func Splitn(s string, n int) []string {
	v := make([]string, 0)
	for i := 0; i < (len(s)/n)+1; i++ {
		bp := i * n
		ep := bp + n
		if bp >= len(s) {
			break
		}
		if ep > len(s) {
			ep = len(s)
		}

		v = append(v, s[bp:ep])
	}
	return v
}

// rune support
func Splitrn(s string, n int) []string {
	v := make([]string, 0)

	sub := ""
	sublen := 0
	for _, c := range s {
		cs := string(c)
		if sublen+len(cs) > n {
			v = append(v, sub)
			sub = ""
			sublen = 0
		}

		sub += cs
		sublen += len(cs)
	}

	if sublen > 0 {
		v = append(v, sub)
	}
	return v
}

// rune support, utf8 3byte, but ui width is 2
func Splitrnui(s string, n int) []string {
	v := make([]string, 0)

	sub := ""
	sublen := 0
	for _, c := range s {
		cs := string(c)
		uilen := IfElseInt(len(cs) == 1, 1, 2)
		if sublen+uilen > n {
			v = append(v, sub)
			sub = ""
			sublen = 0
		}

		sub += cs
		sublen += uilen
	}

	if sublen > 0 {
		v = append(v, sub)
	}
	return v
}

// line support
// TODO one line exceed n???
func Splitln(s string, n int) []string {
	return Splitsn(s, n, "\n")
}

// string seperator support
func Splitsn(s string, n int, sep string) []string {
	v := make([]string, 0)

	ls := strings.Split(s, sep)

	sub := ""
	sublen := 0
	for _, line := range ls {
		if sublen+1+len(line) > n {
			v = append(v, sub)
			sub = ""
			sublen = 0
		}

		sub += line + "\n"
		sublen += len(line) + 1
	}

	if sublen > 0 {
		v = append(v, sub)
	}
	return v
}

func StrPrepend(s string, b byte) string {
	return string(append([]byte{b}, bytes.NewBufferString(s).Bytes()...))
}

func StrPrepend2(s string, b byte) string {
	return string([]byte{b}) + s
}

func StrReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 仅Title第一个字节
func Title(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

// 仅Title第一个字节
func Untitle(s string) string {
	if len(s) > 0 {
		return strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

func IsNumberic(s string) bool {
	if strings.Count(s, ".") > 1 {
		return false
	}
	for _, c := range s {
		if unicode.IsNumber(c) || c == '.' {
		} else {
			return false
		}
	}
	return true
}

func IsInteger(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsFloat(s string) bool {
	for _, c := range s {
		if unicode.IsDigit(c) || c == '.' {
		} else {
			return false
		}
	}

	return true
}

func IsPrint(s string) bool {
	for _, c := range s {
		if !unicode.IsPrint(rune(c)) {
			return false
		}
	}
	return true
}

// type String struct{ s string }
/*
func NewString(s string) *String { return &String{s} }
func (this *String) Raw() string { return this.s }

func (this *String) Mid(from, length int) *String { return NewString(this.s[from:length]) }
*/

// 以类方法的方式使用string相关函数，使用时可以拷贝过去
// 不过还是有许多代码要写的

type Str string

func (this Str) Mid(from, length int) Str { return Str(this[from:length]) }
func CutSuffix(s string, cutlen int) string {
	return s[:len(s)-cutlen]
}

// why there is a newline suffix?
func JsonEncode(v interface{}) (js string, err error) {
	w := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	err = enc.Encode(v)
	js = string(w.Bytes())
	return
}

func StrHaveNocase(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
func StrsHaveNocase(ss []string, s string) bool {
	es := strings.ToLower(s)
	for _, v := range ss {
		if strings.ToLower(v) == es {
			return true
		}
	}
	return false
}

func StrElideMid(s string, wtlen int) string {
	if len(s) <= wtlen {
		return s
	}
	ns := fmt.Sprintf("%s..%s", s[:wtlen/2], s[len(s)-wtlen/2:])
	return ns
}

// todo
func Camel2Snak(s string) string {
	return s
}
func Snak2Camel(s string) string {
	return s
}
