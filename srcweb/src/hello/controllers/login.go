package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	
	this.TplNames = "register.html"
}
func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	email := this.Input().Get("email")
	passwd := this.Input().Get("passwd")
	if beego.AppConfig.String("email") == email &&
		beego.AppConfig.String("passwd") == passwd {
			this.Ctx.SetCookie("email",email,30,"/")
			this.Ctx.SetCookie("passwd",passwd,30,"/")
		this.Redirect("/view",301)
	}else {
	this.Redirect("/category",301)
	}
	return
}
