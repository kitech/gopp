package sqlu

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/kitech/gopp"

	spjson "github.com/bitly/go-simplejson"
)

const (
	SQLITE_META_DATABASE   = ""
	SQLITE_META_TABLE      = "sqlite_master"
	SQLITE_TABLE_WRAP_CHAR = "[" // ]
	SQLITE_FIELD_WRAP_CHAR = "\n"
	SQLITE_VALUE_WRAP_CHAR = ""

	MYSQL_META_DATABASE   = ""
	MYSQL_META_TABLE      = ""
	MYSQL_TABLE_WRAP_CHAR = "`"
	MYSQL_FIELD_WRAP_CHAR = "`"
	MYSQL_VALUE_WRAP_CHAR = "'" // "

	PGSQL_META_DATABASE   = ""
	PGSQL_META_TABLE      = ""
	PGSQL_TABLE_WRAP_CHAR = ""
	PGSQL_FIELD_WRAP_CHAR = ""
	PGSQL_VALUE_WRAP_CHAR = "'"
)

var lastsqlitefile string = os.Getenv("HOME") + "/fedyui.db3"
var lastsqlitecon *sql.DB

func ValuepByColumnType(cty *sql.ColumnType) reflect.Value {
	rtype := cty.ScanType()
	// log.Println(rtype, rtype.String())
	rv := reflect.New(rtype)

	return rv
}

type SqlRows []map[string]gopp.Any
type SqlRow map[string]gopp.Any

// func (me SqlValue) Any() any { return me }
func (me SqlRows) Rowcnt() int {
	return len(me)
}
func (me SqlRows) Colcnt() int {
	if len(me) == 0 {
		return 0
	}
	return len(me[0])
}
func (me SqlRows) Names() []string {
	if len(me) == 0 {
		return nil
	}
	var res []string
	for k, _ := range me[0] {
		res = append(res, k)
	}
	return res
}

func (me SqlRows) Rowat(i int) SqlRow {
	return me[i]
}
func (me SqlRows) Colat(i int, name string) gopp.Any {
	return me[i][name]
}
func (me SqlRows) Colat2(i int, name string) any {
	return me[i][name].I
}
func (me SqlRow) Colat(name string) gopp.Any {
	return me[name]
}
func (me SqlRow) Colat2(name string) any {
	return me[name].I
}

// ///
func Rows2Spjson(rows *sql.Rows) *spjson.Json {
	res := Rows2Table(rows)
	log.Println(res)
	bdata, err := json.Marshal(res)
	gopp.ErrPrint(err)
	jso, err := spjson.NewJson(bdata)
	log.Println(jso)
	return jso
}

