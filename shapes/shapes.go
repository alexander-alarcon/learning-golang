// Package shapes provides geometric shape types and functions
// to calculate their perimeter and area, with basic validation.
package shapes

import (
	"errors"
	"math"
)

// Shape represents any geometric shape that can compute
// its perimeter and area.
type Shape interface {
	Perimeter() (float64, error)
	Area() (float64, error)
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle represents a circle with a given radius.
type Circle struct {
	Radius float64
}

// Triangle represents a triangle using its three side lengths.
type Triangle struct {
	A float64
	B float64
	C float64
}

// Common validation errors returned by shape operations.
var (
	ErrInvalidDimension = errors.New("invalid dimension")
	ErrInvalidTriangle  = errors.New("invalid triangle sides")
)

// Perimeter returns the perimeter of the rectangle.
// Returns an error if any dimension is zero or negative.
func (rectangle Rectangle) Perimeter() (float64, error) {
	if rectangle.Width <= 0 || rectangle.Height <= 0 {
		return 0, ErrInvalidDimension
	}

	width := rectangle.Width
	height := rectangle.Height

	return 2 * (width + height), nil
}

// Area returns the area of the rectangle.
// Returns an error if any dimension is zero or negative.
func (rectangle Rectangle) Area() (float64, error) {
	if rectangle.Width <= 0 || rectangle.Height <= 0 ||
		math.IsNaN(rectangle.Width) || math.IsNaN(rectangle.Height) {
		return 0, ErrInvalidDimension
	}

	width := rectangle.Width
	height := rectangle.Height

	return width * height, nil
}

// Perimeter returns the circumference of the circle.
// Returns an error if the radius is zero or negative.
func (circle Circle) Perimeter() (float64, error) {
	if circle.Radius <= 0 {
		return 0, ErrInvalidDimension
	}
	return 2 * math.Pi * circle.Radius, nil
}

// Area returns the area of the circle.
// Returns an error if the radius is zero or negative.
func (circle Circle) Area() (float64, error) {
	if circle.Radius <= 0 || math.IsNaN(circle.Radius) || math.IsInf(circle.Radius, 0) {
		return 0, ErrInvalidDimension
	}

	return math.Pi * circle.Radius * circle.Radius, nil
}

// isValid checks whether the triangle has valid side lengths
// and satisfies the triangle inequality theorem.
func isValid(triangle Triangle) bool {
	a, b, c := triangle.A, triangle.B, triangle.C

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a+b <= c || a+c <= b || b+c <= a {
		return false
	}

	return true
}

// Area returns the area of the triangle using Heron's formula.
// Returns an error if the triangle is invalid.
func (triangle Triangle) Perimeter() (float64, error) {
	if !isValid(triangle) {
		return 0, ErrInvalidTriangle
	}

	return triangle.A + triangle.B + triangle.C, nil
}

// Perimeter returns the perimeter of the triangle.
// Returns an error if the triangle is invalid.
func (triangle Triangle) Area() (float64, error) {
	if !isValid(triangle) {
		return 0, ErrInvalidTriangle
	}

	semiPerimeter := (triangle.A + triangle.B + triangle.C) / 2
	area := math.Sqrt(
		semiPerimeter *
			(semiPerimeter - triangle.A) *
			(semiPerimeter - triangle.B) *
			(semiPerimeter - triangle.C),
	)

	return area, nil
}
