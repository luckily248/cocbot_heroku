package main

import (
	"cocbot_heroku/cmd/dataserver/healthcheck"

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
