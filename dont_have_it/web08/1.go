package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("users.tmpl")
	if err!=nil{
		fmt.Println("66")
		return
	}
	//渲染模板
	ho:="小王子"
	t.Execute(w,ho)
}
func sss(w http.ResponseWriter,r *http.Request){
	//定义模板
	//在模板中添加新函数
	t,err:=template.New("sss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string)template.HTML{
			return template.HTML(s)
		},
		//解析模板
	}).ParseFiles("./sss.tmpl")
	if err!=nil{
		fmt.Println("66")
		return
	}
	//渲染模板
	str1:="<script>alert(123);</script>"//经过了转义,html/template这个包东西
	str2:="<a href='http://liwenzhou.com'>liwenzhou博客</a>"
	t.Execute(w,map[string]string{
		"str1":str1,
		"str2":str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/sss", sss)
	err := http.ListenAndServe(":9102", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:%v", err)
		return
	}
}