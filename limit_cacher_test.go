package gopp

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestLC0(t *testing.T) {
	lc := LimitCacherNew[string](3, 3*time.Second, func(t []string, full bool, timeo bool) {
		log.Println("notified", t, full, timeo)
	})
	for i := range 10 {
		lc.Add(fmt.Sprintf("log%d", i))
		SleepMs(800)
	}

	SleepSec(1)
	lc.Add("log888")
	lc.Flush()
	log.Println("hhehhe")
	SleepSec(1)

}
