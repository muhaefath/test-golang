package main

import (
	_ "net/http/pprof"

	"github.com/astaxie/beego"
	beegoContext "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"

	_ "test-golang/models/registration"
	_ "test-golang/routers"
)

var originalRecoveryFunc func(*beegoContext.Context)

func init() {
	originalRecoveryFunc = beego.BConfig.RecoverFunc

}

func main() {
	// CORS: allow all origins in any mode
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"https://*.local"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.Run()
}
