package main

import (
	"github.com/astaxie/beego"
	_ "gocmdb/routers"
)

func main() {
	beego.Run(":9801")
}
