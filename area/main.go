package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (s triangle) getArea() float64 {
	return 0.5 * s.base * s.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(i shape) {
	fmt.Println(i.getArea())
}

func main() {
	test01 := &triangle{
		height: 2,
		base:   4,
	}

	test02 := &square{
		sideLength: 5,
	}

	printArea(test01)
	printArea(test02)
}
