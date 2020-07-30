/*
分页: Count
	Limit(显示多少条,从什么位置开始)
	All
ceil floor

视图函数  {{.var | funcName}}
beego.AddFuncMap("name(html)",funcName)

参数 | 前面的值 返回值 {{}}里面的值
QueryTable

model1.

[]代表多个
需要链接的外部表


pk  -> name *struct `orm:"rel(fk/m2m)"`

fk -> name []*struct `orm:"reverse(many)"`

Filter 过滤器函数  ("查询的字段/表__字段名",value)

多表查询使用的是惰性查询 不指定表不查询
RelatedSle("表名")
RelatedSle 只能查询有数据的内容,没有数据就查询不出来

*/


package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c * MainController) Hello(){
	c.Layout="hello.html"
	c.Data["good"] ="hello"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["abc"] ="abc.html"
	c.TplName="world.html"

}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type IndexController struct{
	beego.Controller
}
func (c *IndexController) Post(){
	c.Data["testData"] = "这是一个测试数据"
	c.TplName = "test.html"
}
func (c *IndexController) Get(){
	//c.Data["testData"] = "这是一个测试数据"
	name := c.GetString(":name")
	path := c.GetString(":path")
	ext := c.GetString(":ext")
	beego.Info("name =",name)
	beego.Info("path =",path)
	beego.Info("ext =",ext)
	c.TplName = "index.tpl"
}