// Rows2Assoc
func Rows2Table(rows *sql.Rows) SqlRows {
	var retrows = SqlRows{}

	coltys, err := rows.ColumnTypes()
	colnames, err := rows.Columns()
	for i := 0; rows.Next(); i++ {
		// log.Println(i)
		var valvars []any

		retrow := map[string]gopp.Any{}

		for j := 0; j < len(coltys); j++ {
			colty := coltys[j]
			// colname := colnames[j]
			rv := ValuepByColumnType(colty)
			// log.Println(i, j, colname, rv)
			valvars = append(valvars, rv.Interface())
		}

		err = rows.Scan(valvars...)
		gopp.ErrPrint(err, i, valvars)
		// log.Println(i, valvars, reflect.TypeOf(valvars[0]))

		for j := 0; j < len(coltys); j++ {
			vv := valvars[j]
			switch v := vv.(type) {
			case *sql.NullString:
				// log.Println(i, j, v.String)
				retrow[colnames[j]] = gopp.ToAny(v.String)
			case *sql.NullFloat64:
				retrow[colnames[j]] = gopp.ToAny(v.Float64)
			case *sql.NullByte:
				retrow[colnames[j]] = gopp.ToAny(v.Byte)
			case *sql.NullBool:
				retrow[colnames[j]] = gopp.ToAny(v.Bool)
			case *sql.NullTime:
				retrow[colnames[j]] = gopp.ToAny(v.Time)
			case *sql.NullInt64:
				retrow[colnames[j]] = gopp.ToAny(v.Int64)
			case *sql.NullInt32:
				retrow[colnames[j]] = gopp.ToAny(v.Int32)
			case *sql.NullInt16:
				retrow[colnames[j]] = gopp.ToAny(v.Int16)
			default:
				log.Println("wtelse", i, j, reflect.TypeOf(vv))
			}
		}
		// log.Println(i, retrow)
		retrows = append(retrows, retrow)
	}
	// log.Println(retrows)
	return retrows
}
func Rows2Each(rows *sql.Rows, f func(rc int, row map[string]any)) {
	var retrows = SqlRows{}

	coltys, err := rows.ColumnTypes()
	colnames, err := rows.Columns()
	for i := 0; rows.Next(); i++ {
		// log.Println(i)
		var valvars []any

		retrow := map[string]gopp.Any{}
		retarr := []any{}

		for j := 0; j < len(coltys); j++ {
			colty := coltys[j]
			// colname := colnames[j]
			rv := ValuepByColumnType(colty)
			// log.Println(i, j, colname, rv)
			valvars = append(valvars, rv.Interface())
		}

		err = rows.Scan(valvars...)
		gopp.ErrPrint(err, i, valvars)
		// log.Println(i, valvars, reflect.TypeOf(valvars[0]))

		for j := 0; j < len(coltys); j++ {
			vv := valvars[j]
			switch v := vv.(type) {

			case *sql.NullString:
				// log.Println(i, j, v.String)
				retrow[colnames[j]] = gopp.ToAny(v.String)
				retarr = append(retarr, v.String)
			case *sql.NullFloat64:
				retrow[colnames[j]] = gopp.ToAny(v.Float64)
				retarr = append(retarr, v.Float64)
			case *sql.NullByte:
				retrow[colnames[j]] = gopp.ToAny(v.Byte)
				retarr = append(retarr, v.Byte)
			case *sql.NullBool:
				retrow[colnames[j]] = gopp.ToAny(v.Bool)
				retarr = append(retarr, v.Bool)
			case *sql.NullTime:
				retrow[colnames[j]] = gopp.ToAny(v.Time)
				retarr = append(retarr, v.Time)
			case *sql.NullInt64:
				retrow[colnames[j]] = gopp.ToAny(v.Int64)
				retarr = append(retarr, v.Int64)
			case *sql.NullInt32:
				retrow[colnames[j]] = gopp.ToAny(v.Int32)
				retarr = append(retarr, v.Int32)
			case *sql.NullInt16:
				retrow[colnames[j]] = gopp.ToAny(v.Int16)
				retarr = append(retarr, v.Int16)
			default:

				log.Println("wtelse", i, j, reflect.TypeOf(vv))
			}
		}
		// log.Println(i, retrow)
		retrows = append(retrows, retrow)
		var rowany = map[string]any{}
		for j, vx := range retarr {
			rowany[colnames[j]] = vx
		}
		f(i, rowany)
	}
	// log.Println(retrows)

}

const (
	DBTY_NONE   = 0
	DBTY_MYSQL  = 1
	DBTY_PGSQL  = 2
	DBTY_SQLITE = 3
	DBTY_ORACLE = 4
)

type SqlBuilder struct {
	strings.Builder

	dbtype int
	more   bool
}

type MysqlBuilder struct {
	SqlBuilder
}
type PgsqlBuilder struct {
	SqlBuilder
}
type SqliteBuilder struct {
	SqlBuilder
}

func NewSqlBuilder() *SqlBuilder {
	sb := strings.Builder{}
	return &SqlBuilder{Builder: sb}
}
func (sb *SqlBuilder) Tosql() string {
	sb.Checksql()
	return sb.Builder.String()
}
func (sb *SqlBuilder) Checksql() error {
	var err error
	if sb.more == true {
		err = fmt.Errorf("Maybe not normal terminated sql: %s", sb.Builder.String())
		log.Println(err)
	}
	return err
}

//	func (sb *SqlBuilder) Reset() *SqlBuilder {
//		sb.Builder.Reset()
//		return sb
//	}

func (sb *SqlBuilder) Embed(sb2 *SqlBuilder) *SqlBuilder {
	sb.WriteString("(")
	sb.WriteString(sb2.String())
	sb.WriteString(")")
	return sb
}

func (sb *SqlBuilder) Sp() *SqlBuilder {
	sb.WriteString(" ")
	return sb
}
func (sb *SqlBuilder) Fh() *SqlBuilder {
	sb.WriteString(";")
	return sb
}
func (sb *SqlBuilder) Star() *SqlBuilder {
	sb.WriteString("*")
	return sb
}
func (sb *SqlBuilder) Dh() *SqlBuilder {
	sb.WriteString("*")
	return sb
}
func (sb *SqlBuilder) Eq() *SqlBuilder {
	sb.WriteString("=")
	return sb
}
func (sb *SqlBuilder) Neq() *SqlBuilder {
	sb.WriteString("=")
	return sb
}

