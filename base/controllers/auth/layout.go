package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()

	// Prepare 登录验证
	token, e := c.ParseToken()
	if e != nil {
		panic(e)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	user := claims["username"].(string)
	fmt.Println(user)
	if !ok {
		c.Abort("permission")
		return
	}
}
