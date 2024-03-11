package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

type Shape interface {
	getArea() float64
}

func main() {

	t := Triangle{height: 12, base: 6}
	s := Square{sideLength: 12}

	printArea(t)
	printArea(s)
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh Shape) {
	fmt.Println(sh.getArea())
}
