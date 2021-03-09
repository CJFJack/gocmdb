package utils

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {
	beego.Info(password)
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash)
}

func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
