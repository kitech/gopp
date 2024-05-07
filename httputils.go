package gopp

import (
	"encoding/base64"
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
	HttpBasicAuthHeaderTmpl = "Authorization: Basic %s" // %s=base64encode(user:pass)

	HttpContentTypeXWWWUrlEncoded = "application/x-www-form-urlencoded"

	HttpUserAgentCurl     = "curl/8.6.0"
	HttpUserAgentFirefox  = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"
	HttpUserAgentChrome   = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36"
	HttpUserAgentChromium = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chromium/28.0.1500.52 Safari/537.36"
	HttpUserAgentAndroid  = "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	HttpUserAgentMacos    = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"

	HttpDateFmtStr = HttpDateFmt
)

const (
	HtmlNewline = "<br/>\n"
)

func HttpBasicAuthHeader(user, pass string) string {
	val := fmt.Sprintf("%s:%s", user, pass)
	encval := base64.StdEncoding.EncodeToString([]byte(val))
	return fmt.Sprintf(HttpBasicAuthHeaderTmpl, encval)
}
func HttpBasicAuthHeader2(user, pass string) string {
	val := fmt.Sprintf("%s:%s", user, pass)
	encval := base64.StdEncoding.EncodeToString([]byte(val))
	return fmt.Sprintf("Basic %s", encval)
}
func HttpBasicAuthHeader3(user, pass string) (string, string) {
	val := fmt.Sprintf("%s:%s", user, pass)
	encval := base64.StdEncoding.EncodeToString([]byte(val))
	return "Authorization", fmt.Sprintf("Basic %s", encval)
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
