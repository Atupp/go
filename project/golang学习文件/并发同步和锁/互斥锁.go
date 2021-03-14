package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex//互斥锁
)
func add(){
	for i := 0; i < 5000; i++ {
		lock.Lock()//锁住
		x=x+1//三步，1从全局中取x，2 X+1，3 送回去X//那么修改的时候，给原数据加个锁
		lock.Unlock()//解锁
	}
	wg.Done()
}

func main() {
wg.Add(2)
go add()
go add()
wg.Wait()
fmt.Println(x)
}
