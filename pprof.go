package gopp

import (
	"fmt"
	"net"
	"net/http"

	// "net/http/pprof"
	_ "net/http/pprof"
	"runtime"
	rtdbg "runtime/debug"
	rtpprof "runtime/pprof"
)

/*
usage: go tool pprof http://localhost:port/debug/pprof/profile
*/

const PprofPortDefault = 3860

var PprofPort = PprofPortDefault
var PprofUrl = fmt.Sprintf("http://%s:%d/debug/pprof", "127.0.0.1", PprofPortDefault)

// 使用默认的http server handler
func PprofEnable(port int, allip bool) {
	PprofPort = IfElseInt(port <= 0, PprofPortDefault, port)
	PprofUrl = fmt.Sprintf("http://%s:%d/debug/pprof", "127.0.0.1", PprofPort)
	threadProfile = rtpprof.Lookup("threadcreate")
	setupRuntimeMemoryGCTuner()
	go func() {
		addr := fmt.Sprintf("%s:%d", IfElse2(allip, "0.0.0.0", "127.0.0.1"), port)
		hturl := fmt.Sprintf("http://%s", addr)
		Println("Gopprof listen", hturl, "rcips", Retn(GetLocalIPs()))
		err := http.ListenAndServe(addr, nil)
		ErrPrint(err, addr)
	}()
}

var threadProfile *rtpprof.Profile // rtpprof.Lookup("threadcreate")

func setupRuntimeMemoryGCTuner() {
	cp := NewCodePager()
	sec := "main"
	http.HandleFunc("/tuner", func(w http.ResponseWriter, r *http.Request) {
		maxosth := rtdbg.SetMaxThreads(128) // <=0 crash
		maxmem := rtdbg.SetMemoryLimit(-1)
		gcpcnt := rtdbg.SetGCPercent(100)

		cp.APf(sec, "rcvars: [maxproc:0+, maxmem:MB, maxosth:0+, gcpercent:0-100").Nlweb(sec).Nlweb(sec)
		cp.APf(sec, "Current rtvars:").Nlweb(sec)
		cp.APf(sec, "MAXPROCS: %v", runtime.GOMAXPROCS(0)).Nlweb(sec)
		cp.APf(sec, "Osth: curr/max: %v/%v", threadProfile.Count(), maxosth).Nlweb(sec)
		cp.APf(sec, "MemLimit: %vM", maxmem/MB).Nlweb(sec)
		cp.APf(sec, "GCPercent/GOGC: %v", gcpcnt).Nlweb(sec)

		// restore
		rtdbg.SetMaxThreads(maxosth)
		rtdbg.SetMemoryLimit(maxmem)
		rtdbg.SetGCPercent(gcpcnt)

		res := cp.ExportAll()
		w.Write([]byte(res))
	})
	// setvar?{maxproc=3, maxosth=(1-32), maxmem=30(M), percent=(0-100), }
	http.HandleFunc("/tuner/setvar", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		cp := NewCodePager()
		sec := "main"
		for key, vals := range r.Form {
			val := MustInt(vals[0])
			oldval := -1
			switch key {
			case "maxproc":
				oldval = runtime.GOMAXPROCS(val)
			case "maxosth":
				if val > 0 && val < 10000 {
					oldval = rtdbg.SetMaxThreads(val)
				} else {
					Warn("invalid val: want 1-10000, but", val)
				}
			case "maxmem":
				oldval = int(rtdbg.SetMemoryLimit(int64(val) * MB))
				oldval /= MB
			case "gcpercent":
				oldval = rtdbg.SetGCPercent(val)
			default:
				Warn("nocat", key, vals, r.URL.String())
			}
			cp.APf(sec, "var %v => %v", oldval, val).Nlweb(sec)
		}
		res := cp.ExportAll()
		w.Write([]byte(res))
	})

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
