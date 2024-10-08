package sqlu

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/kitech/gopp"
)

type IDBTableez interface {
	Table() string
	Keys() []string
}

// todo move to sqlu
// 现在与db的命名方式的映射，就是只把首字母大写，其他不变
// todo 结构体字段名和数据库字段名不一样的情况
// todo 带tag的结构体
// 可用于update/select/delete
// 如果是insert, wheres is values
// keys 只对非 insert 生效
// vx 支持结构体指针或者非指针
func UpsqlFields(isinsert bool, vx any, keys ...string) (sets string, wheres string, binds []any, err error) {
	var binds_ []any
	var binds2_ []any // where binds
	var sets_ []string
	var wheres_ []string

	vo := reflect.ValueOf(vx)
	vt := vo.Type()
	if vt.Kind() == reflect.Pointer {
		vo = vo.Elem()
		vt = vt.Elem()
	}

	// log.Println(vo, vt)
	for i := 0; i < vo.NumField(); i++ {
		fo := vt.Field(i)
		fname := fo.Name
		fname4db := gopp.Untitle(fname)
		fv := vo.FieldByName(fname)
		// 现在与db的命名方式的映射，就是只把首字母大小，其他不变
		// todo use field tagname
		// log.Println(fname, fv.Interface(), gopp.Retn(fo.Tag.Get("gorm")), gopp.Retn(fo.Tag.Get("dbm")))
		fvx := fv.Interface()
		if !gopp.Empty(fvx) {
			// 特殊类型转换
			switch tv := fvx.(type) {
			case time.Time:
				fvx = tv.UnixMilli()
			}

			if isinsert {
				sets_ = append(sets_, fname4db)
				wheres_ = append(wheres_, "?")
				binds_ = append(binds_, fvx)
			} else {
				if gopp.StrsHaveNocase(keys, fname4db) {
					wheres_ = append(wheres_, fname4db+"=?")
					binds2_ = append(binds2_, fvx)
				} else {
					// todo if != ? then set
					sets_ = append(sets_, fname4db+"=?")
					binds_ = append(binds_, fvx)
				}
			}
		}

	}

	sets = strings.Join(sets_, ", ")
	// todo where cond not only `and`
	wheres = strings.Join(wheres_, gopp.IfElse2(isinsert, ", ", " and "))

	binds = append(binds_, binds2_...)

	// simple check
	if isinsert {
		if len(wheres_) != len(binds) || len(sets_) != len(binds) {
			log.Println("maybe somerr???", len(sets), len(wheres_), len(binds))
		}
		gopp.ZeroPrint(len(sets), "nothing sets???", vx)
	} else {
		if len(wheres_)+len(sets_) != len(binds) {
			log.Println("maybe somerr???", len(sets_), len(wheres_), len(binds), wheres_, binds)
			// log.Printf("%+#v\n", vx)
		}
		gopp.ZeroPrint(len(binds2_), "empty wheres???", keys, vx)
	}
	gopp.ZeroPrint(len(binds), "nothing binds???", vx)

	return
}

// SelsqlFields 和 UpsqlFields 还能合并吗，处理方式着实不同
func SelsqlFields(vx any, keys ...string) (wheres string, binds []any, err error) {
	var binds2_ []any // where binds
	var wheres_ []string

	vo := reflect.ValueOf(vx)
	vt := vo.Type()
	if vt.Kind() == reflect.Pointer {
		vo = vo.Elem()
		vt = vt.Elem()
	}

	for i := 0; i < vo.NumField(); i++ {
		fo := vt.Field(i)
		fname := fo.Name
		fname4db := gopp.Untitle(fname)
		fv := vo.FieldByName(fname)
		// 现在与db的命名方式的映射，就是只把首字母大小，其他不变
		// todo use field tagname
		// log.Println(fname, fv.Interface(), gopp.Retn(fo.Tag.Get("gorm")), gopp.Retn(fo.Tag.Get("dbm")))
		fvx := fv.Interface()
		if !gopp.Empty(fvx) {
			// 特殊类型转换
			switch tv := fvx.(type) {
			case time.Time:
				fvx = tv.UnixMilli()
			}

			if gopp.StrsHaveNocase(keys, fname4db) {
				wheres_ = append(wheres_, fname4db+"=?")
				binds2_ = append(binds2_, fvx)
			}

		} else {
			if gopp.StrsHaveNocase(keys, fname4db) {
				gopp.Warn("Wow key empty", fname4db)
			}
		}

	}

	// todo where cond not only `and`
	wheres = strings.Join(wheres_, " and ")

	binds = binds2_

	// simple check
	gopp.ZeroPrint(len(binds2_), "empty wheres???", keys, vx)
	gopp.FalsePrint(len(wheres_) == len(binds), "cond binds not matchs", wheres_, binds)
	gopp.ZeroPrint(len(binds), "nothing binds???", vx)

	return
}

