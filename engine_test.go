package tbox

import (
	"log"
	"testing"
)

func TestTableConnection(t *testing.T) {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/test?checkConnLiveness=false&loc=Local" +
		"&parseTime=true&readTimeout=5s&timeout=10s" +
		"&writeTimeout=5s&maxAllowedPacket=0&charset=utf8mb4"
	log.Println("dsn: ", dsn)
	enc := New(dsn, WithEnableJsonTag(), WithEnableTableNameFunc(), WithOutputCmd(), WithTagKey("gorm"))
	// enc := New(dsn, WithEnableJsonTag(), WithEnableTableNameFunc(), WithOutputCmd(), WithTagKey("xorm"))
	// enc := New(dsn, WithEnableJsonTag(), WithEnableTableNameFunc(), WithOutputCmd(), WithTagKey("db"))

	log.Println(enc.GetColumns("user"))
	log.Println(enc.camelCase(lintName("uID_info"))) // UIDInfo
	log.Println(enc.Run())
}
