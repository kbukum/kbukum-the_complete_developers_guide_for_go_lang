package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Printf("Area for %+v is %2.f", s, s.getArea())
	fmt.Println()
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	side float64
}

func (t square) getArea() float64 {
	return t.side * t.side
}

func main() {
	t := triangle{
		base:   3,
		height: 4,
	}
	printArea(t)
	s := square{
		side: 4,
	}
	printArea(s)
}
