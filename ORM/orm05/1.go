package main

import (
	"fmt"
	"github.com/jinzhu/gorm"                  //导入ORM工具
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
db,err:=gorm.Open()
db.First()
}
