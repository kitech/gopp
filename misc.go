package gopp

import (
	"fmt"
)

const (
	KB = 1024
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	ZB = PB * KB
)

// todo move to bytes package
func Bytes2Hum[T int64 | uint64 | int | usize](bv T) string {
	const units = "ZPTGMKB"
	var values = [...]uint64{ZB, PB, TB, GB, MB, KB, 0}

	var rv string
	for i, val := range values {
		// log.Println(i, val, bv)
		if uint64(bv) >= val {
			if val == 0 {
				val = 1
			}
			if val == 1 {
				rv += fmt.Sprintf("%d", uint64(bv)/val)
			} else {
				rv += fmt.Sprintf("%d%v.", uint64(bv)/val, string(units[i]))
			}
			bv = T(uint64(bv) % uint64(val))
		}
	}
	return rv
}

// 简短格式， 23.123M
func Bytes2Humz[T int64 | uint64 | int](bv T) string {
	const units = "ZPTGMKB"
	var values = [...]uint64{ZB, PB, TB, GB, MB, KB, 0}

	var rv string
	for i, val := range values {
		// log.Println(i, val, bv)
		if uint64(bv) >= val {
			if val == 0 {
				val = 1
				rv += fmt.Sprintf("%dB", uint64(bv)/val)
			} else if values[i+1] == 0 {
				bvnxt := T(uint64(bv) % uint64(val))
				rv += fmt.Sprintf("%d.%d%s", uint64(bv)/val, uint64(bvnxt)/1, string(units[i]))
			} else {
				bvnxt := T(uint64(bv) % uint64(val))
				rv += fmt.Sprintf("%d.%d%s", uint64(bv)/val, uint64(bvnxt)/values[i+1], string(units[i]))
			}
			break
		}
	}
	return rv
}
