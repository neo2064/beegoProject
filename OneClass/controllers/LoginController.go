package controllers

import (
	"OneClass/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct{
	beego.Controller
}

func (lg * LoginController)Login(){

	//userName := lg.GetString("")
	lg.TplName="login.html"

}
func (lg * LoginController)Post(){
	lg.TplName="login.html"

	lgOrm := orm.NewOrm()
	lgUser := new(models.User)

	lgUser.Name = lg.GetString("userName")
	beego.Info(lgUser.Name)
	//lgOrm.Read(lgUser,"name")
	//lgOrm.Delete(&lgUser,"name")
	err := lgOrm.Read(lgUser,"name")
	if err!=nil {
		beego.Info("用户不存在")
		beego.Info(err)
		return
	}

	if lgUser.PassWord != lg.GetString("password") {
		beego.Info("密码错误")
		return
	}

	//lg.Ctx.WriteString("登录成功")

	//lg.TplName="add.html/"
	lg.Redirect("/add/?userName="+lgUser.Name,302)

}