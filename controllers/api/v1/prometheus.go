package v1

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocmdb/base/controllers/auth"
	"gocmdb/base/response"
	"gocmdb/models"
	"gocmdb/services"
)

type PrometheusController struct {
	auth.APIController
}

func (c *PrometheusController) Register() {
	c.Data["json"] = response.Ok
	model := models.NewNode()
	if err := c.ParseJson(model); err == nil {
		beego.Info(fmt.Sprintf("%#v", model))
		if err = services.NodeService.Register(model); err != nil {
			c.Data["json"] = response.BadRequest
		}
	} else {
		c.Data["json"] = response.BadRequest
	}
}

func (c *PrometheusController) Config() {
	/*
		[
			{
				"key": "",
				"target: [
					{"addr": ""}, {"addr": ""}
				]
			}
		]
	*/
	uuid := c.GetString("uuid")
	rt := services.JobService.GetByUUID(uuid)
	c.Data["json"] = response.NewJsonResponse(200, "ok", rt)

}
