package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

func main() {
	/*
		var p person
		p.name="123"
		p.age=18
		p.city="北京"
		fmt.Printf("p=%#v\n",p)//%#v,打印的更全面
		fmt.Println(p.city)//可以通过点去访问
		//匿名结构体，适用于只用一次
		var user struct{
			name string
			married bool
		}
		user.name="555"
		user.married=false
		fmt.Printf("%#v",user)

	*/
	//通过指针，去操作这个结构体,如果这个结构体变量是指针，那么可以直接变量名称.结构体字段，去访问
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	(*p2).name = "555"
	p2.age = 21 //可以简写，在GO中
	//取结构体地址进行实例化（赋值）
	p3 := &person{}
	p3.age = 16
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
}
