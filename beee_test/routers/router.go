package routers

import (
	"beee_test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/hello", &controllers.MainController{}, "*:Hello")
}
