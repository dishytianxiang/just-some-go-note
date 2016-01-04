package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
	"github.com/astaxie/beego/context"
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
	if user == "" && ctx.Request.RequestURI != "/" {
        ctx.Redirect(302, "/")
    }
}
