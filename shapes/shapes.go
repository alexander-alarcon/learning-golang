// Package shapes provides shapes functions related to area and perimeter
package shapes

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle
// rules:
// - if any dimension is equal to 0, return -1
// - if any dimension is negative, convert it to positive
func (rectangle Rectangle) Perimeter() float64 {
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
func (rectangle Rectangle) Area() float64 {
	if rectangle.Width == 0 || rectangle.Height == 0 {
		return -1.0
	}

	width := math.Abs(rectangle.Width)
	height := math.Abs(rectangle.Height)

	return width * height
}

func (circle Circle) Perimeter() float64 {
	if circle.Radius <= 0 {
		return -1.0
	}

	return 2 * math.Pi * circle.Radius
}

func (circle Circle) Area() float64 {
	if circle.Radius <= 0 {
		return -1.0
	}

	return math.Pi * circle.Radius * circle.Radius
}

func (triangle Triangle) Perimeter() float64 {
	if triangle.Base <= 0 || triangle.Height <= 0 {
		return -1.0
	}

	if triangle.Base == triangle.Height {
		return 3 * triangle.Base
	}

	hypothenuse := math.Hypot(triangle.Base, triangle.Height)
	return hypothenuse + triangle.Base + triangle.Height
}

func (triangle Triangle) Area() float64 {
	if triangle.Base <= 0 || triangle.Height <= 0 {
		return -1.0
	}

	return 0.5 * triangle.Base * triangle.Height
}
