package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {//循环的是jobs的通道
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}
func main() {

}
