package gopp

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"slices"
	"strings"
	"time"
)

const (
	HtmlNewline = " <br/>\n"
)
const (
	HttpCTHtml         = "text/html; charset=utf-8"
	HttpCTJson         = "application/json; charset=utf-8"
	HttpCTAppXml       = "application/xml; charset=utf-8"
	HttpCTXml          = "text/xml; charset=utf-8"
	HttpCTText         = "text/plain; charset=utf-8"
	HttpCTStream       = "application/octet-stream"
	HttpCTFormData     = "form/data"
	HttpBasicAuthTmpl  = "Authorization: Basic %s"  // %s=base64encode(user:pass)
	HttpBearerAuthTmpl = "Authorization: Bearer %s" // %s=base64encode(user:pass)

	HttpCTXWWWUrlEncoded = "application/x-www-form-urlencoded"

	HttpUACurl     = "curl/8.6.0"
	HttpUAFirefox  = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"
	HttpUAChrome   = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36"
	HttpUAChromium = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chromium/28.0.1500.52 Safari/537.36"
	HttpUAAndroid  = "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	HttpUAMacos    = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"

	XmlHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>"

	HttpDateFmtStr = HttpDateFmt
)

var HttpUserAgents = []string{HttpUACurl, HttpUAFirefox, HttpUAChrome, HttpUAChromium, HttpUAAndroid, HttpUAMacos}

func HttpBasicAuthVal(user, pass string) string {
	val := fmt.Sprintf("%s:%s", user, pass)
	encval := base64.StdEncoding.EncodeToString([]byte(val))
	return fmt.Sprintf(HttpBasicAuthTmpl, encval)
}
func HttpBasicAuthKV(user, pass string) (string, string) {
	val := fmt.Sprintf("%s:%s", user, pass)
	encval := base64.StdEncoding.EncodeToString([]byte(val))
	return "Authorization", fmt.Sprintf("Basic %s", encval)
}
func HttpBasicAuthLine(user, pass string) string {
	k, v := HttpBasicAuthKV(user, pass)
	return fmt.Sprintf("%v: %v", k, v)
}

func HttpBearerAuthVal(user, pass string) string {
	val := fmt.Sprintf("%s:%s", user, pass)
	encval := base64.StdEncoding.EncodeToString([]byte(val))
	return fmt.Sprintf("Bearer %s", encval)
}
func HttpBearerAuthKV(user, pass string) (string, string) {
	val := fmt.Sprintf("%s:%s", user, pass)
	encval := base64.StdEncoding.EncodeToString([]byte(val))
	return "Authorization", fmt.Sprintf("Bearer %s", encval)
}
func HttpBearerAuthToken(token string) (string, string) {
	return "Authorization", fmt.Sprintf("Bearer %s", token)
}
func HttpBearerAuthLine(user, pass string) string {
	k, v := HttpBearerAuthKV(user, pass)
	return fmt.Sprintf("%v: %v", k, v)
}

// parse range: bytes=87458121-97175688
func HttpRangeParse(v string) (int64, int64, error) {

	return -1, -1, nil
}
func HttpRangeParseMust(v string) (int64, int64) {
	b, e, err := HttpRangeParse(v)
	ErrPrint(err, v)
	return b, e
}
func HttpRangeReq(bpos, epos int64) (v string) {
	v = fmt.Sprintf("range: bytes=%v-%v", bpos, epos)
	return
}
func HttpRangeResp(bpos, epos, tlen int64) (v string) {
	// Content-Range: bytes 42-1233/1234
	v = fmt.Sprintf("Content-Range: bytes %v-%v/%v", bpos, epos, tlen)
	return
}
func HttpRangeParseReq(v string) (int64, int64, int64, error) {
	return -1, -1, -1, nil
}

func Httpcode2err(code int) error {
	err := fmt.Errorf("%v %v", code, http.StatusText(code))
	return err
}

func HttpRespRaw(w http.ResponseWriter, code int, v []byte, cctype string, headers map[string]string) error {
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	if cctype != "" {
		w.Header().Set("content-type", cctype)
	}
	w.Header().Set("content-length", ToStr(len(v)))
	// w.Header().Set("content-type", HttpContentTypeJson)
	// jdata, err := json.Marshal(v)
	// gopp.ErrPrint(err, reflect.TypeOf(v), v)
	wn, err := w.Write(v)
	ErrPrint(err, wn, code, cctype, len(v), cctype)
	log.Println("resp", wn, err, code, len(v), SubStr(string(v), 54))
	// log.Println("resp", wn, err, code, len(v), gopp.SubStr(string(v), 64))
	return err
}
func HttpRespText(w http.ResponseWriter, code int, v string, cctype string, headers map[string]string) error {
	return HttpRespRaw(w, code, []byte(v), cctype, headers)
}

func HttpRespJson(w http.ResponseWriter, code int, v any, headers map[string]string) error {

	jdata, err := json.Marshal(v)
	ErrPrint(err, reflect.TypeOf(v), v)
	wn, err := w.Write(jdata)
	ErrPrint(err, wn, code, reflect.TypeOf(v), len(jdata), SubStr(string(jdata), 64))

	return HttpRespRaw(w, code, jdata, HttpCTJson, headers)
}

func HttpRespXml(w http.ResponseWriter, code int, data any, headers map[string]string) error {
	var bdata, err = xml.Marshal(data)
	ErrPrint(err)
	if err == nil {
		err = HttpRespRaw(w, code, bdata, HttpCTText, headers)
	}
	return err
}

// opt: timeout, nokeepalive. skipverify, nohttp2, noredir
func HttpClientSetoptLite(htcli *http.Client) *http.Client {
	tpx := htcli.Transport
	var tp *http.Transport
	if tpx != nil {
		// http.DefaultTransport
		tp = tpx.(*http.Transport)
	} else {
		tp = &http.Transport{}
	}
	tp.DisableKeepAlives = true

	var tlscfg *tls.Config = tp.TLSClientConfig
	if tlscfg == nil {
		tlscfg = &tls.Config{InsecureSkipVerify: true}
		tp.TLSClientConfig = tlscfg
	}
	tp.MaxConnsPerHost = 0
	tp.MaxIdleConns = 1
	tp.MaxIdleConnsPerHost = 1
	tp.IdleConnTimeout = 1 * time.Millisecond
	tp.ForceAttemptHTTP2 = false
	htcli.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	_ = http.DefaultMaxIdleConnsPerHost
	// htcli.CloseIdleConnections()
	HttpClientEnvProxy(htcli)
	return htcli
}

func FindEvnProxy() string {
	envpxys := []string{"http_proxy", "https_proxy", "socks_proxy", "socks4_proxy", "socks5_proxy"}
	for _, line := range os.Environ() {
		pos := strings.Index(line, "=")
		name := strings.ToLower(line[:pos])
		if slices.Contains(envpxys, name) {
			return line[:pos+1]
		}
	}
	return ""
}
