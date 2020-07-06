package routers

import (
	"github.com/astaxie/beego"
	"golang-web-EvanLi/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("sys_user/login_form", &controllers.SysUserController{}, "Get:LoginForm")
	beego.Router("sys_user/login_form", &controllers.SysUserController{}, "Post:LoginAction")
}
