package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
	_ "hello/routers"
//	"fmt"
//	"github.com/astaxie/beego/context"
)

func init() {
	models.RegisterDB()
}
//var FilterUser = func(){
//	user :=this.GetSession("user")
//	fmt.Println(user)
//	fmt.Println("hello filterUser")
//}
func main() {

	beego.SessionOn = true
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
//	beego.InsertFilter("/*",beego.BeforeRouter,beego.InsertFilter("/*",beego.BeforeRouter,FilterUser))
	beego.Run()
}

