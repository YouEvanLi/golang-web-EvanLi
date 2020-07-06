package controllers

import (
	_"fmt"
	"github.com/astaxie/beego"
)

type SysBaseController struct {
	beego.Controller
	ControllerName 	string  //controller 名称
	ActionName		string  //具体方法
	CurrentUrl		string  // 路由
}

type ResponseJson struct {
	Code 	int
	Message string
	Data 	interface{}
}

func (b *SysBaseController)CheckAuth()  {
	b.ControllerName , b.ActionName = b.GetControllerAndAction()
	b.CurrentUrl = b.ControllerName +"."+ b.ActionName
	
}