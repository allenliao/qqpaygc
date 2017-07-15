package routers

import (
	"qqpaygc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Login", &controllers.LoginController{})
	beego.Router("/Index", &controllers.IndexController{})

	beego.Router("/Suggest", &controllers.BetsuggestController{})
	beego.Router("/Suggest/:BUCode(BU[0-9]+)", &controllers.BetsuggestController{})
}
