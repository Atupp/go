package main

//
import (
	"fmt"
	"html/template"
	"net/http"
)

//
//type User struct {
//	Name string
//	Age int
//	gender string
//}
//func sayhello(w http.ResponseWriter ,r *http.Request){
////定义模板
//
//	//解析模板
//	t,err:=template.ParseFiles("./f.tmpl")
//	if err!=nil {
//		fmt.Println("error:%v",err)
//		return
//	}
//	//渲染模板
//	u1:=User{
//		"安畅",
//		8,
//		"男",
//	}
//	t.Execute(w,u1)
//}
//func main() {
//http.HandleFunc("/",sayhello)
//err:=http.ListenAndServe(":9000",nil)
//	if err!=nil {
//		fmt.Printf("err:%v",err)
//	}
//	return
//}

// main.go
type person struct {
	name string
	age int
	add string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	z:=person{
		"赵泽宇",
		18,
		"深圳",
	}
	bbb:=tmpl.Execute(w, z.add)
	if bbb!=nil {
		fmt.Println("666")
		return
	}

}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
