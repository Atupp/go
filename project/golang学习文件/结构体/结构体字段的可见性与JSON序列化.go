package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性与JSON序列化
//GO中标识符首字母大写，那么对外可访问
type student struct {
	ID   int    //这个就是结构体的字段
	Name string //结构体字段
}
type class struct {
	Title    string`json:"title" db:"student" xml:"ss"`//比如说这个title，如果他首字母小写，那么其他函数是调用不了的，那么它的值是零值
	//但是，如果非要用小写的表示也行，那么需要用到tag,比如说在bd里面是student,xml里面是ss
	Students []student
}
func newstudent(id int,name string)student{
	return student{
		id,
		name,
	}
}

func main() {
	c1 := class{
		Title:    "火箭101班",
		Students: make([]student, 0, 20),
	}
	//往班级中添加学生
	for i := 0; i < 10; i++ {
		//添加10学生
		tmpstu:=newstudent(i,fmt.Sprintf("stu%02d",i))
		c1.Students=append(c1.Students,tmpstu)
	}
	fmt.Println(c1.Students)
	//JSON序列化：GO语言中的数据→JSON格式的  字符串   网址是www.json.cn
	data,err:=json.Marshal(c1)//把C1的内容转化成字符串
	if err!=nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%s\n",data)
	fmt.Printf("%T\n",data)
	//JSON反序列化：JSON→GO语言中数据
	jsonstr:=`{"Title":"火箭101班","Students":[{"ID":0,"Name":"小王子"},{"ID":1,"Name":"娜扎"},{"ID":2,"Name":"stu02"}]}`
	var c2 class
	err=json.Unmarshal([]byte(jsonstr),&c2)
	if err!=nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%#v\n",c2)
	}
