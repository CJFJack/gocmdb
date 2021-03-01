package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"gocmdb/base/controllers/base"
	"gocmdb/config"
	"gocmdb/forms"
	"gocmdb/services"
	"time"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Prepare() {
	c.EnableXSRF = false
}

func (c *AuthController) Login() {
	// post请求，数据验证
	// 验证成功
	// 验证失败

	if c.Ctx.Request.Method == "POST" {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
			"data": map[string]string{},
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		config.Cache.Incr("login")
		form := forms.LoginForm{}
		if rawData, err := c.ParseJson(form); err == nil {
			user := services.UserService.GetByName(rawData["username"].(string))
			if user == nil {
				// 用户不存在
				result["code"] = 500
				result["msg"] = "用户不存在"
				logs.Error(fmt.Sprintf("用户不存在： %s", rawData["username"].(string)))
				return
			} else if user.ValidPassword(rawData["password"].(string)) {
				logs.Info(fmt.Sprintf("用户认证成功： %s", rawData["username"].(string)))
				// 用户密码正确
				result["data"].(map[string]string)["username"] = user.NickName
				result["data"].(map[string]string)["uuid"] = fmt.Sprintf("%s", uuid.NewV4())
				result["msg"] = "登录成功"
				// 设置token
				claims := make(jwt.MapClaims)
				claims["username"] = user.NickName
				maxAge, _ := time.ParseDuration(beego.AppConfig.DefaultString("JWTTokenMaxAge", "24h"))
				claims["exp"] = time.Now().Add(time.Hour * maxAge).Unix() //1天有效期，过期需要重新登录获取token
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				// 使用自定义字符串加密 and get the complete encoded token as a string
				tokenString, _ := token.SignedString([]byte(beego.AppConfig.DefaultString("JWTTokenKey", "CMDB")))
				result["data"].(map[string]string)["token"] = tokenString

			} else {
				// 用户密码不正确
				result["code"] = 500
				result["msg"] = "用户名或密码错误"
				logs.Error(fmt.Sprintf("用户名或密码错误： %s", rawData["username"].(string)))
				return
			}
		} else {
			result["code"] = 500
			result["msg"] = "用户名或密码错误"
			return
		}

	}

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
