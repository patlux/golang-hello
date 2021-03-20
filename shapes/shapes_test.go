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
	got := Perimeter(Rectangle{Width: 10.0, Height: 10.0})
	want := 40.0
	assertNumbers(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{Width: 2.5, Height: 2.0})
	want := 5.0
	assertNumbers(t, got, want)
}
