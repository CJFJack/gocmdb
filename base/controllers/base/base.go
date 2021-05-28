package base

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

// 所有业务控制器基础控制器
type BaseController struct {
	beego.Controller
}

func (c *BaseController) ParseJson(jsonData interface{}) error {
	c.Ctx.Input.CopyBody(10 * 1024 * 1024)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &jsonData)
	if err != nil {
		return err
	}
	err = c.ParseForm(jsonData)
	return err
}

func (c *BaseController) ParsePostForm(form interface{}) (map[string]interface{}, error) {
	c.Ctx.Input.CopyBody(10 * 1024 * 1024)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	return form.(map[string]interface{}), err
}

func (c *BaseController) Render() error {
	c.ServeJSON()
	return nil
}
