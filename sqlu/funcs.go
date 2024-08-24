package sqlu

import (
	"database/sql"

	"github.com/kitech/gopp"
)

func IsDupkey(err error) bool {
	if err == nil {
		return false
	}
	if gopp.ErrHave(err, "UNIQUE constraint failed") {
		return true
	}
	return false
}

func DBSetExtraOpts(db *sql.DB) {

}

func SqliteDBSetExtraOpts(db *sql.DB) {

}
