package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"fmt"
	"reflect"
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
	/*if beego.AppConfig.String("email") == email &&
		beego.AppConfig.String("passwd") == passwd {
			this.Ctx.SetCookie("email",email,30,"/")
			this.Ctx.SetCookie("passwd",passwd,30,"/")
		this.Redirect("/view",301)
	}else {
	this.Redirect("/category",301)
	}*/
	users,err := models.GetUser(email,passwd)
	if err != nil || len(users) == 0{
		this.Data["msg"] = "user email or passwd wrong"
		this.TplNames = "home.html"
		return
	}
	fmt.Println(users)
	fmt.Println(reflect.TypeOf(users[0]))
	var people models.User
	people = users[0]
	fmt.Println(people.NickName)
	this.SetSession("user", people.NickName)
	this.Redirect("/view",301)
	return
}
