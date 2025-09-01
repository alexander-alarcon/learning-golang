// Package shapes provides shapes functions related to area and perimeter
package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle
// rules:
// - if any dimension is equal to 0, return -1
// - if any dimension is negative, convert it to positive
func Perimeter(rectangle Rectangle) float64 {
	if rectangle.Width == 0 || rectangle.Height == 0 {
		return -1.0
	}

	width := math.Abs(rectangle.Width)
	height := math.Abs(rectangle.Height)

	return 2 * (width + height)
}

// Area returns the area of a rectangle
// rules:
// - if any dimension is equal to 0, return -1
// - if any dimension is negative, convert it to positive
func Area(rectangle Rectangle) float64 {
	if rectangle.Width == 0 || rectangle.Height == 0 {
		return -1.0
	}

	width := math.Abs(rectangle.Width)
	height := math.Abs(rectangle.Height)

	return width * height
}
