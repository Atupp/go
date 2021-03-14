package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//文档：https://www.liwenzhou.com/posts/Go/06_slice/    视频：https://www.bilibili.com/video/BV1Db41137Se
func main() {
	//var a map[string]int//只声明，但是不初始化，就是零值,内存上没有相对应的位置
	//a=make(map[string]int,8)//map初始化并且赋值,8是容量
	//fmt.Println(a==nil)
	////如何赋值/添加键值对
	//a["安通"]=100
	//a["安畅"]=200
	//fmt.Println(a)
	//fmt.Printf("%#v\n",a)//%#v显示带井号，比平常多了一个#,这样会把引号表示出来，MAP的类型
	////声明MAP的同时完成初始化
	//b:=map[int]bool{//声明+初始化
	//	1:true,//赋值
	//	2:false,
	//}
	//fmt.Printf("%v\n",b)
	//fmt.Printf("%T",b)
	//var a=make(map[string]int,8)
	//a["小王"]=100
	//a["小赵"]=99
	//判断某个键存在？
	//value,ok:=a["小李"]//如果存在，小李对应的值赋给value且ok为true，如果不存在，value的值是零值且ok为false
	//if ok==true {//if 变量=if 这个变量为true？
	//	fmt.Println("小李存在这个里面中",value)
	//}else {
	//	fmt.Println("没有这个人")
	//}
	//for range遍历,键值对添加的顺序和遍历出来的顺序无关
	//for key,value:=range a{
	//	fmt.Println(key,value)
	//	}
	////只遍历一个key,不用下划线代替,遍历value，前面用下划线
	//for key:=range a{
	//	fmt.Println(key)
	//	}
	//删除键值对
	//delete(a,"小王")
	//	fmt.Println(a)
	//按照固定顺序去遍历
	//var b = make(map[string]int,5)
	//b["656"]=600
	//b["777"]=100
	//fmt.Printf("%02d", b)
	var a = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%2d", i)//那这个函数表示啥意思？？两位数的10进制，stu是学生的简写
		value:=rand.Intn(100)//0-99的随机整数
		a[key]=value
	}
	//for k,v:=range a{
	//	fmt.Println(k,v)
	//}
	//按照key的从小到大的顺序去遍历
	//1先取出来所有key，存放切片中（为啥是切片）
	keys:=make([]string,0,100)//([]string,100)=([]string,100,100),所以说不要简写
	for k:=range a{//一个数对应一个切片的值，就用到了for range，一个一个放入
		keys=append(keys,k)
	}
	//2做排序
	sort.Strings(keys)//keys目前是个有序的切片
	//3按照排序后的key对v排序
	for _,v:=range keys{//遍历切片，前面的是切片的下坐标，后面是值
		fmt.Println(v,a[v])
	}
}
