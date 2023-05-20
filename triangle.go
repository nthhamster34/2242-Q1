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
    return (t.height * t.base)/2
}

func (t Triangle) perimiter() float64 {
    h := math.Sqrt(math.Pow(t.height, 2) + math.Pow(t.base, 2))
    return h + t.height + t.base
}

func main() {
  triangle := Triangle {
    base: 10, height: 12,
  }
  fmt.Println(triangle.area())
  fmt.Println(triangle.perimiter())
}
