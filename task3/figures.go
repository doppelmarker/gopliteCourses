package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	a, b float64
}

func (s Square) area() float64 {
	return s.a * s.b
}

func (s Square) perimeter() float64 {
	return 2 * (s.a + s.b)
}

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func figures() {
	var s Figure = Square{1.3, 2.8}
	var c Figure = Circle{4.2}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())
	fmt.Println()
}
