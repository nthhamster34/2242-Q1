package main

import "testing"

func TestTriangle(t *testing.T) {
	triangle := Triangle{
		base:   16,
		height: 10,
	}

	expected_a := 80.0 // Updated expected area
	got_a := triangle.area()

	expected_p := 41.612496949731394 // Updated expected perimeter
	got_p := triangle.perimeter()

	if got_a != expected_a {
		t.Errorf("Expected %v got %v", expected_a, got_a)
	}

	if got_p != expected_p {
		t.Errorf("Expected %v got %v", expected_p, got_p)
	}
}
