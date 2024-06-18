package main

import (
	"fmt"
	"math"
)

type Point struct { // Point with math-accroding floats
	//values are private
	x float64
	y float64
}

func (p Point) Dist(other Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2)) // formula of distance between two points
}

func NewPoint(x, y float64) Point {
	return Point{x, y} // make new point
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(0, 12)
	fmt.Println(a.Dist(b))
}
