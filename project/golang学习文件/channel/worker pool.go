package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) { //每个工人有三个属性，ID，执行哪个任务，任务结果是什么？
	for j := range jobs {                                  //循环的是jobs的通道,打印出的哪个任务。
		fmt.Printf("worker:%d start job:%d\n", id, j) //哪个工人在干哪个任务
		results <- j * 2                              //干的活是把任务*2传送到results通道中。
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j) //哪个工人结束哪个任务
	}
}
func main() {
	jobs := make(chan int, 100)    //通道内容是job，一共5个。
	results := make(chan int, 100) //结果在result中
	// 开启3个goroutine,有三个人在做//工作池
	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j //把0-4存入jobs，一共5个

	}
	// 输出结果
	for i := 0; i < 5; i++ {
		ret:= <-results//遍历通道。返回两个数据，一个是通道数据，一个是正确与否。
			fmt.Println(ret)
	}
}
