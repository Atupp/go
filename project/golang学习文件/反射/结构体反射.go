package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"name" ini:"n"`
	Score int `json:"score" ini:"s"`
}
func main() {
	stu1:=student{
		"小王子",
		65,
	}
	t:=reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n" ,t.Name() ,t.Kind())
	for i:=0;i<t.NumField();i++{//后面是取得的是结构体的字段数量
		fileobj:=t.Field(i)//这个是每一行字段的相关信息，包括名字、类型、标签
		fmt.Printf("name:%v  kind:%v  tag:%v\n",fileobj.Name,fileobj.Type,fileobj.Tag)
		fmt.Println(fileobj.Tag.Get("json"),fileobj.Tag.Get("ini"))//Get表示的是获取
	}
	//根据名字去取结构体中的字段
	fileobj2,ok:=t.FieldByName("score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v",fileobj2.Name,fileobj2.Type,fileobj2.Tag)
	}
}
