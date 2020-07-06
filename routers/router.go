package routers

import (
	"golang-web-EvanLi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("sys_user/login_form",&controllers.SysUserController{},"Get:LoginForm")
	beego.Router("sys_user/login_form",&controllers.SysUserController{},"Post:LoginAction")
}
