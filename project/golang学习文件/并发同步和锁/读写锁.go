package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x int64
	lock sync.Mutex
	wg sync.WaitGroup
	rwlock sync.RWMutex
	)
func read(){
	//lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Microsecond)
	//lock.Unlock()
	wg.Done()
	rwlock.RUnlock()
}
func write(){
	//lock.Lock()
	rwlock.Lock()
	x=x+1
	time.Sleep(time.Microsecond*10)
	//lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}
func main() {
start:=time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
