package main

import "fmt"
import "time"
import "strconv"

// CHANNELS
// Channels allow us to pass data between go routines

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {

	pizzaNum++

	// Convert int into a string
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and Send for Sauce")

	// Send the pizzaName onto the channel for the next
	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func addSauce(stringChan chan string) {

	// Receive the value passed on the channel
	pizza := <-stringChan

	fmt.Println("Add Sauce and Send", pizza, "for Toppings")

	// Send the pizzaName onto the channel for the next
	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func addToppings(stringChan chan string) {

	// Receive the value passed on the channel
	pizza := <-stringChan

	fmt.Println("Add Toppings to", pizza, "and Ship")

	time.Sleep(time.Millisecond * 10)

}

func main() {

	// Make creates a channel that can hold a string
	// int channel intChan := make(chan int)
	stringChan := make(chan string)

	// Cycle through and make 3 pizzas
	for i := 0; i < 3; i++ {

		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)

	}

}
