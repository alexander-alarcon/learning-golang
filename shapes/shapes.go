// Package shapes provides shapes functions related to area and perimeter
package shapes

import "math"

type PerimeterOptions struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle
// rules:
// - if any dimension is equal to 0, return -1
// - if any dimension is negative, convert it to positive
func Perimeter(options PerimeterOptions) float64 {
	if options.Width == 0 || options.Height == 0 {
		return -1.0
	}

	width := math.Abs(options.Width)
	height := math.Abs(options.Height)

	return 2 * (width + height)
}

// Area returns the area of a rectangle
// rules:
// - if any dimension is equal to 0, return -1
// - if any dimension is negative, convert it to positive
func Area(options PerimeterOptions) float64 {
	if options.Width == 0 || options.Height == 0 {
		return -1.0
	}

	width := math.Abs(options.Width)
	height := math.Abs(options.Height)

	return width * height
}
