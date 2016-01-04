package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	user :=this.GetSession("user")
	fmt.Println(user)
	this.Data["user"] = user
	this.TplNames = "view.html"
}
func (this *ViewController) Document() {
	this.TplNames = "document.html"
}
