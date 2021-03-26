package main

import (
	"fmt"
)

type square struct {
	length int
	width  int
}

type circle struct {
	radius int
}

func (s square) area() float64 {
	return float64(s.length * s.width)
}

func (c circle) area() float64 {
	return 3.14 * (float64(c.radius) * float64(c.radius))
}

type shape interface {
	area() float64
}

func info(s shape) {

	switch s.(type) {
	case circle:
		fmt.Println("Area of circle:", s.(circle).area())
	case square:
		fmt.Println("Area of square:", s.(square).area())
	}
}

func main() {

	c := circle{
		radius: 10,
	}

	s := square{
		length: 4,
		width:  5,
	}

	info(c)
	info(s)
}
