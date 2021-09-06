package controllers

import (
	"devops/models"
	"devops/utils"
	"fmt"
	"strconv"
)

type GetpagController struct {
	BaseController
}

func (this *GetpagController) Get() {

	pa := this.GetString("id", "1")
	paint, _ := strconv.Atoi(pa)

	fmt.Println(pa)
	pre_page :=3
	totals,_:=models.GetArticeleCout()

	res :=utils.Paginator(paint, pre_page, totals)
	datalist,_:=models.GetArticelePag(pre_page, (paint-1)*3)

	this.Data["datas"] = datalist
	this.Data["paginator"] = res
	this.Data["totals"] = totals

	this.TplName = "home.html"
	}








