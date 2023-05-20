package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) area() float64 {
	return (t.base * t.height) / 2
}

func (t Triangle) perimeter() float64 {
	side := math.Sqrt(math.Pow(t.base/2, 2) + math.Pow(t.height, 2))
	return 2*side + t.base
}

func main() {
	triangle := Triangle{
		base:   16, // Updated base value
		height: 10, // Updated height value
	}
	fmt.Println(triangle.area())
	fmt.Println(triangle.perimeter())
}
