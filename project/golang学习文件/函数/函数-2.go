package main

import "fmt"

func add(x,y int)int  {
	return x+y
}
func calc(x,y int,op func(int,int) int) int {//那么OP是函数，也是接收变量吗？
	return op(x,y)
}
func main() {
	//fmt.Println("sdssd")
	//defer fmt.Println("1")//相当于堆栈一样，后进先出
	//defer fmt.Println("2")
	//defer fmt.Println("3")
	//fmt.Println("end...")
r1:=calc(100,200,add)//这一步相当于执行操作，执行操作可以指定函数的输入变量是谁，操作集是谁
fmt.Println(r1)
fmt.Printf("%T",r1)
}
