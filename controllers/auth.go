package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/satori/go.uuid"
	"gocmdb/base/controllers/base"
	"gocmdb/config"
	"gocmdb/services"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Prepare() {
	c.EnableXSRF = false
}

func (c *AuthController) ParseJson() (map[string]interface{}, error) {
	var m map[string]interface{}
	c.Ctx.Input.CopyBody(10 * 1024 * 1024)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	return m, err
}

func (c *AuthController) Login() {
	// post请求，数据验证
	// 验证成功
	// 验证失败

	result := map[string]interface{}{
		"code":     0,
		"msg":      "ok",
		"username": "",
		"data":     map[string]string{},
	}
	if c.Ctx.Request.Method == "POST" {
		config.Cache.Incr("login")

		if rawData, err := c.ParseJson(); err == nil {
			user := services.UserService.GetByName(rawData["username"].(string))
			if user == nil {
				// 用户不存在
				result["code"] = 500
				result["msg"] = "用户不存在"
				logs.Error(fmt.Sprintf("用户不存在： %s", rawData["username"].(string)))
			} else if user.ValidPassword(rawData["password"].(string)) {
				logs.Info(fmt.Sprintf("用户认证成功： %s", rawData["username"].(string)))
				// 用户密码正确
				// 设置session
				sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
				c.SetSession(sessionKey, user.ID)
				result["data"].(map[string]string)["username"] = user.NickName
				result["data"].(map[string]string)["uuid"] = fmt.Sprintf("%s", uuid.NewV4())
				result["msg"] = "登录成功"

			} else {
				// 用户密码不正确
				result["code"] = 500
				result["msg"] = "用户名或密码错误"
				logs.Error(fmt.Sprintf("用户名或密码错误： %s", rawData["username"].(string)))
			}
		} else {
			result["code"] = 500
			result["msg"] = "用户名或密码错误"
		}
	}

	result["data"].(map[string]string)["token"] = c.XSRFToken()
	result["data"].(map[string]string)["uuid"] = c.XSRFToken()

	c.Data["json"] = result
	c.ServeJSON()
}

// 退出登录
func (c *AuthController) Logout() {
	c.DestroySession()
	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "注销成功",
	}
}
