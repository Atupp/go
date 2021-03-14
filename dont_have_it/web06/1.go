package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数
	k:= func(name string)(string,error) {
		return name+"真帅",nil
	}

	t := template.New("f.tmpl")
	//在解析之前,告诉模板加了一个函数
	t.Funcs(template.FuncMap{
		"kua99":k,
	})

	_,err:=t.ParseFiles("./f.tmpl")//解析模板
	//创建模板的新函数的时候,模板命名一定要在解析模板里面(可以解析多个模板)
	if err != nil {
		fmt.Println("66")
		return
	}
	name := "小王子"
	//渲染模板
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9500", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:%v", err)
		return
	}
}
