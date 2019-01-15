package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go MultiByTwo(in, out)
	go MultiByTwo(in, out)
	go MultiByTwo(in, out)
	in <- 1
	in <- 2
	in <- 3

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	//time.Sleep(time.Second)

}

func MultiByTwo(in <-chan int, out chan<- int) {
	i := <-in
	result := i * 2
	//fmt.Println(result)
	out <- result
}
