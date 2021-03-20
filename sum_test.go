package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Should sum a array of any size", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}

		actual := Sum(given)
		expected := 15

		if actual != expected {
			t.Errorf("Expected: %d / Received: %d, %v", expected, actual, given)
		}
	})

	t.Run("Should sum a array with a fixed size", func(t *testing.T) {
		given := [5]int{3, 5, 6, 2, 8}

		actual := Sum(given[:]) // pass a slice
		expected := 24

		if actual != expected {
			t.Errorf("Expected: %d / Received: %d, %v", expected, actual, given)
		}
	})

}

func assertSum(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected: %d / Received: %d", got, want)
	}
}

func TestSumAll(t *testing.T) {

	t.Run("Should sum a list of slices", func(t *testing.T) {
		actual := SumAll([]int{3, 4}, []int{2, 7})
		expected := []int{7, 9}
		assertSum(t, actual, expected)
	})

}

func TestSumAllTails(t *testing.T) {

	t.Run("Sum all tails", func(t *testing.T) {
		actual := SumAllTails([]int{3, 4, 2}, []int{2, 7})
		expected := []int{6, 7}
		assertSum(t, actual, expected)
	})

	t.Run("Sum all tails even its empty", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{2, 7})
		expected := []int{0, 7}
		assertSum(t, actual, expected)
	})

}
