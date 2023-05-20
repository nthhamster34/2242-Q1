package main

import (
  "testing"
)

func TestSquare(t *testing.T) {
  perimiter, area := square(2)

  if perimiter != 8 {
    t.Error("Expected 8 got", perimiter)
  }

  if area != 4 {
    t.Error("Expected 4 got", area)
  }
  
}
