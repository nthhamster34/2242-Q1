package main

import(
  "testing"
)

func TestCircle(t *testing.T) {
    circle := Circle {
      radius: 12,
    }
    
    if circle.area() != 452.0 {
      t.Error("Expected 452")
    }

    if circle.perimiter() != 75.0 {
      t.Error("Expected 75")
    }
    
}
