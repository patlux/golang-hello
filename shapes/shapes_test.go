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

func assertArea(t *testing.T, shape Shape, want float64) {
	t.Helper();
	got := shape.Area()
	assertNumbers(t, got, want);
}

func TestArea(t *testing.T) {

	t.Run("Should calc area for a rectangle", func(t *testing.T) {
		assertArea(t, Rectangle{Width: 2.5, Height: 2.0}, 5.0)
	})

	t.Run("Should calc area for a circle", func(t *testing.T) {
		assertArea(t, Circle{6.5}, 132.73228961416876)
	})
}
