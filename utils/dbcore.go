package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitMysql1()  {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	driverName := beego.AppConfig.String("driverName")
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	orm.Debug = true


	//登陆方法
	errs :=orm.RegisterDriver(driverName,orm.DRMySQL)
	if errs != nil {
		fmt.Println("注册设备连接失败")

	}
	fmt.Println("msyql驱动注册成功")

	err :=orm.RegisterDataBase("default" ,driverName, dbConn )


	if err != nil {
		fmt.Println("mysql连接失败")

	}

	fmt.Println("MySQL数据库连接成功")
	fmt.Println("初始化表")
	//orm.RegisterModel(new(models.Article))
	//orm.RunSyncdb("default", false, true)
}

