package controllers

import (
	"devops/models"
	"fmt"
	"strconv"
	"time"
)

type EditingArticlesController struct {
	BaseController
}


func (this *EditingArticlesController) Get() {

}

func (this *EditingArticlesController) Post() {
	pa := this.GetString("id")
	title := this.GetString("title")
	tags := this.GetString("tags")
	content := this.GetString("content")
	paint, _ := strconv.Atoi(pa)
    fmt.Println(content)
	timeUnix:=time.Now().Unix()   //已知的时间戳
	formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")
    art := models.Article{paint,title,tags,content, formatTimeStr}

    models.UpdateArticele(art.Id,art.Title,art.Tags,art.Content,art.Createtime)
	var response map[string]interface{}
    response = map[string]interface{}{"code": 1, "message": "ok", "textmsg":"更新成功"}

	this.Data["json"] = response
	this.ServeJSON()
	fmt.Println(pa)

}
