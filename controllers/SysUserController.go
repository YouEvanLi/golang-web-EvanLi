package controllers

import (
	"fmt"
)

type SysUserController struct {
	SysBaseController
}

func (c *SysUserController) LoginForm() {
	c.TplName = "sysuser/login_form.tpl"
}

func (c *SysUserController) LoginAction() {
	username := c.Ctx.Request.Form.Get("username")
	fmt.Println(username)
<<<<<<< HEAD
=======

>>>>>>> d6221628a0020094f12082f8625f17b75cf8a2b5
}
