package controllers

import (
	"dataserver/models"
	"fmt"

	"github.com/astaxie/beego"
)

// get weather information
type ShowWeatherController struct {
	beego.Controller
}

func (c *ShowWeatherController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Get
// @Description get ShowWeatherModel by id
// @Param	id		path 	string	true		"The key for cityid"
// @Success 200 {object} models.ShowWeatherModel
// @Failure 404  not found
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ShowWeatherController) GetOne() {

	cityid := c.Ctx.Input.Params[":id"]
	//fmt.Printf("cityid:%s\n", cityid)
	if cityid == "" {
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = string("is empty")
	} else {
		ob, err := models.GetOneData(cityid)
		if err != nil {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = err.Error()
			fmt.Printf("err:%s\n", err.Error())
		} else {
			c.Data["json"] = ob
			fmt.Printf("ob:%v\n", ob)
		}
	}
	c.ServeJson()
}

// @Title Get All (NOT USE)
// @Description get ShowWeatherModel (not use)
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ShowWeatherModel
// @Failure 403
// @router / [get]
func (c *ShowWeatherController) GetAll() {

}
