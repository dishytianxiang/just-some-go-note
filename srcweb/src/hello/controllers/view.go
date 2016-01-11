package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ViewController struct {
	beego.Controller
}
type Person struct {
	Name   string
	Age    int
	Sex    string
	Addr   string
}
func (this *ViewController) Get() {
	user :=this.GetSession("user")
	fmt.Println(user)
	this.Data["user"] = user
	this.TplNames = "view.html"
}
func (this *ViewController) Document() {
	var data = make(map[int]Person,10)
	var trueName string
	for i:=0;i<10;i++ {
		trueName = "dishy" + strconv.Itoa(i)
		data[i] = Person{Name: trueName,Age: 25,Sex: "female",Addr: "baoanbaoti"}
	}
	this.Data["data"] = data
	this.TplNames = "document.html"
}
