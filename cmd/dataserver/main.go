package main

import (
	_ "dataserver/routers"

	"dataserver/healthcheck"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	toolbox.AddHealthCheck("database", &healthcheck.CitycontentCheck{})
	toolbox.AddHealthCheck("citycontent", &healthcheck.DatabaseCheck{})
	//collector.Run()
	beego.Run()
}
