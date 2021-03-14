package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func hello(a int){
	fmt.Println("hello world",a)
	wg.Done()
}
func main() {//主goroutine,主的运行完，次的运行不运行也会关闭
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go hello(i)//并发了，同一段时间内，这一段时间指的是程序开始到最后.次goroutine
	}
	fmt.Println("456")
	//time.Sleep(time.Second)//等一会
	wg.Wait()
}
