package controllers

/*用户将会把请求通过路由发送给controller，而controller则会调用对应的方法*/

import (
	"apitest/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

/* 声明一个Controller并内嵌了beego.Controller，
这样这个Controller便自动拥有了所有 beego.Controller 的方法
类似于继承*/
type TestController struct {
	beego.Controller
}

/*这里采用的是注解路由的方法，这样便无需在router中注册路由了。
主需要在方法上加上// @router*便可，有2个参数，使用空格分隔，各参数的含义分别为：
请求的路由地址，支持正则路由和自定义路由、支持的请求方法，需要放入[]中，多个方法用“,”分割
注解路由生成的路由会放到 “/routers/commentsRouter.go” 文件中。


同时，在注解路由的上面可进行注释，用于api自动化文档，各参数含义如下：
@Title api的名称
@Description api的详细描述
@Param 需要传递给服务器的参数，有5个参数，使用空格分隔，各参数的含义分别为：
参数名、参数类型、参数类型、是否必须、注释
@Success 成功返回给客户端的信息，有3个参数，使用空格分隔，各参数的含义分别为：
status code、返回的类型（必须用{}包含）、返回的对象或字符串信息
@Failure 失败返回的信息，有2个参数，使用空格分隔，各参数的含义分别为：
status code、错误信息
注释完后需要在“/routers/router.go” 文件中写入解析。


beego有多种方法可以获得用户传递的数据，如：
GetString(key string) string
GetStrings(key string) []string
GetInt(key string) (int64, error)
GetBool(key string) (bool, error)
GetFloat(key string) (float64, error)
当需要获取Request Body里的JSON或XML的数据时，需要在配置文件里设置 copyrequestbody = true
然后在Controller中使用json.Unmarshal方法获取数据


beego里数据有好几种输出：
1.直接输出字符串，用法：
beego.Controller.Ctx.WriteString()
2.模板数据输出，用法：
beego.Controller.Date["名字"]=数据
beego.Controller.TplName=模板文件
3.json格式数据输出，用法：
beego.Controller.Date["json"]=数据
beego.ControllerServeJSON()
4.xml格式数据输出，用法：
beego.Controller.Date["xml"]=数据
beego.ControllerServeXML()
5.jsonp调用，用法：
beego.Controller.Date["jsonp"]=数据
beego.ControllerServeJSONP()*/

// @Title 获得所有agency
// @Description 返回所有的agency数据
// @Success 200 {object} models.Agency
// @router / [get]
func (T *TestController) GetAll() {
	agencys, err := models.GetAllAgency()
	if err != nil {
		T.Ctx.WriteString(err.Error())
	}
	T.Data["json"] = agencys
	T.ServeJSON()
}

// @Title 获得一个agency
// @Description 返回某agency数据
// @Param	Ano		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Agency
// @router 	/:Ano [get]
func (T *TestController) GetbyId() {
	Ano := T.GetString(":Ano")
	agency, err := models.GetAgencybyId(Ano)
	if err != nil {
		T.Ctx.WriteString(err.Error())
	}
	T.Data["json"] = agency
	T.ServeJSON()
}

// @Title  添加agency
// @Description  添加agency的描述
// @Param      body          body   models.Agency true          "body for user content"
// @Success 200 {int} models.Agency.Ano
// @Failure 403 body is empty
// @router /add [post]
func (T *TestController) Add() {
	var a models.Agency
	json.Unmarshal(T.Ctx.Input.RequestBody, &a)
	result, err := models.AddAgency(&a)
	if err != nil {
		T.Ctx.WriteString(err.Error())
	}
	if result != true {
		T.Ctx.WriteString("add fail")
	}
	T.Data["json"] = "add success"
	T.ServeJSON()
}

// @Title 修改agency
// @Description 修改agency的内容
// @Param      body          body   models.Agency true          "body for user content"
// @Success 200 {int} models.Agency
// @Failure 403 body is empty
// @router /update [post]
func (T *TestController) Update() {
	var a models.Agency
	json.Unmarshal(T.Ctx.Input.RequestBody, &a)
	result, err := models.UpdateAgency(&a)
	if err != nil {
		T.Ctx.WriteString(err.Error())
	}
	if result != true {
		T.Ctx.WriteString("update fail")
	}
	T.Data["json"] = "update success"
	T.ServeJSON()
}

// @Title 删除一个agency
// @Description 删除agency数据
// @Param      Ano           path string    true          "The key for staticblock"
// @Success 200 {object} models.Agency
// @router /:Ano [delete]
func (T *TestController) Delete() {
	Ano := T.GetString(":Ano")
	result, err := models.DeleteAgency(Ano)
	if err != nil {
		T.Ctx.WriteString(err.Error())
	}
	if result != true {
		T.Ctx.WriteString("Delete fail")
	}
	T.Data["json"] = "Delete success"
	T.ServeJSON()
}
