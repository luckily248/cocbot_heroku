package controllers

import (
	. "dataserver/models"
	"fmt"

	"github.com/astaxie/beego"
)

// get city citycontent
type CityContentController struct {
	beego.Controller
}

func (c *CityContentController) URLMapping() {
	c.Mapping("GetOneCity", c.GetOneCity)
	c.Mapping("GetOneCityContent", c.GetOneCityContent)
}

// @Title Get
// @Description get citycontent
// @Success 200 {object} models.citycontent
// @Failure 404 not found
// @router /citycontent [get]
func (c *CityContentController) GetOneCityContent() {

	citycontent, err := GetCityContent()
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = err.Error()
		fmt.Printf("err:%s\n", err.Error())
	} else {
		c.Data["json"] = citycontent
		fmt.Printf("citycontent:%v\n", citycontent)
	}
	c.ServeJson()

}

// @Title Get
// @Description get City by name
// @Param	name		path 	string	true	"cityname"
// @Success 200 {object} models.City
// @Failure 403 :name is empty
// @Failure 404 not found
// @router /:name [get]
func (c *CityContentController) GetOneCity() {
	cityname := c.Ctx.Input.Params[":name"]
	if cityname == "" {
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = string("is empty")
	} else {
		city, err := GetOneCityid(cityname)
		if err != nil {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = err.Error()
			fmt.Printf("err:%s\n", err.Error())
		} else {
			c.Data["json"] = city
			fmt.Printf("city:%v\n", city)
		}
	}
	c.ServeJson()

}
