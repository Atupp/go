package main

import "fmt"

func modify1(x int){
	x=100
}
func mdify2(y *int)  {
	*y=100
}
func main() {
	/*
a:=10
b:=&a//*int表示整型类型的指针，同时*地址=那个值，&这个是取地址操作
fmt.Printf("%v %p\n",a,&a)
fmt.Println(b)
c:=*b
fmt.Println(c)

	 */
	a:=1
	modify1(a)
	fmt.Println(a)
}
