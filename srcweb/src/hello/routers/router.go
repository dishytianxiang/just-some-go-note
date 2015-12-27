package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/view", &controllers.ViewController{})
	beego.Router("/view/document", &controllers.ViewController{},"get:Document")
	beego.Router("/logout", &controllers.LogoutController{})
}
