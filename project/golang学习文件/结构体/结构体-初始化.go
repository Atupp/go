package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

func main() {
	//1键值对初始化
p4:=person{
	name: "555",
	age: 8,
	city: "北京",
}
fmt.Printf("%#v\n",p4)
	//2值的列表初始化
	p5:=person{
		"小王子",
		"北京",
		18,
	}
	fmt.Printf("%#v\n",p5)
}
