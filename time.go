package gopp

import (
	"fmt"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
)

// 中文常用格式
func TimeToFmt1(t time.Time) string { return t.Format(CleanDateFmt) }
func TimeToFmt1Now() string         { return time.Now().Format(CleanDateFmt) }
func TimeFromFmt1(s string) (time.Time, error) {
	return time.Parse(CleanDateFmt, s)
}
func TimeRfc3389ms(t time.Time) string { return t.Format(Rfc3339DatemsFmtStr) }
func TimeRfc7231(t time.Time) string   { return t.Format(Rfc7231DateFmtStr) }

// rounded float point part
// origin 8h38m46.115296675s
// now 8h38m46.115s
func Duround(d time.Duration) string {
	return d.String()
}

const CleanDateFmt = "2006-01-02 15:04:05"
const CleanDatemsFmt = "2006-01-02 15:04:05.999"
const HttpDateFmt = "Mon, 02 Jan 2006 15:04:05 GMT" // "Sat, 30 Sep 2017 00:10:59 GMT"
const DavDateFmt = HttpDateFmt
const Rfc7231DateFmtStr = HttpDateFmt
const Rfc3339DatemsFmtStr = "2006-01-02T15:04:05.999Z"
const Rfc3339DateFmtStr = "2006-01-02T15:04:05Z"

var StartTime = time.Now()
var ZeroTime time.Time

func Dur2hum(d time.Duration) string {
	unitMeasures := []time.Duration{365 * 24 * time.Hour, 12 * 24 * time.Hour, 24 * time.Hour, time.Hour, time.Minute, time.Second, time.Millisecond, time.Microsecond, time.Nanosecond}
	// unitWords := []string{"year", "month", "day", "hour", "minute", "second", "millisecond", "microsecond", "nanosecond"}
	unitShorts := []string{"y", "M", "d", "h", "m", "s", "ms", "µs", "ns"}

	sawsec := false
	r := ""
	for idx, du := range unitMeasures {
		m := d.Nanoseconds() / du.Nanoseconds()
		if m == 0 {
		} else {
			// 5s100ms => 5.100s
			unit := unitShorts[idx]
			sawsec = IfElse(sawsec, sawsec, unit == "s").(bool)
			var sfx = unit
			switch unit {
			case "s":
				sfx = "."
			case "ms":
				sfx = IfElseStr(sawsec, "s", unit)
			}
			r += fmt.Sprintf("%d%s", m, sfx)
		}
		d -= time.Duration(m) * du
		if idx >= 6 {
			break
		}
	}
	if strings.HasSuffix(r, ".") {
		r = r[:len(r)-1] + "s"
	}

	return r
}
func SinceHum(t time.Time) string { return Dur2hum(time.Since(t)) }

// offset [0-14)
func SetTimezone(offset int) (olocal *time.Location) {
	olocal = time.Local
	if offset >= 0 && offset < 24 {
		secondsEastOfUTC := int((time.Duration(offset) * time.Hour).Seconds())
		zone := time.FixedZone(fmt.Sprintf("UTC+%d", offset), secondsEastOfUTC)
		time.Local = zone
	}
	return
}
func SetLocal(offset int) *time.Location { return SetTimezone(offset) }

func CondWait(timeoutms int, f func() bool) {
	for {
		time.Sleep(time.Duration(timeoutms) * time.Millisecond)
		if f() {
			break
		}
	}
}

func TimeUnixMS(t time.Time) int64 {
	return t.Unix()*1000 + int64(t.Nanosecond()/1000000)
}
func TimeUnixMSStr(t time.Time) string { return fmt.Sprintf("%d", TimeUnixMS(t)) }
func TimeUnixMSNow() int64 {
	t := time.Now()
	return t.Unix()*1000 + int64(t.Nanosecond()/1000000)
}
func TimeUnixMSStrNow() string { return fmt.Sprintf("%d", TimeUnixMSNow()) }
func TimeFromUnixMS(tsms int64) time.Time {
	return time.Time(time.Unix(tsms/1000, tsms%1000*1000000))
}
func TimeFromUnixMSStr(tsms string) (time.Time, error) {
	ts, err := strconv.ParseInt(tsms, 10, 64)
	return TimeFromUnixMS(ts), err
}

func SleepSec(sec int) { time.Sleep(time.Duration(sec) * time.Second) }
func SleepMs(msec int) { time.Sleep(time.Duration(msec) * time.Millisecond) }
func SleepUs(usec int) { time.Sleep(time.Duration(usec) * time.Microsecond) }

func DurandMs(basems int, rdms int) time.Duration {
	var rdval = 0
	if rdms != 0 {
		rdval = mrand.Int() % rdms
		rdval = Abs(rdval)
	}
	return time.Duration(basems+rdval) * time.Millisecond
}
func DurandSec(basesec int, rdsec int) time.Duration {
	var rdval = 0
	if rdsec != 0 {
		rdval = mrand.Int() % rdsec
		rdval = Abs(rdval)
	}
	return time.Duration(basesec+rdval) * time.Second
}

// 写法比较：
// 1*time.Second
// gopp.Secondof(1)
func Secondof(n int) time.Duration  { return time.Second * time.Duration(n) }
func Msecondof(n int) time.Duration { return time.Millisecond * time.Duration(n) }
