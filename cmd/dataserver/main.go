package main

import (
	_ "heroku-dataserver/cmd/dataserver/routers"

	"heroku-dataserver/cmd/dataserver/healthcheck"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	toolbox.AddHealthCheck("database", &healthcheck.CitycontentCheck{})
	//collector.Run()
	beego.Run()
}
