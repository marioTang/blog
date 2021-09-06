package main

import (
	"devops/models"
	_ "devops/routers"
	"devops/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	utils.InitMysql1()
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Article))
	orm.RegisterModel(new(models.Album))
	models.GetArticeleCout()

	beego.Run("0.0.0.0:8090")
}


