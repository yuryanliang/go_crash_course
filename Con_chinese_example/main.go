package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var ch chan int = make(chan int)
	go func() {
		fmt.Printf("开始阻塞1。。。")
		time.Sleep(time.Second)
		fmt.Printf("结束阻塞1。。。")
		ch <- 4
	}()

	fmt.Printf("ch ==>%d", <-ch)
}
