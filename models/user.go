package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(100)"`
	Nickname string `orm:"size(100)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/beego_leaning?charset=utf8")
	orm.RegisterModel(new(User))
}

func Creat(name string, nickname string) (user User) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	fmt.Println("models被调用到了")
	newuser := new(User)
	newuser.Name = name
	newuser.Nickname = nickname
	o.Insert(newuser)
	return *newuser
}

func GetUser() (user []User) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	fmt.Println("models被调用到了")
	qs := o.QueryTable("user")
	var users []User
	var count int
	if count, err := qs.All(&users); err == nil {
		fmt.Println("查询数据成功")
		fmt.Println(count)
	}
	fmt.Println(count)
	return users
}

func Delete(id int64) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	newuser := new(User)
	newuser.Id = id
	o.Delete(newuser)
}

func GetOneUser(name string) (user User) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	user1 := User{Name: name}
	err := o.Read(&user1, "Name")
	fmt.Printf("ERR: %v\n", err)
	fmt.Println(user1.Id, user1.Name)
	return user1
}
