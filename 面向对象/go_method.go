package main

import (
	"fmt"
	"math"
)
type Rectangle struct{
	width,height float64
}
type Circle struct{
	radius float64
}
func (r Rectangle) area() float64{
	return r.width*r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi 
}
func main() {
	r1 := Rectangle{2,2}
	r2 := Rectangle{4,5}
	r3 := Circle{10}
	r4 := Circle{2}

	fmt.Println("Area of r1 is :",r1.area())
	fmt.Println("Area of r2 is :",r2.area())
	fmt.Println("Area of r3 is :",r3.area())
	fmt.Println("Area of r4 is :",r4.area())
}