package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["test-golang/controllers:UrlConvertController"] = append(beego.GlobalControllerRouter["test-golang/controllers:UrlConvertController"],
		beego.ControllerComments{
			Method:           "ShortenUrl",
			Router:           `/shorten_url`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("params", param.InBody),
			),
			Params: nil})
	beego.GlobalControllerRouter["test-golang/controllers:UrlConvertController"] = append(beego.GlobalControllerRouter["test-golang/controllers:UrlConvertController"],
		beego.ControllerComments{
			Method:           "RedirectUrl",
			Router:           `/redirect_url`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("url"),
			),
			Params: nil})
	beego.GlobalControllerRouter["test-golang/controllers:UrlConvertController"] = append(beego.GlobalControllerRouter["test-golang/controllers:UrlConvertController"],
		beego.ControllerComments{
			Method:           "StatsUrl",
			Router:           `/stats_url`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("params", param.InBody),
			),
			Params: nil})

}
