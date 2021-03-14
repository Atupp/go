package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main() { //主goroutine,主的运行完，次的运行不运行也会关闭
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(a int) {
			fmt.Println("hello", a)//因为越到后面，I的取值范围越大
			wg.Done()
		}(i)//这个是啥意思？意思是传入外界参数的位置
	}
	fmt.Println("123")
	//time.Sleep(time.Second)//等一会
	wg.Wait()
}