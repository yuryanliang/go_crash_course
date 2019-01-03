package main

import (
	"fmt"
	"math"
)

//define interface
type Shape interface {
	//put the method you want your interface to implement here
	area() float64
	perimeter() float64
}

//use the interface with two structures
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

//structure methods
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) perimeter() float64 {
	return c.radius * 2 * math.Pi
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

//interface func
func getArea(s Shape) float64 {
	return s.area()
}

func getPerimeter(s Shape) float64 {
	return s.perimeter()
}

//sample data
func main() {
	circle := Circle{0, 0, 10}
	rectangle := Rectangle{10, 10}

	fmt.Println(getArea(circle))
	fmt.Println(getPerimeter(rectangle))
}
