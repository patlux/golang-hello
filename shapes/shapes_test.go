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

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{2.5, 2.0}, 5.0},
		{Circle{6.5}, 132.73228961416876},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertNumbers(t, got, tt.want)
	}
}
