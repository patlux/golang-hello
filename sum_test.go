package main

import "testing"

func TestSum(t *testing.T) {
	given := [5]int{1, 2, 3, 4, 5}

	actual := Sum(given)
	expected := 15

	if actual != expected {
		t.Errorf("Expected: %d / Received: %d, %v", expected, actual, given)
	}
}
