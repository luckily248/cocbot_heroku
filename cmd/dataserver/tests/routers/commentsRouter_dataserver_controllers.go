package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
