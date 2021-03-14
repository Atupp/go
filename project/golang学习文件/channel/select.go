package main

import "fmt"

func main() {
	ch :=make(chan int,1)
	for i := 0; i < 10; i++ {
		select {
		case x:=<-ch://每个case对应一个通道的发送和接受；select会一直等待，case们都是同步的，直到某个case通信完成。然后执行对应的分支。
			fmt.Println(x)
			case ch<-i://没有操作
		default:
			fmt.Println("默认")
		}
	}
}
