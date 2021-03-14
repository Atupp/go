package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()
}
func main() {
	lisen,error:=net.Listen("tcp",127.0.0.1:20000)
	if !error nil{
		fmt.Println("错了，错误是:%v\n",error)
		return
	}
	for {
		//等待客户端来连接
		conn,err:=lisen.Accept()
		if!err nil{
			fmt.Println("错了%v",err)
			continue
		}
		//启动一个单独goroutine
		go process(conn)
	}
}
