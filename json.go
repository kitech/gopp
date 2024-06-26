package gopp

import (
	"encoding/json"
	"io"
	"strings"

	spjson "github.com/bitly/go-simplejson"
)

// //////
func JsonMarshalMust(v any) string {
	bcc, err := json.Marshal(v)
	ErrPrint(err, v)
	return string(bcc)
}

////////

type Spjson struct {
	*spjson.Json
}

func (me *Spjson) Origin() *spjson.Json {
	return me.Json
}
func originspjson2wrap(jo *spjson.Json) *Spjson {
	return &Spjson{jo}
}

// key1.key2.key3
func (me *Spjson) GetPathDot(paths string) *Spjson {
	jo := me.Origin()
	rjo := jo.GetPath(strings.Split(paths, ".")...)
	return originspjson2wrap(rjo)
}

// key1.idx1.key2.idx2
func (me *Spjson) GetPathIndexDot(pathidxs string) *Spjson {
	jo := me.Origin()

	var rv = jo
	for i, p := range strings.Split(pathidxs, ".") {
		if IsInteger(p) {
			idx := MustInt(p)
			if idx < 0 {
				Warnp("idx < 0", idx, i, pathidxs)
				rv = jo.GetIndex(idx) // what error?
			} else {
				rv = jo.GetIndex(idx)
			}
		} else {
			rv = jo.Get(p)
		}
	}
	return originspjson2wrap(rv)
}

func JsonNew(body []byte) (*Spjson, error) {
	jo, err := spjson.NewJson(body)
	if err != nil {
		return nil, err
	}
	me := originspjson2wrap(jo)
	return me, err
}
func JsonNewEmpty() *Spjson {
	me := originspjson2wrap(spjson.New())
	return me
}
func Spjsonof(jo *spjson.Json) *Spjson {
	return originspjson2wrap(jo)
}

//////

type Json struct {
	jo *spjson.Json
}

func NewEmptyJson() *Json {
	return &Json{spjson.New()}
}

func NewJson(body []byte) (*Json, error) {
	j, err := spjson.NewJson(body)
	return &Json{j}, err
}

func NewJsonFromReader(r io.Reader) (*Json, error) {
	j, err := spjson.NewFromReader(r)
	return &Json{j}, err
}

func NewJsonFromObject(j *spjson.Json) *Json {
	return &Json{j}
}

func (j *Json) Ori() *spjson.Json { return j.jo }

func (j *Json) GetPathDot(branch string) *Json {
	return &Json{j.jo.GetPath(strings.Split(branch, ".")...)}
}

// just wrapper
func (j *Json) GetPath(branch ...string) *Json {
	return &Json{j.jo.GetPath(branch...)}
}

func (j *Json) GetIndex(index int) *Json {
	return &Json{j.jo.GetIndex(index)}
}

func (j *Json) Get(key string) *Json {

	return &Json{j.jo.Get(key)}
}

func (j *Json) CheckGet(key string) (*Json, bool) {
	o, ok := j.jo.CheckGet(key)
	return &Json{o}, ok
}

func (j *Json) CheckGetDot(key string) (*Json, bool) {
	ps := strings.Split(key, ".")

	jo := j.jo
	tok := false
	for i := 0; i < len(ps); i++ {
		tok = false
		if tmpo, ok := jo.CheckGet(ps[i]); ok {
			jo = tmpo
			tok = true
		} else {
			return nil, false
		}
	}

	return &Json{jo}, tok
}

func (j *Json) MustInt(args ...int) int {
	return j.jo.MustInt(args...)
}

func (j *Json) MustString(args ...string) string {
	return j.jo.MustString(args...)
}

func (j *Json) MustArray(args ...[]interface{}) []interface{} {
	return j.jo.MustArray(args...)
}

func (j *Json) MustMap(args ...map[string]interface{}) map[string]interface{} {
	return j.jo.MustMap(args...)
}

// /
func JsonDecStrMap(data []byte) (map[string]string, error) {
	ret := map[string]string{}
	err := json.Unmarshal(data, &ret)
	return ret, err
}
