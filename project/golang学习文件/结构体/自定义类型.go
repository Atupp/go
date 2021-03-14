package main

import "fmt"

//自定义类型
type myint int
//类型别名
type newint=int//类型别名，类型还是int
func main() {
	var i myint
	fmt.Printf("%T\n %v\n",i,i)
	var b newint
	fmt.Printf("%T\n %v\n",b,b)
}
