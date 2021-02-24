package auth

import (
	"github.com/astaxie/beego"
	"gocmdb/base/controllers/base"
	"gocmdb/models"
	"gocmdb/services"
	"html/template"
	"net/http"
	"strings"
)

//所有需要认证才能访问的基础控制器
type AuthorizationController struct {
	base.BaseController
	LoginUser *models.User
}

// 获取nav
func (c *AuthorizationController) getNav() string {
	controllerName, _ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}

// Prepare 用户认证检查
func (c *AuthorizationController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessionValue := c.GetSession(sessionKey)
	c.Data["loginUser"] = nil
	c.Data["nav"] = c.getNav()

	if sessionValue != nil {
		if pk, ok := sessionValue.(int); ok {
			if user := services.UserService.GetByPk(pk); user != nil {
				c.Data["loginUser"] = user
				c.LoginUser = user
				c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
				return
			}
		}
	}

	action := beego.AppConfig.DefaultString("auth::LoginAction", "AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)

}
