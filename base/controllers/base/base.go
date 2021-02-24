package base

import "github.com/astaxie/beego"

// 所有业务控制器基础控制器
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
}
