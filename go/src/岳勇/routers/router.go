package routers

import (
	"岳勇/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
}
