package main

import "fmt"

func a() {
	fmt.Println("func a")
}
func b() {
	//假如b函数可能触发PANIC，那么就要在出触发前通过defer，定义延迟调用的匿名函数（函数延迟）
	defer func() {
		err := recover() //recover是把当前的错误信息收集一下
		if err != nil {
			fmt.Println("你有错啦")
		}

	}()
	panic("panic in b") //（）里面的东西表示提示内容
}
func c() {
	fmt.Println("func c")
}
func main() {
	a()
	b()
	c()
}
