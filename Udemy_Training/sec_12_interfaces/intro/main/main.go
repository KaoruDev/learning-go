package main

import (
	"math"
	"fmt"
)

type Square struct {
	name string
	side float64
}

func (sq Square) area() float64 {
	return math.Pow(sq.side, 2)
}

type Circle struct {
	name string
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

type Shape interface {
	area() float64
}

func info(shape Shape) {
	fmt.Println(shape)
	fmt.Println(shape.area())
}

func main() {
	s := Square{"square", 10}
	c := Circle{"circle", 5}
	info(s)
	info(c)
}
