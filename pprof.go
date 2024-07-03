package gopp

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
)

/*
usage: go tool pprof http://localhost:port/debug/pprof/profile
*/

func PprofEnable(port int, allip bool) {
	go func() {
		addr := fmt.Sprintf("%s:%d", IfElse2(allip, "0.0.0.0", "127.0.0.1"), port)
		hturl := fmt.Sprintf("http://%s", addr)
		Println("Gopprof listen", hturl, "rcips", Retn(GetLocalIPs()))
		err := http.ListenAndServe(addr, nil)
		ErrPrint(err, addr)
	}()
}

func GetLocalIPs() ([]net.IP, error) {
	var ips []net.IP
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
			if ipnet.IP.To16() != nil {
				// log.Println(ipnet.IP)
			}
		}
	}
	// log.Println(ips)
	return ips, nil
}
