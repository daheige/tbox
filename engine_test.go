package tbox

import (
	"log"
	"testing"
)

func TestTableConnection(t *testing.T) {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/test?checkConnLiveness=false&loc=Local" +
		"&parseTime=true&readTimeout=5s&timeout=10s" +
		"&writeTimeout=5s&maxAllowedPacket=0&charset=utf8mb4"
	enc := New(dsn, WithUcFirstOnly(), WithEnableJsonTag(), WithEnableTableNameFunc())

	log.Println(enc.GetColumns("user"))
	log.Println(enc.camelCase("uID_info"))
	log.Println(enc.Run())
}
