package models

import "fmt"

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(100)"`
	Nickname string `orm:"size(100)"`
}

func Creat(name string, nickname string) (user User) {

	fmt.Println("models被调用到了")
	newuser := new(User)
	newuser.Id = 1
	newuser.Name = "这是名称"
	newuser.Nickname = "这是昵称"
	return *newuser
}
