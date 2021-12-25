package main

import (
	"flag"
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
	flag.StringVar(&dsn, "dsn", "", "dsn string")
	flag.StringVar(&pkgName, "p", "model", "pkg name")
	flag.StringVar(&pkgPath, "d", "./model", "pkg dir path")
	flag.StringVar(&tagKey, "tag", "db", "tag key")
	flag.StringVar(&tab, "t", "", "table,eg:-t=user;order")

	flag.BoolVar(&isOutputCmd, "v", false, "output cmd,eg:-v=true")
	flag.BoolVar(&ucFirstOnly, "u", true, "uc first only,eg:-u=true")
	flag.BoolVar(&enableTableNameFunc, "m", false, "add TableName func eg:-m=true")
	flag.BoolVar(&enableJsonTag, "j", false, "add json tag eg:-j=true")
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

	enc := tbox.New(dsn, options...)
	if tab != "" {
		tables := strings.Split(strings.TrimSuffix(tab, ";"), ";")
		enc.Run(tables...)
		return
	}

	enc.Run()
}
