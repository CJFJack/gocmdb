package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocmdb/base/controllers/base"
	"gocmdb/base/response"
)

type APIController struct {
	base.BaseController
}

func (c *APIController) Prepare() {
	c.EnableXSRF = false

	token := fmt.Sprintf("Token %s", beego.AppConfig.DefaultString("api::token", ""))
	headerToken := c.Ctx.Input.Header("Authorization")

	if token != headerToken {
		c.Data["json"] = response.UnAuthorization
	}
}
