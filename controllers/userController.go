package controllers

import (
	"beego-learning/models"
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

/*
 * 1、判断用户是否存在
 * 2、存在则验证密码
 * 3、登录成功之后将用户信息写入session
 */
func (c *UserController) Login() {
	name := c.GetString("name")
	user := models.GetOneUser(name)
	if !reflect.ValueOf(user).IsValid() {
		fmt.Println("用户信息不存在")
		c.TplName = "fail.html"
		return
	}
	c.SetSession("userInfo", user)
	c.Ctx.Redirect(302, "/user")
	return
}

func (c *UserController) Get() {
	us := c.GetSession("userInfo")
	if us == nil {
		c.TplName = "userLoginForm.html"
		return
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	fmt.Println("Controller被调用到了")
	var users []models.User
	users = models.GetUser()
	c.Data["list"] = users
	c.TplName = "user.html"
}

func (c *UserController) UserForm() {
	c.TplName = "userForm.html"
}

func (c *UserController) CreateUser() {
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

func (c *UserController) DeleteUser() {
	id := c.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	fmt.Println(err)
	models.Delete(intId)
	c.TplName = "success.html"
}

func (c *UserController) GetloginPage() {
	c.TplName = "userLoginForm.html"
}
