package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"gocmdb/base/controllers/base"
	"gocmdb/config"
	"gocmdb/services"
	"time"
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
				result["data"].(map[string]string)["username"] = user.NickName
				result["data"].(map[string]string)["uuid"] = fmt.Sprintf("%s", uuid.NewV4())
				result["msg"] = "登录成功"
				// 设置token
				claims := make(jwt.MapClaims)
				claims["username"] = user.Name
				claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //20天有效期，过期需要重新登录获取token
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				// 使用自定义字符串加密 and get the complete encoded token as a string
				tokenString, _ := token.SignedString([]byte(beego.AppConfig.DefaultString("JWTTokenKey", "CMDB")))
				result["data"].(map[string]string)["token"] = tokenString

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

	} else {
		result["code"] = 500
		result["msg"] = "请求方法不正确"
	}
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
	c.ServeJSON()
}
