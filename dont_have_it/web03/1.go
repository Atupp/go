package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayhello(w http.ResponseWriter,r *http.Request ){//一个请求,一个相应
	b,_:=ioutil.ReadFile("./1.txt")
_,_=fmt.Fprintln(w,string(b))
}


func main() {
	http.Handle()
	http.HandleFunc("/hello",sayhello)
	//启动服务端
	err:=http.ListenAndServe("9090",nil)
	if err!=nil {
		fmt.Print("66")
		return
	}
}
