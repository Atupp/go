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
	Name string `gorm:"column:6666"`//把这列的列名修改成666
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
//1把User表名字改为profiles
func (User) TableName() string {
	return "profiles"
}

func main() {
	//连接MYSQL数据库
	db,err:=gorm.Open("mysql","root:ANTONG7230800@(127.0.0.1:3306)/db1")
	if err!=nil {
		panic(err)
	}
	defer db.Close()
	//在表名有前缀,只会修改默认表明,单独指定的不会修改.
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "prefix_" + defaultTableName;//必须在创建表时间之前使用
	}
	//禁用表名字复数形式
	db.SingularTable(true)
	//创建表 自动连接
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserInfo{})//不回删除数据,名字改后,只会重新建一个
	//创建数据行
	u1:=UserInfo{1,"兄奥","男","篮球"}
	u2:=UserInfo{2,"解决","女","跳舞"}
	//把数据放入数据库
	//在数据库中,创建+改名
	db.Table("乱七八糟").CreateTable(&User{})
	db.Create(&u1)
	db.Create(&u2)
	//查询数据
	{
		var u UserInfo
		db.First(&u, u2)
		fmt.Println(u)
		//更新
		db.Model(&u).Update("hobby", "是多少")
		//删除
		db.Delete(&u)
	}

}
