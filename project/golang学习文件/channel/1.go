package main

import "fmt"

func main() {
	//有缓冲区通道
var ch1 chan int//引用类型，必须分配内存（初始化）才能使用
ch1=make(chan int,1)//缓冲区大小，1表示最多放一个值，叫异步通道
	ch1=make(chan int)//无缓冲通道，类似快递员，必须把快递给你手上才行，不然就等，又叫同步通道（必须同时完成）
ch1<-10
x:=<-ch1
fmt.Println(x)
len(ch1)//取得通道元素数量
cap(ch1)//拿到通道容量
close(ch1)//手动关闭了通道
}
