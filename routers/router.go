// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

/*beego有四种路由：固定路由、正则路由、自动路由和注解路由
固定路由为最简单的路由方式，例子如下：
beego.Router("/", &controllers.MainController{})，两个参数，一个路由，一个对应的方法
正则路由可通过正则表达式定义需要匹配比较复杂的路由，例子如下：
beego.Router(“/api/:id([0-9]+)“, &controllers.RController{})
自定义正则匹配 //匹配 /api/123 :id = 123
自动路由需要把路由的controllers注册到自动路由中，例子如下：
beego.AutoRouter(&controllers.ObjectController{})
beego便会自动获得controllers中的方法并生成路由，例子如下：
/object/login   调用ObjectController中的Login方法
/object/logout  调用ObjectController中的Logout方法
注解路由间/controllers/agency.go的注解


*/
import (
	"apitest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/object",
		// 	beego.NSInclude(
		// 		&controllers.ObjectController{},
		// 	),
		// ),
		// beego.NSNamespace("/user",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
		beego.NSNamespace("/agency",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
