/*1.基本路由
post->post 方法  get请求找-> Get方法  注意 方法名首字母一定要大写

2.高级路由
在执行 beego.Router 方法的时候 有第三个参数
	1. beego.Router("/",&Controller.ControllerStruct(),"get:Func"  )  1 vs 1
	2. beego.Router("/",%Controller.ControllerStruct(),"get,post:Func") 1 vs 多
	3. beego.Router("/",&Controller.ControllerStruct(),"get:Get;post:Post" 多 vs 多
	4. beego.Router("/",&Controller.ControllerStruct(),"*:Func;get:Get" 多 vs 1 + 1 对 1
							全部执行 func方法  但是 get 请求执行 Get 方法->如果指定了则执行指定的

	刷新细节:刷新是重新请求上一次的方法(get post)
	注意细节: 1.方法名一定大写
			2.当你自定以了请求方法,那就不会去执行默认的 get post 方法需要你重新指定
			3.*代表全部 格式: "请求:方法名;请求2:方法 2"  方法没有括号

3.正则路由
? : 0个或1个
	?:id  -> 将 0 个或 1 个值赋值给 :id
	this.GetString(":id")

	:id  至少有一个

	*.*
	:path  -> .前面的*
	:ext   ->  .后面的*
	* -> :splat  包括所有的 splat的内容

* : 0 个或多个
\w: 字母
[0-9]:数字一个
[0-9]+ :


*/

package routers

import (
	"OneClass/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.MainController{},"get:Hello")
    beego.Router("/index/*", &controllers.IndexController{})
    beego.Router("/login", &controllers.LoginController{},"get:Login;post:Post")
    beego.Router("/register", &controllers.RegisterController{},"get:Register;post:Post")
    beego.Router("/content", &controllers.ContentController{},"get:ShowContent;post:Post")
    beego.Router("/add/*", &controllers.AddController{},"get:Add;post:Post")

}
