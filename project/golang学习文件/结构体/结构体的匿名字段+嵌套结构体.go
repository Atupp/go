package main

import "fmt"

//结构体的匿名字段
type person struct {
	string
	int
	address
}
type address struct {
	MPH string
	street string
}
func main() {
	/*
	p:=person{
		"小王子",
		21,
		address{//嵌套结构体
			"203",
			"花园大道",

		},
	}
	fmt.Println(p)
	fmt.Println(p.string,p.int,p.address)

	 */
	//结构体继承

}
