package main

import "fmt"

//写入专用通道
//一次写10条
func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i	//主线程不能产生死锁,所以此处报错
		fmt.Println("send:", i)
	}
}

//消费专用通道
//一次只取10条
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}


ch := make(chan int)
//只生产和消费10条记录
go produce(ch)
go consumer(ch)