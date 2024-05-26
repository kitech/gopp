package gopp

import (
	"fmt"
	"testing"
	"time"
)

func TestRetryer0(t *testing.T) {

	{
		rter := NewRetry()
		n, dur := rter.NextWait()
		// log.Println(n, dur)
		if n != 0 || dur != 0*time.Millisecond {
			t.Error("wait error", n, dur)
		}
	}

	{
		rter := NewRetry(2)
		n, dur := rter.NextWait()
		// log.Println(n, dur)
		if n != 0 || dur != 0*time.Millisecond {
			t.Error("wait error", n, dur)
		}
	}

	{

		var cnt = 0
		rter := NewRetryFnOnly(func() error {
			cnt += 1
			if cnt > 3 {
				return nil
			}
			return fmt.Errorf("tst %d", cnt)
		})
		rter.Do()
		// log.Println(cnt)
		if cnt != 4 {
			t.Error("want retry 4 times, but", cnt)
		}
	}

	{

		var cnt = 0
		rter := NewRetryFnOnly(func() error {
			cnt += 1
			if cnt > 6 {
				return nil
			}
			return fmt.Errorf("tst %d", cnt)
		})
		rter.Do(3)
		// log.Println(cnt)
		if cnt != 4 {
			t.Error("want retry 4 times, but", cnt)
		}
	}

	{

		var cnt = 0
		rter := NewRetryFnOnly(func() error {
			cnt += 1
			if cnt > 6 {
				return nil
			}
			return fmt.Errorf("tst %d", cnt)
		}, 2)
		rter.Do(3)
		// log.Println(cnt)
		if cnt != 4 {
			t.Error("want retry 4 times, but", cnt)
		}
	}
}
