package main

import "fmt"

//方法，作用于特定变量的函数，建议接收者名字为其类似的第一个小写字母
//方法的定义示例
type Person struct {
	name string
	age int8
}
//Newperson为person类型的构造函数（通过或者函数运行，就构造出来一个人）
func NewPerson(name string,age int8)*Person  {
	return &Person{
		name: name,
		age: age,
	}
}
//dream是为person类型（接收者）定义的方法
func(p Person) Dream()  {
	fmt.Printf("%s的梦想是学好GO语言\n",p.name)
}
//修改年龄的方法,这个是接收者类型是指针
func (p *Person) Setage(newage int8)  {
	p.age = newage
}
//接收者是值类型的修改年龄方法
func (p Person) Setage2(newage int8)  {
	p.age = newage
}
func main() {
//如何使用这些东西，包括构造函数+方法
	P1:=NewPerson("tom",16)//它是person类型的指针变量
	//(*P1).Dream()//GO自动把指针类型转换为相对于的变量，故，可简写,见下方
	P1.Dream()
	fmt.Println(P1.age)
	P1.Setage(19)
	fmt.Println(P1.age)
	P1.Setage2(30)
	fmt.Println(P1.age)
}
//什么时候使用指针接收者
//1需要修改接收者的值；2接收者拷贝起来代价大；3保证一致性，其他方法用了指针，咱也用指针

