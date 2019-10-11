package _struct

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type AbstractShape interface {
	Area() float64
}

type Cycle struct {
	Radius float64
}

func (cycle Cycle) Area() (area float64) {
	return math.Pi * cycle.Radius * cycle.Radius
}

func (cycle Cycle) Perimeter() (perimeter float64) {
	return math.Pi * 2 * cycle.Radius
}

type Rectangle struct {
	Width float64;
	Height float64;
}

func (rectangle Rectangle) Area() (area float64) {
	return rectangle.Width * rectangle.Height
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
	return (rectangle.Height + rectangle.Width) * 2
}


type Triangle struct {
	Width float64
	Height float64
}

func (triangle Triangle) Area() (area float64)  {
	return triangle.Width * triangle.Height / 2
}

func (triangle Triangle) Perimeter() (perimeter float64) {
	return
}