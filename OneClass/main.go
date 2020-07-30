package main

import (
	_ "OneClass/routers"
	"github.com/astaxie/beego"
	_"OneClass/models"
)

func main() {

	beego.AddFuncMap("hi",hello)
	beego.Run()

}
func hello()(out string){
	out = "hello" + "world"
	return out;
}

