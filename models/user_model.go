package models

import (
	"github.com/astaxie/beego/orm"
	//"github.com/prometheus/common/model"
)

type Users struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

//--------------数据库操作-----------------

//插入
func InsertUser(user Users) error {
	o :=orm.NewOrm()
	u := Users{Id: user.Id, Username: user.Username, Password: user.Password, Status: user.Status, Createtime: user.Createtime}
	_, err := o.Insert(&u)
    if err != nil {
		return err
	}
	return nil
}


//根据用户名查询id
func QueryUserWithUsername(username string) int {
	var u Users
//	orm.RegisterModel(new(Users))
	o :=orm.NewOrm()
	err := o.Raw("select id from users where username = ?",username).QueryRow(&u)
	if err != nil {
	}
	return u.Id

}
//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	var u Users
//	orm.RegisterModel(new(Users))
	o :=orm.NewOrm()
	err := o.Raw("select id from users where username = ? and password = ? ",username,password).QueryRow(&u)
	if err != nil {
		//fmt.Println(err)
	}
	return u.Id

}
