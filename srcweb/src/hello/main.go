package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	"hello/models"
	_ "hello/routers"
	//	"fmt"
	//	"github.com/astaxie/beego/context"
)

func init() {
	models.RegisterDB()
	//globalSessions, _ := session.NewManager("redis", `{"cookieName":"gosessionid","gclifetime":3600,"ProviderConfig":"127.0.0.1:6379,100,astaxie"}`)
	//go globalSessions.GC()
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
	beego.AddFuncMap("tooggle1",chu)
	//	beego.InsertFilter("/*",beego.BeforeRouter,beego.InsertFilter("/*",beego.BeforeRouter,FilterUser))
	beego.Run()
}
func chu(i int) (out int) {
	out = i % 2
	return
}
