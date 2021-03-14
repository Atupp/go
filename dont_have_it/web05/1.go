package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type user struct {
	Name string
	Gender string
	Age int
}
func sayHello(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./f.tmpl")
	if err!=nil {
		fmt.Println("66")
		return
	}
	u1:=user{
		"小王子",
		"男",
		18,
	}
	m1:=map[string]interface{}{
		"N":456465,
		"B":66666,
		"C":55555,
	}
t.Execute(w,map[string]interface{}{
	"u1":u1,
	"m1":m1,
})
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
