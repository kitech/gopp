package gopp

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

const (
	HttpContentTypeHtml     = "text/html; charset=utf-8"
	HttpContentTypeJson     = "application/json; charset=utf-8"
	HttpContentTypeXml      = "application/xml; charset=utf-8"
	HttpContentTypeText     = "text/plain; charset=utf-8"
	HttpContentTypeStream   = "application/octet-stream"
	HttpContentTypeFormData = "form/data"

	HttpContentTypeXWWWUrlEncoded = "application/x-www-form-urlencoded"

	HttpUserAgentFF       = "Firefox 123"
	HttpUserAgentChrome   = "Chrome 456"
	HttpUserAgentChromium = "Chromium/86.0"
	HttpUserAgentCurl     = "curl/8.6.0"

	// HttpBasicAuthHeader = "Basic %s"

	HttpDateFmtStr = HttpDateFmt
)

const (
	HtmlNewline = "<br/>\n"
)

func _HttpBasicAuthHeader(key string) string {
	return ""
}

func HttpRespText(w http.ResponseWriter, code int, v string, cctype string) {
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	if cctype != "" {
		w.Header().Set("content-type", cctype)
	}
	// w.Header().Set("content-type", HttpContentTypeJson)
	// jdata, err := json.Marshal(v)
	// gopp.ErrPrint(err, reflect.TypeOf(v), v)
	wn, err := w.Write([]byte(v))
	ErrPrint(err, wn, code, cctype, len(v), cctype)
	log.Println("resp", wn, err, code, len(v), SubStr(v, 54))
	// log.Println("resp", wn, err, code, len(v), gopp.SubStr(string(v), 64))
}

func HttpRespJson(w http.ResponseWriter, code int, v any) error {
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	if v == nil {
		v = map[string]any{}
	}
	w.Header().Set("content-type", HttpContentTypeJson)
	jdata, err := json.Marshal(v)
	ErrPrint(err, reflect.TypeOf(v), v)
	wn, err := w.Write(jdata)
	ErrPrint(err, wn, code, reflect.TypeOf(v), len(jdata), SubStr(string(jdata), 64))

	log.Println("resp", wn, err, code, v, len(jdata))
	// log.Println("resp", wn, err, code, v, string(jdata))
	return err
}

func HttpRespXml(w http.ResponseWriter, code int, data any) {
	var bdata, err = xml.Marshal(data)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("content-length", fmt.Sprintf("%d", len(bdata)))
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(code)
		var n, err = w.Write(bdata)
		if err != nil {
			log.Println(err, n)
		}
	}
}
