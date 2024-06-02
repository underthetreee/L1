package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// Creates new point with coordinates
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Return point coordinates
func (p Point) Coordinates() Point {
	return Point{p.x, p.y}
}

// Calculates distance between two points
func distance(p1, p2 *Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Create two points
	point1 := NewPoint(1, 2)
	point2 := NewPoint(3, 4)

	// Calculates distance between them
	dist := distance(point1, point2)

	fmt.Println(point1.Coordinates(), point2.Coordinates(), dist)
}
