package main

import (
	"fmt"
	"math"
	"time"
)

type Rectangle struct {
	leftX  float64
	TopY   float64
	height float64
	width  float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func main() {
	/*
		listOfNums := []float64{1.1, 2.34}
		fmt.Println("sum is", addNums(listOfNums))

		num1, num2 := twoValue(5)
		fmt.Println("return two values:", num1, num2)

		fmt.Println("sub is", subNums(1, 2, 4))

		fmt.Println("factorial is", factorial(3))

		fmt.Println("safeDiv:", safeDiv(3, 0))
		fmt.Println("safeDiv:", safeDiv(3, 1))

		panicDemo()

		//pointer
		x := 0
		changeXVal(&x)
		fmt.Println(x)
		fmt.Println("Memory Address for x:", &x)

		yPtr := new(int)
		fmt.Printf("%T\n", yPtr)
		changeYVal(yPtr)
		fmt.Println("yPtr is now:", *yPtr)

		rect1 := Rectangle{leftX: 0, TopY: 50, height: 10, width: 10}
		fmt.Println("Area of the rectangle =", rect1.area())

		//interface
		rect := Rectangle1{20, 50}
		circ := Circle{4}

		fmt.Println("Rectangle Area =", getArea(rect))
		fmt.Println("Circle Area =", getArea(circ))

		fmt.Println("What is your name? ")

		var name string

		//fmt.Scan(&name)

		fmt.Println("Hello", name)

	*/

	for j := 0; j < 5; j++ {
		fmt.Println("loop in main start, id =", j)
		go count(j)
		fmt.Println("loop in main after go count, id =", j)
	}

	time.Sleep(time.Millisecond * 11000)

}

func count(id int) {
	fmt.Println("count func start...id =", id)
	for i := 0; i < 5; i++ {
		pre := ""
		if id == 0 {
			pre = "***"
		}
		fmt.Println(pre, "before sleep: count id =", id, "; i =", i, time.Now().Format("2:10:13.000000PM"))

		// Pause the function for 1 second to allow other functions to execute
		time.Sleep(time.Second)
		fmt.Println(pre, "after sleep: count id =", id, "; i =", i, time.Now().Format("2:10:13.000000PM"))

	}
	fmt.Println("count func end...id =", id)

}

type Shape interface {
	area() float64
}

type Rectangle1 struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle1) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {

	return shape.area()

}
func changeYVal(yPtr *int) {
	*yPtr = 100
}

func changeXVal(i *int) {

	*i = 3

}

func panicDemo() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC")

}

func safeDiv(i1 int, i2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := i1 / i2
	return solution
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * factorial(i-1)
}

func subNums(nums ...int) int {
	final := 0
	for _, value := range nums {
		final -= value
	}
	return final
}

func twoValue(i int) (int, int) {
	return i + 1, i + 3
}

func addNums(listOfNums []float64) float64 {
	sum := 0.0
	for _, value := range listOfNums {
		sum += value
	}
	return sum
}
