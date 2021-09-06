package controllers

import (
	"devops/models"
	"time"
)

type EditorController struct {
	BaseController
}


func (this *EditorController) Get() {
	this.TplName = "write.html"
}

func (this *EditorController) Post() {
	if this.Loginuser == nil{
		var response map[string]interface{}
		response = map[string]interface{}{"code": 0, "message": "error", "textmsg":"请登录"}
		this.Data["json"] = response
		this.ServeJSON()
	}else {
		title := this.GetString("title")
		content := this.GetString("origin")
		//content := this.GetString("content")
		tags := this.GetString("tags")
		//blogid := this.GetString("blogid")
		//转换时间
		timeUnix:=time.Now().Unix()   //已知的时间戳
		formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")

		//写入文章
		art := models.Article{0, title, tags,  content, formatTimeStr}
		err := models.AddArticle(art)

		//返回数据给浏览器
		var response map[string]interface{}

		if err == nil {
			//无误
			response = map[string]interface{}{"code": 1, "message": "ok", "textmsg":"创建成功"}
		} else {
			response = map[string]interface{}{"code": 0, "message": "error", "textmsg":"创建失败"}
		}

		this.Data["json"] = response
		this.ServeJSON()

	}



}
