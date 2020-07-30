package controllers

import (
	"OneClass/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func ( rg *RegisterController )Register(){

	rg.TplName="register.html"
}


func (rg *RegisterController) Post(){

	rgUserN := rg.GetString("userName")
	rgPassW :=rg.GetString("password")
	beego.Info(rgUserN)
	beego.Info(rgPassW)

	rgOr := orm.NewOrm()
	var rgData = new(models.User)
	rgData.Name = rgUserN
	rgData.PassWord = rgPassW
	_,err := rgOr.Insert(rgData)
	if err != nil {
		beego.Info("插入数据错误")
		return
	}
	beego.Info("插入数据成功")

	//rg.TplName="login.html"
	rg.Redirect("/login",302)
}
