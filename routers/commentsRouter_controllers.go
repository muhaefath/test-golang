package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["test-golang/controllers:BookController"] = append(beego.GlobalControllerRouter["test-golang/controllers:BookController"],
		beego.ControllerComments{
			Method:           "GetListBook",
			Router:           `/get_list_book`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("genre"),
				param.New("offset"),
				param.New("limit"),
			),
			Params: nil})
	beego.GlobalControllerRouter["test-golang/controllers:BookController"] = append(beego.GlobalControllerRouter["test-golang/controllers:BookController"],
		beego.ControllerComments{
			Method:           "PickUpBook",
			Router:           `/pick_up_book`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("params", param.InBody),
			),
			Params: nil})
	beego.GlobalControllerRouter["test-golang/controllers:BookController"] = append(beego.GlobalControllerRouter["test-golang/controllers:BookController"],
		beego.ControllerComments{
			Method:           "ReturnBook",
			Router:           `/return_book`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("params", param.InBody),
			),
			Params: nil})

}
