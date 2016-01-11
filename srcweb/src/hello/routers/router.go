package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
	"github.com/astaxie/beego/context"
	"fmt"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, loginFilter)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/view", &controllers.ViewController{})
	beego.Router("/view/document", &controllers.ViewController{},"get:Document")
	beego.Router("/logout", &controllers.LogoutController{})
}
func loginFilter(ctx *context.Context) {
    
	user := ctx.Input.CruSession.Get("user")
	//user := this.GetSession("user")
	fmt.Println("loginFilter user" )
		fmt.Println(user)
	if user == nil && ctx.Request.RequestURI != "/" && ctx.Request.RequestURI != "/login"{
        ctx.Redirect(302, "/")
    }
}
