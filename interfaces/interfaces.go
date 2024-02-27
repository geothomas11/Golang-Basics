package main

import "fmt"

type shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}
func PrinntShapeInfo(s shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("perimeter: %f\n", s.Perimeter())

}

func main() {
	circle := Circle{Radius: 5}
	PrinntShapeInfo(circle)
}