// func (sb *SqlBuilder) Table(t string) *SqlBuilder {
// 	sb.WriteString(t)
// 	sb.Sp()
// 	return sb
// }

func (sb *SqlBuilder) From(ts ...string) *SqlBuilder {
	sb.WriteString("FROM ")
	sb.writelistident(ts...)
	sb.Sp()
	return sb
}

func (sb *SqlBuilder) Select0() *SqlBuilder {
	sb.WriteString("SELECT /*Select0*/")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Select(names ...string) *SqlBuilder {
	if len(names) == 0 {
		return sb.SelectAll()
	}
	sb.WriteString("SELECT /*Select*/")
	sb.Sp()
	sb.writelistident(names...)
	sb.Sp()
	return sb
}

func (sb *SqlBuilder) SelectAll() *SqlBuilder {
	sb.WriteString("SELECT /*SelectAll*/")
	sb.Sp().Star().Sp()
	return sb
}
func (sb *SqlBuilder) Delete() *SqlBuilder {
	sb.WriteString("DELETE")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Count() *SqlBuilder {
	sb.WriteString("COUNT(*)")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Countx(s string) *SqlBuilder {
	sb.WriteString(s)
	sb.Sp()
	return sb
}

func (sb *SqlBuilder) Where0() *SqlBuilder {
	sb.WriteString("WHERE /*Where0*/")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Where(lft string, op string, rgt any) *SqlBuilder {
	sb.WriteString("WHERE")
	sb.Sp()
	sb.Cond(lft, op, rgt)
	return sb
}
func (sb *SqlBuilder) Wherex(cond string) *SqlBuilder {
	sb.WriteString("WHERE /*Wherex*/")
	sb.Sp()
	sb.WriteString(cond)
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Cond(lft string, op string, rgt any) *SqlBuilder {
	sb.WriteString(lft)
	sb.Sp().Strf(op).Sp()
	if reflect.TypeOf(rgt).Kind() == reflect.String {
		sb.WriteString(fmt.Sprintf("'%v'", rgt))
	} else {
		sb.WriteString(fmt.Sprintf("%v", rgt))
	}
	sb.Sp()

	return sb
}
func (sb *SqlBuilder) Condx(cond string) *SqlBuilder {
	sb.WriteString(cond)
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) In(lft string, ins ...any) *SqlBuilder {
	sb.WriteString(lft)
	sb.Sp()
	sb.WriteString("IN")
	sb.WriteString("(")
	sb.writelistvalue(ins...)
	sb.WriteString(")")
	return sb
}
func (sb *SqlBuilder) NotIn(lft string, ins ...any) *SqlBuilder {
	sb.WriteString(lft)
	sb.Sp()
	sb.WriteString("NOT IN")
	sb.WriteString("(")
	sb.writelistvalue(ins...)
	sb.WriteString(")")
	return sb
}
func (sb *SqlBuilder) writelistvalue(vals ...any) *SqlBuilder {
	for n, v := range vals {
		if reflect.TypeOf(v).Kind() == reflect.String {
			sb.Strf("'%v'", v)
		} else {
			sb.Strf("%v", v)
		}
		sb.WriteString(gopp.IfElseStr(n == len(vals)-1, "", ","))
	}
	return sb
}
func (sb *SqlBuilder) writelistident(idts ...string) *SqlBuilder {
	for n, v := range idts {
		sb.Strf("%v", v)
		sb.WriteString(gopp.IfElseStr(n == len(idts)-1, "", ","))
	}
	return sb
}
func (sb *SqlBuilder) And(lft string, op string, rgt any) *SqlBuilder {
	sb.WriteString("AND")
	sb.Sp()
	sb.Cond(lft, op, rgt)
	sb.more = true
	return sb
}
func (sb *SqlBuilder) Andx(cond string) *SqlBuilder {
	sb.WriteString("AND")
	sb.WriteString(cond)
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Or(lft string, op string, rgt any) *SqlBuilder {
	sb.WriteString("OR")
	sb.Sp()
	sb.Cond(lft, op, rgt)
	return sb
}
func (sb *SqlBuilder) Orx(cond string) *SqlBuilder {
	sb.WriteString("OR")
	sb.Sp()
	sb.WriteString(cond)
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Like(lft string, cond string) *SqlBuilder {
	sb.WriteString(lft)
	sb.Sp()
	sb.WriteString("LIKE")
	sb.Sp()
	sb.Strf("'%%%s%%'", cond)
	sb.Sp()
	return sb
}

// func (sb *SqlBuilder) Likex(cond string) *SqlBuilder {
// 	// sb.WriteString("LIKE")
// 	sb.Sp()
// 	sb.WriteString(cond)
// 	sb.Sp()
// 	return sb
// }

func (sb *SqlBuilder) Strf(f string, args ...any) *SqlBuilder {
	str := fmt.Sprintf(f, args...)
	sb.WriteString(str)
	return sb
}
func (sb *SqlBuilder) Limit(n int, off ...int) *SqlBuilder {
	sb.WriteString("LIMIT")
	sb.Sp().Strf("%d", n)
	if len(off) > 0 {
		sb.Strf(",%d", off[0])
	}
	sb.Sp()
	return sb
}

func (sb *SqlBuilder) Insert(t string, fields ...string) *SqlBuilder {
	sb.WriteString("INSERT")
	sb.Sp()
	sb.WriteString(t)
	sb.Sp()

	return sb
}
func (sb *SqlBuilder) Values(values ...any) *SqlBuilder {
	sb.WriteString("VALUES (")
	// sb.WriteString("'")
	sb.writelistvalue(values...)
	// sb.WriteString("'")
	sb.WriteString(")")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Replace(fields ...string) *SqlBuilder {
	sb.WriteString("INSERT")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Update(t string) *SqlBuilder {
	sb.WriteString("UPDATE")
	sb.Sp()
	sb.WriteString(t)
	sb.Sp()
	sb.WriteString("SET")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Set(key string, value any) *SqlBuilder {
	sb.WriteString(key)
	sb.Eq()
	sb.writelistvalue(value)
	sb.Sp()

	return sb
}

func (sb *SqlBuilder) Incr(field string, step ...int) *SqlBuilder {
	sb.WriteString(field)
	sb.Eq()
	sb.WriteString(field)
	if len(step) == 0 {
		sb.Strf("+1")
	} else {
		sb.Strf("+%d", step[0])
	}
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Order(fields ...string) *SqlBuilder {
	sb.WriteString("ORDER BY")
	sb.Sp()
	for n, v := range fields {
		sb.WriteString(v)
		sb.WriteString(gopp.IfElseStr(n == len(fields)-1, "", ","))
	}
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Desc() *SqlBuilder {
	sb.WriteString("DESC")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Asc() *SqlBuilder {
	sb.WriteString("ASC")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Group(names ...string) *SqlBuilder {
	sb.WriteString("GROUP BY")
	sb.Sp()
	sb.writelistident(names...)
	return sb
}

// todo
func (sb *SqlBuilder) Join(names ...string) *SqlBuilder {
	sb.WriteString("JOIN")
	sb.Sp()
	sb.writelistident(names...)
	return sb
}

func (sb *SqlBuilder) Create(names ...string) *SqlBuilder {
	sb.WriteString("CREATE")
	sb.Sp()
	sb.writelistident(names...)
	return sb
}
func (sb *SqlBuilder) CreateTable(t string, names ...string) *SqlBuilder {
	sb.WriteString("CREATE")
	sb.Sp()
	sb.writelistident(names...)
	return sb
}
func (sb *SqlBuilder) EndTable() *SqlBuilder {
	sb.WriteString(")")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) Column() *SqlBuilder {
	sb.WriteString(")")
	sb.Sp()
	return sb
}

func (sb *SqlBuilder) CreateIndex(t string, names ...string) *SqlBuilder {
	sb.WriteString("CREATE INDEX ")
	sb.WriteString(strings.Join(names, "_") + "_idx")
	sb.Sp()
	sb.Strf("ON %s(", t)
	sb.Sp()
	sb.writelistident(names...)
	sb.WriteString(")")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) CreateUniqkey(t string, names ...string) *SqlBuilder {
	sb.WriteString("CREATE UNIQUE KEY ")
	sb.WriteString(strings.Join(names, "_") + "_idx")
	sb.Sp()
	sb.Strf("ON %s(", t)
	sb.Sp()
	sb.writelistident(names...)
	sb.WriteString(")")
	sb.Sp()
	return sb
}
func (sb *SqlBuilder) CreateFulltext(names ...string) *SqlBuilder {
	sb.WriteString("CREATE")
	sb.Sp()
	sb.writelistident(names...)
	return sb
}

const (
	Opgt      = ">"
	Oplt      = "<"
	Opgteq    = ">="
	Oplteq    = "<="
	Opne      = "!="
	Opeq      = "="
	OpNotnull = "IS NOT NULL"
	OpNull    = "IS NULL"
)
