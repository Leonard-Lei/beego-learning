package routers

import (
	"beego-learning/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/create", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/user/form", &controllers.UserController{}, "get:UserForm")
	beego.Router("/user/delete/:id", &controllers.UserController{}, "get:DeleteUser")
	beego.Router("/user/toLogin", &controllers.UserController{}, "get:GetloginPage")
	beego.Router("/user/login", &controllers.UserController{}, "post:Login")
}
