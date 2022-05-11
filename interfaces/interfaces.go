//interface is a collection of method signatures and also defines the behaviour of the object
//method signature consist of method name, input parameters and return type

package main

import (
	"fmt"
	"math"
)

//declaring interface
type Shape interface {
	Area() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) Area() float64{
	return r.width * r.height
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s Shape
	fmt.Println(s)
	fmt.Printf("%T", s)
}