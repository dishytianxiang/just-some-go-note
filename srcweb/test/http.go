package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello,this is version 1.")
}

/*type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("hello world")
}
func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}*/
