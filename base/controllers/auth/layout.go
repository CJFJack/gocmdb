package auth

import "github.com/astaxie/beego"

type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()

	c.Layout = "base/layouts/layout.html"
	c.Data["title"] = beego.AppConfig.DefaultString("AppName", "CMDB")

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SectionStyle"] = ""
	c.LayoutSections["SectionScript"] = ""
}
