package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	
	this.TplNames = "view.html"
}
func (this *ViewController) Document() {
	this.TplNames = "document.html"
}