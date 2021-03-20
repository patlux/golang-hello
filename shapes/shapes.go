package shapes

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height 
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func Area(rec Rectangle) float64 {
	return rec.Width * rec.Height
}
