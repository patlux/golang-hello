package shapes

import (
	"testing"
)

func assertNumbers(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	assertNumbers(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(2.5, 2)
	want := 5.0
	assertNumbers(t, got, want)
}
