package main

import (
	"goutils"
	_ "qqpaygc/routers"

	"github.com/astaxie/beego"
)

func main() {
	goutils.InitLogs()
	beego.Run()
}
