package main

import "fmt"

//匿名函数和闭包
func a(name string) func() {
	return func() {//返回一个匿名函数
	fmt.Println(name)
	}//这块表示，这个匿名函数内部对变量的操作，其变量依赖于外部输入
}
func main() {
	//sayhello:=func(){//用一个变量接收，然后带动了这个匿名函数，一会引用这个匿名函数，就直接引用这个变量
	//	fmt.Println("匿名函数")//当这一行末尾加了（）表示，命名匿名函数并且立即执行，就不用下面再引用了
	//}
	//sayhello()
	//闭包
	r:=a("沙河小王子")//a赋值给这个r,这个r是函数类型，相当于换了函数名字
	r()//执行r,r不仅是一个内部的匿名函数，还包括外部变量的引用，r就是一个闭包
	fmt.Printf("%T",r)
}
