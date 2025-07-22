package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}

func (p Point) Distance(o Point) float64 {
	a := p.GetX() - o.GetX()
	b := p.GetY() - o.GetY()
	return math.Sqrt(a*a + b*b)
}

func main() {
	p := NewPoint(-4, -1)
	o := NewPoint(8, 7)
	fmt.Println(p.Distance(o))
}
