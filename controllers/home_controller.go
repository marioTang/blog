package controllers

import (
	"devops/models"
	"devops/utils"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
		pa := 1
		pre_page :=3
		totals,_:=models.GetArticeleCout()
		res :=utils.Paginator(pa, pre_page, totals)
		datalist,_:=models.GetArticelePag(3,(pa-1)*3)
		fmt.Println("首页")
		this.Data["datas"] = datalist
		this.Data["paginator"] = res
		this.Data["totals"] = totals

		this.TplName = "home.html"
	}







