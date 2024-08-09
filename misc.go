package gopp

import (
	"fmt"
)

const (
	KB = 1024
	MB = 1024 * 1024
	GB = 1024 * 1024 * 1024
	TB = 1024 * 1024 * 1024 * 1024
	PB = 1024 * 1024 * 1024 * 1024 * 1024
)

func Bytes2Hum[T int64 | uint64 | int | usize](bv T) string {
	const units = "PTGMKB"
	var values = [...]uint64{PB, TB, GB, MB, KB, 1}

	var rv string
	for i, val := range values {
		// log.Println(i, val, bv)
		if uint64(bv) >= val {
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
	const units = "PTGMKB"
	var values = [...]uint64{PB, TB, GB, MB, KB, 1}

	var rv string
	for i, val := range values {
		// log.Println(i, val, bv)
		if uint64(bv) >= val {
			if val == 1 {
				rv += fmt.Sprintf("%dB", uint64(bv)/val)
			} else {
				bvnxt := T(uint64(bv) % uint64(val))
				rv += fmt.Sprintf("%d.%d%s", uint64(bv)/val, uint64(bvnxt)/values[i+1], string(units[i]))
			}
			break
		}
	}
	return rv
}
