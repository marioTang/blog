package models

import (
	"github.com/astaxie/beego/orm"
)

type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	Createtime int64
}
func InsertAlbum(album Album) error {
	o :=orm.NewOrm()
	u := Album{Filepath: album.Filepath, Filename: album.Filename, Status: album.Status, Createtime: album.Createtime}
	_, error := o.Insert(&u)
	if error != nil {
		return error
	}
	return nil
}
func GetMyads() ([]Album, error) {
	o := orm.NewOrm()
	var all []Album
	//var lsits []orm.ParamsList
	_, err := o.Raw("select id,filepath,filename,status,createtime from album").QueryRows(&all)

	//_, err := o.Raw("select id,filepath,filename,status,createtime from album" ).ValuesList(&lsits)
	if err != nil {
		return nil, err
	}

	return all, nil
}