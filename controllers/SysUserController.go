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
}

