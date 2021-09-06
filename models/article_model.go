package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Content    string
	Createtime string
}

//---------数据处理-----------
func AddArticle(article Article) error {
	err := insertArticle(article)
	return  err
}



func insertArticle(article Article) error {
	o :=orm.NewOrm()
	u := Article{Title: article.Title, Tags:article.Tags, Content: article.Content, Createtime: article.Createtime}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func GetArticele() ([]Article, error) {
	o := orm.NewOrm()
	var all1 []Article
	_, err := o.Raw("select id,title,tags,content,Createtime from article").QueryRows(&all1)

	//_, err := o.Raw("select id,filepath,filename,status,createtime from album" ).ValuesList(&lsits)
	if err != nil {
		return nil, err
	}

	return all1, nil
}
func GetArticelePag(pag int, count int) ([]Article, error) {
	o := orm.NewOrm()
	var all1 []Article
	_, err := o.Raw("select * from article limit ? offset ?",pag,count).QueryRows(&all1)

	if err != nil {
		return nil, err
	}

	return all1, nil
}
func GetArticeleCout() (int64, error) {
	o := orm.NewOrm()
	qs, _ := o.QueryTable("article").Count()
	return qs, nil
}

func GetArticeleDetails(id int) (Article, error) {
	o := orm.NewOrm()
	var u Article
	err := o.Raw("select * from article where id = ?",id).QueryRow(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}
func UpdateArticele(id int, title string, tags string, content string, Createtime string) Article {
	o := orm.NewOrm()
	var u Article
	err := o.Raw("update article set title= ?,tags = ?,content = ?,createtime = ? where id=?" , title,tags,content,Createtime,id,).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}