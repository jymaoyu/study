package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apitest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apitest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apitest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apitest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apitest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["apitest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetbyId",
			Router: `/:Ano`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Ano`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:UserController"] = append(beego.GlobalControllerRouter["apitest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:UserController"] = append(beego.GlobalControllerRouter["apitest/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:UserController"] = append(beego.GlobalControllerRouter["apitest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:UserController"] = append(beego.GlobalControllerRouter["apitest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:UserController"] = append(beego.GlobalControllerRouter["apitest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:UserController"] = append(beego.GlobalControllerRouter["apitest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:UserController"] = append(beego.GlobalControllerRouter["apitest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
