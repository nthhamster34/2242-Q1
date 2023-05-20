package main

import "testing"


func TestTriangle(t *testing.T) {
  triangle := Triangle {
    base: 10, height: 12,
  }

  expected_a := 60.0
  got_a := triangle.area()

  expected_p := 37.62049935181331
  got_p := triangle.perimiter()

  if got_a != expected_a {
    t.Errorf("Expected %v got %v", expected_a, got_a)
  }

  if got_p != expected_p {
    t.Errorf("Expected %v got %v", expected_p, expected_p)
  }
}
