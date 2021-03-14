package main

import "fmt"

/*
两个同步任务
1生成0-100数字发送到ch1
2从ch1取出来，计算数据平方，发送到ch2
*/
func f1(ch chan<- int) { //<-表示通道只接受数据//生成0-100随机数,这个应该是参数，有的参数，输入进入的是接收者。
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
//<-限制数据，只读或者只写，通过通道
//从通道取值的方式1
func f2(cha <-chan int, chb chan<- int) {
	for {
		tmp, ok := <-cha //如果cha中有值，那么Ok就是True.
		if !ok {
			break
		}
		chb<-tmp*tmp//这两种都行，只不过是简写
		//} else {
		//	chb <- tmp * tmp
		//}
	}
	close(chb)
}
func main() {
	ch1 := make(chan int, 100)//ch1是个双向的，可以接受也可以发送值
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)
	//从通道取值的方式2
	for ret := range ch2 { //当遍历通道的时候，如果隐形ok为false，那么会自动关闭，退出循环
fmt.Println(ret)
	}
}
