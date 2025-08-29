package main

import (
	"fmt"
	"math"
)

func main() {
	first := NewPoint(4, 5)
	second := NewPoint(10, 10)

	distance := first.Distance(second)

	fmt.Println("Distance:", distance)
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}
