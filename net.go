package gopp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"
)

// u: http://ip:port
func HttpClientSetProxy(hc *http.Client, pu string) *http.Client {
	pxyurl := pu
	urlo, err := url.Parse(pxyurl)
	ErrPanic(err, pxyurl)

	tp := hc.Transport.(*http.Transport)
	tp.Proxy = http.ProxyURL(urlo)
	cli := &http.Client{}
	cli.Transport = tp

	return hc
}
func HttpClientEnvProxy(hc *http.Client) *http.Client {
	pu := FindEvnProxy()
	if Empty(pu) {
		return hc
	}
	HttpClientSetProxy(hc, pu)
	return hc
}

type HttpClient struct {
	*http.Client

	Req    *http.Request
	Resp   *http.Response
	pxyurl string
}

func HttpClientNew() *HttpClient {
	return NewHttpClient()
}
func NewHttpClient() *HttpClient {
	cli := &HttpClient{}
	c := NewHttpClient2(0)
	HttpClientSetoptLite(c)
	cli.Client = c
	return cli
}
func (me *HttpClient) Timeout(timeo time.Duration) *HttpClient {
	me.Client.Timeout = timeo
	return me
}
func (me *HttpClient) KeepAlive(on bool) *HttpClient {
	tp := me.Client.Transport.(*http.Transport)
	tp.DisableKeepAlives = !on
	return me
}
func (me *HttpClient) Http2(on bool) *HttpClient {
	tp := me.Client.Transport.(*http.Transport)
	tp.ForceAttemptHTTP2 = on
	return me
}
func (me *HttpClient) SkipVerify(skip bool) *HttpClient {
	tp := me.Client.Transport.(*http.Transport)
	tlscfg := tp.TLSClientConfig
	tlscfg.InsecureSkipVerify = skip
	return me
}
func (me *HttpClient) Redirect(on bool) *HttpClient {
	if on {
		me.Client.CheckRedirect = nil // go default
	} else {
		me.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return me
}
func (me *HttpClient) Proxy(pu string) *HttpClient {
	me.pxyurl = pu
	HttpClientSetProxy(me.Client, pu)
	return me
}
func (me *HttpClient) ProxyEnv() *HttpClient {
	pu := FindEvnProxy()
	me.pxyurl = pu
	HttpClientSetProxy(me.Client, pu)
	return me
}
func (me *HttpClient) Getopts() map[string]any {

	tp := me.Client.Transport.(*http.Transport)
	tlscfg := tp.TLSClientConfig

	opts := map[string]any{}
	opts["timeout"] = me.Client.Timeout
	opts["skipverify"] = tlscfg.InsecureSkipVerify
	opts["keepalive"] = !tp.DisableKeepAlives
	opts["http2"] = tp.ForceAttemptHTTP2
	opts["redirect"] = me.Client.CheckRedirect == nil

	opts["proxy"] = me.pxyurl

	return opts
}
func (me *HttpClient) Do() (*http.Response, error) {
	resp, err := me.Client.Do(me.Req)
	me.Resp = resp
	return resp, err
}

func (me *HttpClient) ensurereq() {
	if me.Req == nil {
		req, err := http.NewRequest(http.MethodGet, "", nil)
		ErrPrint(err)
		me.Req = req
	}
}

func (me *HttpClient) Method(m, u string) *HttpClient {
	me.ensurereq()
	me.Req.Method = m
	uo, err := url.Parse(u)
	ErrPrint(err, u)
	me.Req.URL = uo
	// me.Req.RequestURI = u
	me.Req.Header = http.Header{}

	return me
}
func (me *HttpClient) Get(u string) *HttpClient {
	me.Method(http.MethodGet, u)
	return me
}
func (me *HttpClient) Post(u string) *HttpClient {
	me.Method(http.MethodPost, u)
	return me
}
func (me *HttpClient) Put(u string) *HttpClient {
	me.Method(http.MethodPut, u)
	return me
}
func (me *HttpClient) Head(u string) *HttpClient {
	me.Method(http.MethodHead, u)
	return me
}

func (me *HttpClient) BodyRaw(d []byte) *HttpClient {
	me.ensurereq()
	me.Req.Body = io.NopCloser(bytes.NewReader(d))
	me.HeaderKV("content-length", ToStr(len(d)))
	return me
}
func (me *HttpClient) BodyReader(d io.Reader) *HttpClient {
	me.ensurereq()
	me.Req.Body = io.NopCloser(d)
	return me
}
func (me *HttpClient) BodyFile(d string) *HttpClient {
	me.ensurereq()
	fo, err := os.Open(d)
	ErrPrint(err, d)
	me.Req.Body = fo
	fi, err := fo.Stat()
	ErrPrint(err, d)
	me.HeaderKV("content-length", ToStr(fi.Size()))
	return me
}
func (me *HttpClient) BodyJson(d any) *HttpClient {
	bcc, err := json.Marshal(d)
	ErrPrint(err, reflect.TypeOf(d))
	me.BodyRaw(bcc)
	me.HeaderKV("content-type", HttpCTJson)
	return me
}
func (me *HttpClient) PeekBody() io.ReadCloser {
	rx := me.Req.Body

	return rx
}

func (me *HttpClient) HeaderMap(d MapSS) *HttpClient {
	me.ensurereq()
	r := me.Req
	for k, v := range d {
		r.Header.Set(k, v)
	}
	return me
}
func (me *HttpClient) HeaderKV(kvs ...string) *HttpClient {
	me.ensurereq()
	r := me.Req
	for i := 0; i < len(kvs); i += 2 {
		k := kvs[i]
		v := kvs[i+1]
		r.Header.Set(k, v)
	}
	return me
}
func (me *HttpClient) HeaderValues(d url.Values) *HttpClient {
	me.ensurereq()
	r := me.Req
	for k, v := range d {
		r.Header.Set(k, v[0])
	}
	return me
}
func (me *HttpClient) Cookie(k, v string) *HttpClient {
	me.ensurereq()
	ck := http.Cookie{}
	ck.Name = k
	ck.Value = v
	me.Req.AddCookie(&ck)
	return me
}
func (me *HttpClient) Cookies(kvs ...string) *HttpClient {
	me.ensurereq()
	for i := 0; i < len(kvs); i += 2 {
		ck := http.Cookie{}
		ck.Name = kvs[i]
		ck.Value = kvs[i+1]
		me.Req.AddCookie(&ck)
	}
	return me
}
func (me *HttpClient) CookieJar(k, v string) *HttpClient {
	me.ensurereq()
	return me
}
func (me *HttpClient) CookieFile(k, v string) *HttpClient {
	me.ensurereq()
	return me
}

// after request
func (me *HttpClient) RemoteIP() string {
	return ""
}
func (me *HttpClient) ContentLength2() int64 {
	return 0
}
func (me *HttpClient) ContentType2() string {
	return ""
}
func (me *HttpClient) PeekBody2() io.ReadCloser {
	rx := me.Resp.Body

	return rx
}

// timeoms == 0, use default value
func NewHttpClient2(timeoms int) *http.Client {
	cli := &http.Client{}
	// tp1 := (http.DefaultTransport.(*http.Transport))
	// tp := tp1.Clone() // go1.13
	tp := &http.Transport{}
	tp.DisableCompression = false
	if timeoms > 0 {
		todur := time.Duration(timeoms) * time.Millisecond
		cli.Timeout = todur
		tp.TLSHandshakeTimeout = todur
	}

	tlscfg := &tls.Config{}
	tlscfg.InsecureSkipVerify = true
	tp.TLSClientConfig = tlscfg
	tp.DisableKeepAlives = true
	tp.MaxIdleConns = 3
	tp.ForceAttemptHTTP2 = false

	cli.Transport = tp

	return cli
}
func NewHttpClient3(timeoms int, idlecnt int, keepalive bool) *http.Client {
	cli := &http.Client{}
	// tp1 := (http.DefaultTransport.(*http.Transport))
	// tp := tp1.Clone() // go1.13
	tp := &http.Transport{}
	tp.DisableCompression = false
	if timeoms > 0 {
		todur := time.Duration(timeoms) * time.Millisecond
		cli.Timeout = todur
		tp.TLSHandshakeTimeout = todur
	}

	tlscfg := &tls.Config{}
	tlscfg.InsecureSkipVerify = true
	tp.TLSClientConfig = tlscfg
	tp.DisableKeepAlives = !keepalive
	tp.ForceAttemptHTTP2 = false
	tp.MaxIdleConns = IfElseInt(idlecnt == 0, http.DefaultMaxIdleConnsPerHost, idlecnt)

	cli.Transport = tp

	return cli
}

func CloseHttpIdles() {
	tpx := http.DefaultClient.Transport
	if tpx == nil {
		return
	}
	tp := tpx.(*http.Transport)
	tp.CloseIdleConnections()
	// http.DefaultClient.CloseIdleConnections() // go1.12
}

// only ip:port, if ip part is a domain name, this will fail
func ParseUdpAddr(address string) *net.UDPAddr {
	ta := ParseTcpAddr(address)
	ua := &net.UDPAddr{}
	ua.Port = ta.Port
	ua.IP = ta.IP
	return ua
}
func ParseTcpAddr(address string) *net.TCPAddr {
	ao := ParseAddr(address)
	return ao.(*net.TCPAddr)
}
func ParseAddr(address string) net.Addr {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil
	}
	iport, err := strconv.Atoi(port)
	if err != nil {
		return nil
	}
	ip := net.ParseIP(host)
	ao := &net.TCPAddr{}
	ao.IP = ip
	ao.Port = iport
	return ao
}

