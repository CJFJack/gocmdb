package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"gocmdb/controllers"
	v1 "gocmdb/controllers/api/v1"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "X-XSRFToken", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Access-Control-Max-Age"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Access-Control-Max-Age"},
		AllowCredentials: true,
	}))
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UsersController{})
	beego.AutoRouter(&controllers.CloudPlatformController{})
	beego.AutoRouter(&controllers.VirtualMachineController{})

	// Prometheus
	beego.AutoRouter(&controllers.NodeController{})
	beego.AutoRouter(&controllers.JobController{})
	beego.AutoRouter(&controllers.TargetController{})

	// v1
	v1Prometheus := beego.NewNamespace("v1", beego.NSAutoRouter(&v1.PrometheusController{}))
	beego.AddNamespace(v1Prometheus)
}
