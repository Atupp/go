package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"                  //导入ORM工具
	_ "github.com/jinzhu/gorm/dialects/mysql" //导入GORM工具驱动
	"time"
)
//当列名由两个单词组成,那么对应到数据库中,中间用_连接,具体列名怎么修改,看下面
type UserInfo struct {
	ID uint
	Name sql.NullString//表示空值
	Gender string
	Hobby string

}
//定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}


func main() {
	//连接MYSQL数据库
	db,err:=gorm.Open("mysql","root:ANTONG7230800@(127.0.0.1:3306)/db1")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	u1:=UserInfo{ID:2}
	//把数据放入数据库
	//在数据库中,创建+改名
	fmt.Println(db.NewRecord(&u1))
	db.Create(&u1)
	fmt.Println(db.NewRecord(&u1))
}
