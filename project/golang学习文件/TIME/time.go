package main

import (
	"fmt"
	"time"
)

func main() {
	now:=time.Now()
	fmt.Println(now)
	//时间戳
	now.Unix()//返回的是秒数，从1970年1月1日
	now.UnixNano()//返回纳秒数
	//时间戳→时间
	//time.Unix()//TIME包的一个函数
	//时间间隔
	//time.Duration()
	//time.Sleep()//默认是纳秒，函数表示程序暂停，然后继续执行
	//时间操作
	now.Add(time.Hour)
	//时间计时器
	//时间格式化
	a:=now.Format("2006：01：02")//用2006/01/02/15/04/05
	fmt.Println(a)
	//根据时区去解析字符串的时间

}
