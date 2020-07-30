package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


type Home struct {
	Id int
	Name string
	Address string
	Number int
}

type User struct {
	Id int `orm:"pk;auto"`
	Name string `orm:"size(10)""`
	PassWord string `orm:"size(20)"`
}

type Article struct {
	Id int `orm:"pk";auto`
	Title string `orm:"size(50);not null""`
	Type string `orm:"size(20);not null`
	Content string `orm:"size(100);default(这是文章内容)"`
	Img string `orm:"size(30)"`
}
func init() {
	//orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:mysqlroot@tcp(:3306)/mysql?charset=utf8")
	orm.RegisterModel(new(Home),new(User),new(Article))
	orm.RunSyncdb("default",false,true)
}

func test(){
	test := []byte{byte(3)}
	beego.Info(test)

}