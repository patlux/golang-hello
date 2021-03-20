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

	t.Run("Should calc area for a rectangle", func(t *testing.T) {
		rec := Rectangle{Width: 2.5, Height: 2.0}
		got := rec.Area()
		want := 5.0
		assertNumbers(t, got, want)
	})

	t.Run("Should calc area for a circle", func(t *testing.T) {
		circle := Circle{6.5}
		got := circle.Area()
		want := 132.73228961416876 
		assertNumbers(t, got, want)
	})
}
