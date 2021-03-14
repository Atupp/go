package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)
var m=make(map[int]int)
var m2=sync.Map{}

func get(key int)int{
	return m[key]
}
func set(key int,value int){//这个函数就是赋予key对应的值的
	m[key]=value
}
//func main() {
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(i int) {
//			set(i,i+100)
//			fmt.Printf("key-%v value-%v",i,get(i))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}
//原生的map不支持并发读写，需要加锁，或者是使用sync.map
func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i,i+100)
			value,_:=m2.Load(i)
			fmt.Printf("key-%v value-%v\n",i,value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}