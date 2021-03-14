package main

import "fmt"

func main() {
	//元素类型为map的切片
	//	var mapslice=make([]map[string]int,8,8)//这个切片是，树状的，每个切片值都有一堆MAP,定义的时候大的是什么类型，定义方式就是那个大的格式
	//	//含需要完成内部map元素的初始化
	//	mapslice[0]=make(map[string]int,8)//完成内部MAP初始化
	//	mapslice[0]["安畅"]=100//意思是mapslice[0]这个切片[0]值对应的MAP，键值对“安畅”为100
	//fmt.Println(mapslice)
	//值为切片的MAP,一个key，对应一个切片/一堆数
	var slicemap = make(map[string][]int, 8) //完成MAP的初始化
	//判断值有没有存在在这个MAP中
	v, ok := slicemap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		slicemap["中国"] = make([]int, 8)
		slicemap["中国"][0] = 100
		slicemap["中国"][1] = 200
		slicemap["中国"][2] = 300
	}
//for 遍历
	for k,v:=range  slicemap{
		for k,v:=range slicemap{

		}//for  两个值，循环不同的类型，不一样，但是都是对应的=表示把循环赋值给这两个元素，range 变量
fmt.Println(k,v)
	}
}
