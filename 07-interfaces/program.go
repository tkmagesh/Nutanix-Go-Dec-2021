package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func main() {
	c := Circle{Radius: 10}
	PrintArea(c)
	r := Rectangle{Height: 10, Width: 10}
	PrintArea(r)
	shapesWithArea := []ShapeWithArea{
		c, r,
	}
	for _, shape := range shapesWithArea {
		fmt.Println(shape.Area())
	}
}

type ShapeWithArea interface {
	Area() float64
}

func PrintArea(sa ShapeWithArea) {
	fmt.Printf("Area = %f\n", sa.Area())
}
