package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"gocmdb/base/controllers/base"
	"gocmdb/base/errors"
	"gocmdb/models"
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

// 解析Token
func (c *AuthorizationController) ParseToken() (t *jwt.Token, e *errors.Errors) {
	errs := errors.New()
	authString := c.Ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		errs.Add("auth", fmt.Sprintf("AuthString invalid: %s", authString))
		return nil, errs
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(beego.AppConfig.DefaultString("JWTTokenKey", "CMDB")), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				errs.Add("auth", "errInputData")
				return nil, errs
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				errs.Add("auth", "token expired")
				return nil, errs
			} else {
				// Couldn't handle this token
				errs.Add("auth", "errInputData")
				return nil, errs
			}
		} else {
			// Couldn't handle this token
			errs.Add("auth", "errInputData")
			return nil, errs
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		errs.Add("auth", fmt.Sprintf("Token invalid: %s", tokenString))
		return nil, errs
	}
	beego.Debug("Token:", token)
	return token, nil
}
