package controllers

type AboutController struct {
	BaseController
}


func (this *AboutController) Get() {
	this.TplName = "about.html"
}

