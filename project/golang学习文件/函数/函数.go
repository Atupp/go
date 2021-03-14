package main

import "fmt"

//定义没有参数、返回值的函数
//func sayhello()  {
//	fmt.Println("沙河小王子")
//}
//func main() {
//	//调用这个函数
//	sayhello()
//}
//定义一个接收string类型的name函数
func sayhello(name string) { //这个规定的是函数的参数，如果只定义函数的返回值应该是sayhello()(name string)
	fmt.Println(name)
	}
//定义接受多个参数类型的函数，一个返回值
func intsum(a int,b int)int {
	sum:=a+b
	return sum//表示返回，具体返回哪一个参数呢？有很多，return后面加
}
//定义接收可变参数
//可变参数在函数体中是切片
func intsum1(a...int)  {
fmt.Println(a)
fmt.Printf("%T\n",a)
}
//定义有多个返回值的函数,也支持类型简写
func cals(a,b int)(sum,sub int)  {
	sum=a+b
	sub=a-b
	return
}
//定义全局变量num
var num=10
//定义函数
func qjbl()  {
	num:=100//现在函数自己内部访问变量，找不见的话，那就往外找
	fmt.Println("全局变量",num)//可以在函数中访问全局变量

}
func main() {
//	//a := "哈哈哈哈"
//	//sayhello(a)
//	//R:=intsum(10,20)//R的赋值对象是这个intsum函数的返回值，a和b是两个参数
//	//fmt.Println(R)
//	//intsum1(10,30)
//x,y:=cals(100,200)
//fmt.Println(x,y)
//qjbl()
////语句块内部定义的变量,i变量只能在语句块中使用（外不能用内）
//	for i := 0; i < ; i++ {
//		fmt.Println(i)
//	}

	//}
	abc:=qjbl//相当于把函数赋值给这个函数ABC
	fmt.Printf("%T",abc)
}
