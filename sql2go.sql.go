package tbox

import (
	"fmt"
)

// map for converting mysql type to golang types
var typeForMysqlToGo = map[string]string{
	"int":                "int64",
	"integer":            "int64",
	"tinyint":            "int64",
	"smallint":           "int64",
	"mediumint":          "int64",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int64",
	"smallint unsigned":  "int64",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func getType(mysqlType string) string {
	s, ok := typeForMysqlToGo[mysqlType]
	if ok {
		return s
	}

	panic(fmt.Errorf("sql type:%s not found", mysqlType))
}

// RegisterType register mysql to to golang type
func RegisterType(mysqlType string, goType string) {
	if _, ok := typeForMysqlToGo[mysqlType]; ok {
		panic(fmt.Sprintf("the mysql type:%s to go type:%s has exist", mysqlType, goType))
	}

	typeForMysqlToGo[mysqlType] = goType
}
