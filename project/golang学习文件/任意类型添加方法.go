package main

import "fmt"

//任性类型添加方法，但是必须是我代码包的，或者说是我这个GO文件中的
type myint int

func(i myint) intGO()  {
fmt.Println("hi")
}
func main() {
	var m1 myint
	m1.intGO()
}
