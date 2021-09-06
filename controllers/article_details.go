package controllers

import (
	"devops/models"
	"fmt"
	"strconv"
)

type DetailsController struct {
	BaseController
}


func (this *DetailsController) Get() {

	pa := this.GetString("id", "1")
	fmt.Println(pa)
	paint, _ := strconv.Atoi(pa)
	data, _ := models.GetArticeleDetails(paint)
    fmt.Println(data.Title)
	this.Data["Title"] = data.Title
	this.Data["Tags"] = data.Tags
	this.Data["Content"] = data.Content
	this.Data["Createtime"] = data.Createtime
	this.Data["Id"] = data.Id
	this.TplName = "details.html"
}

func (this *DetailsController) Post() {

}
