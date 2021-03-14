package main

import "fmt"

//文档：https://www.liwenzhou.com/posts/Go/05_array/   视频：https://www.bilibili.com/video/BV1Bb41137Ga
func main() {
	//数组初始化
	var a [3]int
	var b = [2]int{222, 666}
	f := b[:]
	f[0] = 12
	var c = [...]string{"哈哈", "你妹"} //与切片的区别是，[...]数组，[]切片
	d := [...]int{0: 100, 3: 10}
	fmt.Println(a) //最普通的打印方式
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", d)
	fmt.Println(d)
	fmt.Println(f)
}
