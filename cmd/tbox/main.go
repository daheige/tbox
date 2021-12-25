package main

import (
	"flag"
	"log"
	"strings"

	"github.com/daheige/tbox"
)

var (
	dsn                 string
	pkgName             string
	pkgPath             string
	isOutputCmd         bool
	tagKey              string
	ucFirstOnly         bool
	enableTableNameFunc bool
	enableJsonTag       bool
	tab                 string
)

func init() {
	flag.StringVar(&dsn, "dsn", "", `mysql dsn,eg:-dsn="root:root1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4"`)
	flag.StringVar(&pkgName, "p", "model", "pkg name,eg:-p=model")
	flag.StringVar(&pkgPath, "d", "./model", "pkg dir path,eg:-d=./model")
	flag.StringVar(&tagKey, "tag", "db", "tag key,eg:-tag=db")
	flag.StringVar(&tab, "t", "", "table,eg:-u=user;order")

	flag.BoolVar(&isOutputCmd, "v", false, "whether output cmd,eg:-v=true")
	flag.BoolVar(&ucFirstOnly, "u", true, "whether uc first only,eg:-u=true")
	flag.BoolVar(&enableTableNameFunc, "m", false, "whether add TableName func eg:-m=true")
	flag.BoolVar(&enableJsonTag, "j", false, "whether add json tag eg:-j=true")
	flag.Parse()
}

func main() {
	options := []tbox.Option{
		tbox.WithPkgName(pkgName),
		tbox.WithPkgPath(pkgPath),
	}

	if isOutputCmd {
		options = append(options, tbox.WithOutputCmd())
	}

	if ucFirstOnly {
		options = append(options, tbox.WithUcFirstOnly())
	}

	if enableTableNameFunc {
		options = append(options, tbox.WithEnableTableNameFunc())
	}

	if enableJsonTag {
		options = append(options, tbox.WithEnableJsonTag())
	}

	var err error
	enc := tbox.New(dsn, options...)
	if tab != "" {
		tables := strings.Split(strings.TrimSuffix(tab, ";"), ";")
		err = enc.Run(tables...)
	} else {
		err = enc.Run()
	}

	if err != nil {
		log.Fatalln("generating code error: ", err)
	}

	log.Println("generating code success")
}
