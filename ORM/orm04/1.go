package main

import (
	"fmt"
	"github.com/jinzhu/gorm"                  //导入ORM工具
	_ "github.com/jinzhu/gorm/dialects/mysql" //导入GORM工具驱动
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	//连接MYSQL数据库
	db, err := gorm.Open("mysql", "root:ANTONG7230800@(127.0.0.1:3306)/db1")
	if err != nil {
		panic(err)
		defer db.Close()
	}
	//把表和模型对应起来
	db.AutoMigrate(&User{})
	//把数据放入数据库
	//u1 := User{ Name: "安通", Age: 18}
	//db.Create(&u1)
	//u2 := User{Name: "赵泽宇", Age: 18}
	//db.Create(&u2)
	////查询
	var u User
db.First(&u)
fmt.Println(u)
var user []User
	db.Where("name = ?", "安通").First(&user)
fmt.Println(user)
}
