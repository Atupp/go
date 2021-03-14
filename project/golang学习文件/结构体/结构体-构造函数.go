package main

import "fmt"

type person struct {
	name, city string
	age        int8
}
//当结构体函数，开销比较大的时候，因为是值类型，是拷贝，所以说用值类型比较合适
//构造函数：调用这个函数，可以得到结构体对应的示例对象
func newperson(name,city string,age int8)*person{//调用这个函数，是创建了一个person类型
	return &person{
		name,
		city,
		age,
	}
}
func main() {
p1:=newperson("小王子","北京",int8(18))//int8(18)=18
fmt.Printf("%T" "%#v\n",p1,p1)
}
