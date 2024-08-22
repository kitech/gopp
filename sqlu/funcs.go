package sqlu

import "github.com/kitech/gopp"

func IsDupkey(err error) bool {
	if err == nil {
		return false
	}
	if gopp.ErrHave(err, "UNIQUE constraint failed") {
		return true
	}
	return false
}
