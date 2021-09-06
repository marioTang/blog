package controllers

import (
	"devops/models"
	"fmt"
	"strconv"
)

type ShowController struct {
	BaseController
}


func (this *ShowController) Get() {

	pa := this.GetString("id")

	fmt.Println(pa)
	paint, _ := strconv.Atoi(pa)
	data, _ := models.GetArticeleDetails(paint)
	this.Data["Id"] = data.Id
	this.Data["Title"] = data.Title
	this.Data["Tags"] = data.Tags
	this.Data["Content"] = data.Content
	this.Data["Createtime"] = data.Createtime


	this.TplName = "write.html"
}

func (this *ShowController) Post() {


}
