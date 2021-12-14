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

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func main() {
	c := Circle{Radius: 10}
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)
	r := Rectangle{Height: 10, Width: 10}
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}

type ShapeWithArea interface {
	Area() float64
}

func PrintArea(sa ShapeWithArea) {
	fmt.Printf("Area = %f\n", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(sa ShapeWithPerimeter) {
	fmt.Printf("Perimeter = %f\n", sa.Perimeter())
}

/*
type Shape interface {
	Area() float64
	Perimeter() float64
}
*/

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(s Shape) {
	PrintArea(s)
	PrintPerimeter(s)
}
