package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    this.Data["isHome"] = true
	this.TplNames = "home.html"
}
