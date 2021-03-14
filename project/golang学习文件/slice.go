package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3] // s := a[low:high]
	//fmt.Printf("%v %v %v\n", s, len(s), cap(s))//%v是默认格式表示
	//MAKE创建切片
	//a := make([]int, 3) //make([]type, size, cap)
	//b := a
	//b[0] = 100
	//fmt.Println(b)
	//切片遍历
	//a := []int{1, 2, 3, 4}
	//for i := 0; i < len(a); i++ {//这个意思是从初始化开始，然后执行循环体，判断终止条件，i++
	//	fmt.Println(i,a[i])
	//	fmt.Printf("%v\n\n",len(a))
	//	for key,value:=range a{
	//		fmt.Println(key,value)
	//for _,value:=range a{//用不着哪个，直接用_代替
	//	fmt.Println(value)
	//二维数组
	//a:= [3][2]string{//第一个[]里面是最大的有几个数组，后面的[]是每个大数组里面有几个小数组
	//	{"sds", "聪明蛋"},
	//	{"qq", "ww"},
	//	{"rrrr", "hhh"},
	//}
	//a[0][0]="安畅"
	//fmt.Println(a)
	//fmt.Println(a[2][0])
	//for _,V1:=range a{
	//	for _,V2:=range V1{
	//		fmt.Println(V2)
	//	}
	//append函数扩容切片
//	var a []int//此时并没有申请内存地址，仅仅是申明了，没有初始化（这个才是申请地址）
//	//a =append(a,10)//append函数必须用某个变量接受,前面是原有的，后面是要添加的
//	for i := 0; i < 2; i++ {
//		b:=[]int{111}
//		a=append(a,b...)//可以直接追加另一个切片
//	}
//	c:=make([]int,5,6)
//	copy(c,a)//前面是目标函数，后面是操作函数，COPY后，两个函数不是同一个底层数组，不同于c:=a(这个里面用的是同一个底层数组)
//	fmt.Printf("%v\n %d  %d\n %p",a,len(a),cap(a),a)
	//切片元素的删除
	//a:=[]string{"1","2","3","4"}
	//a=append(a[0:2],a[3:]...)//为什么会出现替换现象？不是后面跟着吗？怎么还挤出去了？
	//fmt.Println(a)
	//a:=[]int{1,2,3}
	//b:=[]int{5}
	//a=append(a[0:2],b...)//append概念是覆盖/叠加
	//fmt.Println(a)
	//	var a = make([]string, 5, 10)//长度为5，容量是10，会保证后面是满的，所以前5个是零值（空格）
	//	for i := 0; i < 10; i++ {
			//a = append(a, fmt.Sprintf("%v", i))
			//fmt.Sprintf("%d", i)//这个是什么？为什么打印不出来东西？这个打印函数是干嘛用的
			//fmt.Println(i)
		//}
		//fmt.Println(a)
	//切片元素排序
	var a=[...]int{1,5,2,4}//这个是个数组
	sort.Ints(a[:])//这个包是用来排序的,但是接收的是切片
	fmt.Println(a)
	for k,v:=range a{
		fmt.Println(k)
		fmt.Println(v)
	}
}
