package controllers

import (
	"OneClass/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddController struct {
	beego.Controller
}


func (ad *AddController)Add(){
	ad.TplName="add.html"
	aName := ad.GetString("userName")
	beego.Info("aName = ",aName)
	aOrm := orm.NewOrm()
	aUser := new(models.User)
	aUser.Name = aName
	err := aOrm.Read(aUser,"name")
	if err != nil {
		beego.Info("用户名不存在")
		return
	}

	ad.Data["name"] = aUser.Name

	hello123 := "good"
	beego.Info(hello123)
	aArticle := new(models.Article)
	aArticle.Id = aUser.Id
	beego.Info("aArticle = ",aArticle.Id)
	err2 := aOrm.Read(aArticle,"id")
	if err2!=nil {
		beego.Info("表中无记录")
		return
	}

	aType := []string{"体育新闻","财经新闻","科技新闻"}
	ad.Data["title"] = aArticle.Title
	ad.Data["type"] = aType
	ad.Data["content"] = aArticle.Content
	ad.Data["img"] = aArticle.Img

}

func (ad *AddController)Post(){


	file,head,err := ad.GetFile("uploadname")
	if err != nil {
		beego.Info("文件不存在")
		return
	}
	defer file.Close()
	fileName := head.Filename;
	beego.Info(fileName)
	ad.SaveToFile("uploadname","./static/img/"+fileName)
	ad.Ctx.WriteString("success")


}