package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2.0 // area = (b * h) / 2.0.
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func main() {
	t := triangle{base: 4, height: 3}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)
}
