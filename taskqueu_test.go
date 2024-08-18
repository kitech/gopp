package gopp

import (
	"log"
	"testing"
)

func TestTqez1(t *testing.T) {
	cm := NewTaskqez()

	cm.Add(func() {
		log.Println("tq111")
		SleepMs(13)
	})
	SleepMs(12)
	cm.Add(func() {
		log.Println("tq222")
		SleepMs(23)
	})
	SleepMs(3)
	cm.Add(func() {
		log.Println("tq333")
	})

	SleepSec(1)
}
func TestTqez2(t *testing.T) {

}

func TestCmer1(t *testing.T) {
	cm := NewCallMerger()
	cm.Add("test1", DurandMs(123, 123), func() {
		log.Println("tq111")
	})
	SleepMs(12)
	cm.Add("test1", DurandMs(123, 123), func() {
		log.Println("tq222")
	})
	SleepMs(23)
	cm.Add("test1", DurandMs(123, 123), func() {
		log.Println("tq333")
	})

	SleepSec(1)
}
