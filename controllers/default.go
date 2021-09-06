package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


//固定路由匹配
//func (this *MainController) Get() {
//	beego.Info(" 固定路由的 get类型的方法 ")
//	this.Ctx.Output.Body([]byte(" 固定路由 Get 请求!!"))
//}
//func (this *MainController) Post() {
//	beego.Info(" 固定路由的 post类型的方法 ")
//	this.Ctx.Output.Body([]byte(" 固定路由 Post 请求!!"))
//}
//
//func (this *MainController) Delete() {
//	beego.Info(" 固定路由的 Delete类型的方法请求 ")
//	this.Ctx.Output.Body([]byte("固定路由的 Delete 请求"))
//}
//
//func (this *MainController) Options() {
//	beego.Info(" 固定路由的 options 类型的方法请求 ")
//	this.Ctx.Output.Body([]byte("固定路由的 options 请求"))
//}


    //正则路由匹配
//type RegController struct {
//	beego.Controller
//}
//func (this *RegController) Get() {
	//*全匹配
	//beego.Info("全匹配：" + this.Ctx.Input.URL())
	//this.Ctx.Output.Body([]byte("请求URL：" + this.Ctx.Input.URL()))

	//变量匹配
	//id := this.Ctx.Input.Param(":name")
	//beego.Info("Id:" + id)
	//this.Ctx.ResponseWriter.Write([]byte("Id:" + id))

	//*.*匹配
	//path := this.Ctx.Input.Param(":path")
	//beego.Info(path)
	//ext := this.Ctx.Input.Param(":ext")
	//beego.Info(ext)
	//this.Ctx.ResponseWriter.Write([]byte("filePath: " + path + " , ext: " + ext))

	//int类型匹配 只能匹配int类型的
	//id := this.Ctx.Input.Param(":id")
	//this.Ctx.ResponseWriter.Write([]byte("int类型变量值：" + id))
	//
	////string类型匹配 只能匹配string类型
	//userName := this.Ctx.Input.Param(":username")
	//this.Ctx.ResponseWriter.Write([]byte("string类型变量值：" + userName))
    //}

    //自定义路由
//type CustomController struct {
//	beego.Controller
//}
//func (this *CustomController) GetUserInfo() {
//	beego.Info("获取用户信息")
//
//	username := this.GetString("username")
//
//	userid := this.GetString("userid")
//
//	this.Ctx.Output.Body([]byte("获取用户信息请求,用户名：" + username + " , 用户编号：" + userid))
//}
//func (this *CustomController) RegisterInfo() {
//	beego.Info("注册用户信息")
//	username := this.GetString("username")
//	userid := this.GetString("userid")
//	this.Ctx.Output.Body([]byte("用户注册信息，username：" + username + " , userid：" + userid))
//}