type FmtWriter struct {
	io.Writer
}

func NewFmtwriter(w io.Writer) *FmtWriter {
	this := &FmtWriter{w}
	return this
}
func (this *FmtWriter) Print(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(this.Writer, a...)
	return
}
func (this *FmtWriter) Printf(format string, a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(this.Writer, format, a...)
	return
}
func (this *FmtWriter) Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(this.Writer, a...)
	return
}

func HTWFlush(w http.ResponseWriter) bool {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
		return true
	}
	return false
}

func XferCopy(c1, c2 net.Conn, close bool) (int64, int64, error) {
	donec12 := make(chan bool, 1)
	donec21 := make(chan bool, 1)
	go func() {
		cpn, err := io.Copy(c2, c1)
		if ErrHave(err, "use of closed network connection") ||
			ErrHave(err, "read/write on closed pipe") {
		} else {
			ErrPrint(err, cpn)
		}
		donec12 <- true
	}()
	go func() {
		cpn, err := io.Copy(c1, c2)
		if ErrHave(err, "use of closed network connection") {
		} else {
			ErrPrint(err, cpn)
		}
		donec21 <- true
	}()

	// 双向转发的终止方法
	// first cycle, any done
	select {
	case <-donec12:
	case <-donec21:
	}
	// second cycle, all done or timeout
	select {
	case <-donec12:
	case <-donec21:
	case <-time.After(135 * time.Second):
	}

	c1.Close()
	c2.Close()
	return 0, 0, nil
}
