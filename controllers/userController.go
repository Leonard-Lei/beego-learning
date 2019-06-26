package controllers

import (
	"beego-learning/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	fmt.Println("Controller被调用到了")
	var users []models.User
	users = models.GetUser()
	c.Data["json"] = users
	c.ServeJSON()
	c.TplName = "user.html"
}

func (c *UserController) Post() {
	fmt.Println("Controller被调用到了")
	log := logs.NewLogger(10000)
	name := c.GetString("name")
	nickname := c.GetString("nickname")
	log.Informational(name)
	user := models.Creat(name, nickname)
	log.Informational("这是从modles层返回的数据")
	log.Informational(user.Name)
	c.TplName = "success.html"
}
