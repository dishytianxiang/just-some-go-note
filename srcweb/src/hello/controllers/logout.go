package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.Redirect("/",301)
	return
}