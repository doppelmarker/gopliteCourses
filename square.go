package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	var end Point
	end.x += int(s.a) + s.start.x
	end.y -= int(s.a) + s.start.y
	return end
}

func (s Square) Perimeter() int {
	return 4 * int(s.a)
}

func (s Square) Area() int {
	return int(s.a * s.a)
}

func squareTask() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
