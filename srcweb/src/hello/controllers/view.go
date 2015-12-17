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
