package gopp

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

// u: http://ip:port
func ProxyHttpClient(u string) *http.Client {
	tp := &http.Transport{}
	pxyurl := u
	urlo, err := url.Parse(pxyurl)
	if err != nil {
		log.Panicln(err, pxyurl)
	}

	tp.Proxy = http.ProxyURL(urlo)
	cli := &http.Client{}
	cli.Transport = tp

	return cli
}

type HttpClient struct {
	c *http.Client
}

func NewHttpClient() *HttpClient {

	return &HttpClient{}
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
