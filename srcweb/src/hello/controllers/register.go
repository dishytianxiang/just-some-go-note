package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"hello/models"
	"strings"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	
	this.TplNames = "register.html"
}
func (this *RegisterController) Post() {
	email := this.Input().Get("email")
	passwd := this.Input().Get("passwd")
	nickName := this.Input().Get("nickName")
	fmt.Println("email = " + email)
	fmt.Println("passwd = " + passwd)
	str := models.AddUser(email,passwd,nickName)
	fmt.Println("str = " + str)
	bo := strings.EqualFold(str, "exist")
	fmt.Println(bo)
	if strings.EqualFold(str, "exist") || strings.EqualFold(str, "servererr") {
		this.Data["content"] ="this email has register!!"
		this.TplNames = "register.html"
		return
	}
	//this.Ctx.WriteString("hello world")
	this.SetSession("user",nickName)
	this.Ctx.Redirect(302, "/view")
	return
}
