package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
}
func (this *LoginController) Post() {
	passwd := this.Ctx.Request.Form["passwd"]
	email := this.Ctx.Request.Form["email"]
	fmt.Print("======",passwd)
	fmt.Print("======",email)	
	fmt.Print("======",this.Ctx.Request.Form["email"])
	
	this.TplNames = "main.html"
}
