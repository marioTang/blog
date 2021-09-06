package controllers

import (
	"devops/models"
	"github.com/prometheus/common/log"
)

type AlbumController struct {
	BaseController
}


func (this *AlbumController) Get() {
	albums,err := models.GetMyads()
	if err !=nil{
		log.Error(err)
	}
	this.Data["Album"] = albums
	this.TplName="album.html"
}