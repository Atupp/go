package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./index2.tmpl")
	if err!=nil{
		fmt.Println("66")
		return
	}
	//渲染模板
	msg:="小王子"
	t.Execute(w,msg)
}

func home(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./home2.tmpl")
	if err!=nil{
		fmt.Println("66")
		return
	}
	//渲染模板
	ho:="小王子"
	t.Execute(w,ho)
}
func home2(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/home2.tmpl")
	if err!=nil{
		fmt.Println("66")
		return
	}
	//渲染模板
	ho:="李逸"
	t.ExecuteTemplate(w,"home2.tmpl",ho)
}

func index2(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/index2.tmpl")
	if err!=nil{
		fmt.Println("66")
		return
	}
	//渲染模板
	ho:="小王子"
	t.ExecuteTemplate(w,"index2.tmpl",ho)
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home2", home2)
	http.HandleFunc("/index2", index2)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:%v", err)
		return
	}
}
