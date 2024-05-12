package sqlu

import (
	"log"
	"testing"
)

func Test1(t *testing.T) {
	sb := NewSqlBuilder()
	sb.SelectAll().From("sqlite_master")
	sb.Where("id", Opeq, 12345).And("ctime", Oplteq, "jkaowefweaf")
	sb.Order("id", "name", "ctime").Desc()
	sb.Limit(3, 2)
	log.Println(sb.String())
}
func Test2(t *testing.T) {
	sb := NewSqlBuilder()
	sb.SelectAll().From("sqlite_master")
	sb.Where("id", Opeq, 12345).And("ctime", Oplteq, "jkaowefweaf")
	sb.Order("id", "name", "ctime").Desc()
	sb.Limit(3, 2)
	log.Println(sb.String())
}
