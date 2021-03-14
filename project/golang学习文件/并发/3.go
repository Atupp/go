package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var wg sync.WaitGroup
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	wg.Done()
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(6)//1，只占一个CPU，有俩任务，但是只用一个人
	wg.Add(2)
	go a()
	go b()
	time.Sleep(time.Second)
	wg.Wait()
}
