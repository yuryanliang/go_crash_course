package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum1 = 0
var pizzaName1 = ""

func main() {
	stringChan := make(chan string)
	for i := 0; i < 1; i++ {
		go makeDoughNew(stringChan)
		go addSauceNew(stringChan)
		go Topping(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}
}

func makeDoughNew(stringChan chan string) {
	pizzaNum1++
	pizzaName1 := "pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough for", pizzaName)

	stringChan <- pizzaName1
	time.Sleep(time.Millisecond * 10)

}

func addSauceNew(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("add Sauce for", pizza)
	stringChan <- pizza
	//time.Sleep(time.Millisecond * 10)

}
func Topping(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("add Topping for", pizza)
	time.Sleep(time.Millisecond * 10)

}
