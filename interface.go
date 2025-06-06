package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	length, breadth float32
}

func (c Circle) Area() float32{
	return math.Pi*c.radius*c.radius
}

func (c Circle) Perimeter() float32{
	return 2*math.Pi*c.radius
}

func (r Rectangle) Area() float32{
	return r.length*r.breadth
}

func (r Rectangle) Perimeter() float32{
	return 2*(r.length+r.breadth)
}

func main(){
	var s Shape = Circle{radius:7}

	fmt.Println("Area of c : ", s.Area())
	fmt.Println("Perimeter of c : ", s.Perimeter())

	s = Rectangle{length: 12, breadth: 8}

	fmt.Println("Area of r : ", s.Area())
	fmt.Println("Perimeter of r : ", s.Perimeter())
}
