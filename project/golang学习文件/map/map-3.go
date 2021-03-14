package main

import (
	"fmt"
	"strings"
)

func main() {
	//统计一个字符串中每个单词的出现次数，how do you do
	//1定义一个MAP，前面是单词，后面是单词的次数
	//2
	//3遍历单词做统计
	var a =make(map[string]int,10)
	var b ="how do you do"
	//1.字符串都有哪些单词,得到的是一个切片
	words:=strings.Split(b," ")//
	//遍历单词做统计
	for _,word:=range words{
		v,ok:=a[word]
		if ok{
			//MAP中有这个单词的统计记录
			a[word]=v+1
		}else {
			//MAP中没有这个单词
			a[word]=1
		}
	}
	for k,v:=range a{
		fmt.Println(k,v)

	}
}
