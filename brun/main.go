package main

import (
	"github.com/bairn/props/ini"
	"github.com/bairn/props/kvs"
	_ "github.com/bairn/account"
	"github.com/bairn/infra"
	"github.com/bairn/infra/base"
)

func main() {
	file := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()
}