// todo how update mtimems only
// 由于结构体名字不规范，需要传递tbl名字
// 使用ez后续是没有考虑更复杂的sql的情况，后续可能有更完整的实现，像。。。
func DBAddrowez(db *sql.DB, tbl string, dup2up bool, vx any, keys ...string) (sql.Result, error) {
	sets, wheres, binds, err := UpsqlFields(true, vx)
	gopp.ErrPrint(err, tbl)

	var sql = fmt.Sprintf("insert into %v (%v) values (%v)", tbl, sets, wheres)
	// log.Println(sets, wheres, binds)
	res, err := DBExec(db, sql, binds...)
	if !IsDupkey(err) {
		gopp.ErrPrint(err, sql, len(binds), binds)
	}
	if IsDupkey(err) {
		if dup2up && len(keys) > 0 {
			// 现在的实现方式可以直接更新，会过滤掉新row vx中的空值
			// gopp.Trace("DB: Add => update", tbl, keys)
			res, err = DBUprowez(db, tbl, vx, keys...)
		}
	}
	return res, err
}

// 由于结构体名字不规范，需要传递tbl名字
func DBDelrowez(db *sql.DB, tbl string, vx any, keys ...string) (sql.Result, error) {
	_, wheres, binds, err := UpsqlFields(false, vx, keys...)
	gopp.ErrPrint(err, tbl)

	var sql = fmt.Sprintf("delete from %v where %v", tbl, wheres)
	// log.Println(sets, wheres, binds)
	res, err := DBExec(db, sql, binds...)
	return res, err
}

// todo 现在还不能update keys 中的字段
// 由于结构体名字不规范，需要传递tbl名字
func DBUprowez(db *sql.DB, tbl string, vx any, keys ...string) (sql.Result, error) {
	sets, wheres, binds, err := UpsqlFields(false, vx, keys...)
	gopp.ErrPrint(err, tbl)

	var sql = fmt.Sprintf("update %v set %v where %v", tbl, sets, wheres)
	// log.Println(sets, wheres, binds)
	res, err := DBExec(db, sql, binds...)
	return res, err
}

var prepedstmts sync.Map // string => sql.Stmt

// todo return affect rows, insertid
func DBExec(db *sql.DB, sqltpl string, binds ...any) (sql.Result, error) {
	// tx, err := db.Begin()
	// gopp.ErrPrint(err)

	var stmt *sql.Stmt
	if stmtx, ok := prepedstmts.Load(sqltpl); ok {
		stmt = stmtx.(*sql.Stmt)
	} else {
		stmt2, err := db.Prepare(sqltpl)
		gopp.ErrPrint(err, sqltpl)
		if err != nil {
			return nil, err
		}
		// defer stmt.Close()
		prepedstmts.Store(sqltpl, stmt2)
		stmt = stmt2
	}

	res, err := stmt.Exec(binds...)
	_ = res
	// gopp.ErrPrint(err, res == nil, sql)
	if err == nil {
		// tx.Commit()
	} else {
		// tx.Rollback()
	}

	return res, err
}
