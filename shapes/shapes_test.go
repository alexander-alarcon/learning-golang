package shapes

import (
	"LearningGo/testhelpers"
	"math"
	"testing"
)

func TestShapePerimeter(t *testing.T) {
	tests := []testhelpers.TestCase[Shape, float64]{
		{
			Name:          "rectangle_perimeter",
			Input:         Rectangle{Width: 3.0, Height: 4.0},
			Expected:      14.0,
			ExpectedError: nil,
		},
		{
			Name:          "square_perimeter",
			Input:         Rectangle{Width: 5.0, Height: 5.0},
			Expected:      20.0,
			ExpectedError: nil,
		},
		{
			Name:          "negative_width",
			Input:         Rectangle{Width: -3.0, Height: 4.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "negative_height",
			Input:         Rectangle{Width: 3.0, Height: -4.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "negative_both_dimensions",
			Input:         Rectangle{Width: -3.0, Height: -4.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "large_dimensions",
			Input:         Rectangle{Width: 1e6, Height: 1e6},
			Expected:      4e6,
			ExpectedError: nil,
		},
		{
			Name:          "huge_dimensions",
			Input:         Rectangle{Width: 1e9, Height: 1e9},
			Expected:      4e9,
			ExpectedError: nil,
		},
		{
			Name:          "small_dimensions",
			Input:         Rectangle{Width: 1e-6, Height: 1e-6},
			Expected:      4e-6,
			ExpectedError: nil,
		},
		{
			Name:          "tiny_dimensions",
			Input:         Rectangle{Width: 1e-9, Height: 1e-9},
			Expected:      4e-9,
			ExpectedError: nil,
		},
		{
			Name:          "zero_width",
			Input:         Rectangle{Width: 0.0, Height: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "zero_height",
			Input:         Rectangle{Width: 5.0, Height: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "zero_both_dimensions",
			Input:         Rectangle{Width: 0.0, Height: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:     "circle_perimeter",
			Input:    Circle{Radius: 3.0},
			Expected: 18.84955592153876,
		},
		{
			Name:          "zero_radius",
			Input:         Circle{Radius: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "negative_radius",
			Input:         Circle{Radius: -3.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "large_radius",
			Input:         Circle{Radius: 1e6},
			Expected:      2 * math.Pi * 1e6,
			ExpectedError: nil,
		},
		{
			Name:          "huge_radius",
			Input:         Circle{Radius: 1e9},
			Expected:      2 * math.Pi * 1e9,
			ExpectedError: nil,
		},
		{
			Name:          "tiny_radius",
			Input:         Circle{Radius: 1e-9},
			Expected:      2 * math.Pi * 1e-9,
			ExpectedError: nil,
		},
		{
			Name:          "small_radius",
			Input:         Circle{Radius: 1e-6},
			Expected:      2 * math.Pi * 1e-6,
			ExpectedError: nil,
		},
		{
			Name:          "right_triangle_perimeter",
			Input:         Triangle{A: 3.0, B: 4.0, C: 5.0},
			Expected:      12.0,
			ExpectedError: nil,
		},
		{
			Name:          "equilateral_triangle_perimeter",
			Input:         Triangle{A: 5.0, B: 5.0, C: 5.0},
			Expected:      15.0,
			ExpectedError: nil,
		},
		{
			Name:          "isosceles_triangle_perimeter",
			Input:         Triangle{A: 5.0, B: 5.0, C: 3.0},
			Expected:      13.0,
			ExpectedError: nil,
		},
		{
			Name:          "negative_A_side_perimeter",
			Input:         Triangle{A: -3.0, B: 4.0, C: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "negative_B_side_perimeter",
			Input:         Triangle{A: 3.0, B: -4.0, C: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "negative_C_side_perimeter",
			Input:         Triangle{A: 3.0, B: 4.0, C: -5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "zero_A_side_perimeter",
			Input:         Triangle{A: 0.0, B: 4.0, C: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "zero_B_side_perimeter",
			Input:         Triangle{A: 3.0, B: 0.0, C: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "zero_C_side_perimeter",
			Input:         Triangle{A: 3.0, B: 4.0, C: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "zero_all_sides_perimeter",
			Input:         Triangle{A: 0.0, B: 0.0, C: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "triangle_inequality_violation",
			Input:         Triangle{A: 1.0, B: 2.0, C: 10.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "triangle_degenerate_case",
			Input:         Triangle{A: 2.0, B: 3.0, C: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
	}

	testhelpers.RunTableTestWithError(t, tests, Shape.Perimeter, testhelpers.AssertFloatEqual, testhelpers.AssertEqual)
}

func TestShapeArea(t *testing.T) {
	tests := []testhelpers.TestCase[Shape, float64]{
		{
			Name:          "rectangle_area",
			Input:         Rectangle{Width: 3.0, Height: 4.0},
			Expected:      12.0,
			ExpectedError: nil,
		},
		{
			Name:          "square_area",
			Input:         Rectangle{Width: 5.0, Height: 5.0},
			Expected:      25.0,
			ExpectedError: nil,
		},
		{
			Name:          "negative_width",
			Input:         Rectangle{Width: -3.0, Height: 4.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "negative_height",
			Input:         Rectangle{Width: 3.0, Height: -4.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "negative_both_dimensions",
			Input:         Rectangle{Width: -3.0, Height: -4.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "large_dimensions",
			Input:         Rectangle{Width: 1e6, Height: 1e6},
			Expected:      1e12,
			ExpectedError: nil,
		},
		{
			Name:          "huge_dimensions",
			Input:         Rectangle{Width: 1e9, Height: 1e9},
			Expected:      1e18,
			ExpectedError: nil,
		},
		{
			Name:          "small_dimensions",
			Input:         Rectangle{Width: 1e-6, Height: 1e-6},
			Expected:      1e-12,
			ExpectedError: nil,
		},
		{
			Name:          "tiny_dimensions",
			Input:         Rectangle{Width: 1e-9, Height: 1e-9},
			Expected:      1e-18,
			ExpectedError: nil,
		},
		{
			Name:          "zero_width",
			Input:         Rectangle{Width: 0.0, Height: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "zero_height",
			Input:         Rectangle{Width: 5.0, Height: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "zero_both_dimensions",
			Input:         Rectangle{Width: 0.0, Height: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "decimal_dimensions",
			Input:         Rectangle{Width: 2.5, Height: 4.2},
			Expected:      10.5,
			ExpectedError: nil,
		},
		{
			Name:          "high_precision_dimensions",
			Input:         Rectangle{Width: 1.0000000001, Height: 1.0000000002},
			Expected:      1.0000000003,
			ExpectedError: nil,
		},
		{
			Name:          "infinite_width",
			Input:         Rectangle{Width: math.Inf(1), Height: 5.0},
			Expected:      math.Inf(1),
			ExpectedError: nil,
		},
		{
			Name:          "nan_height",
			Input:         Rectangle{Width: 5.0, Height: math.NaN()},
			Expected:      math.NaN(),
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "circle_area",
			Input:         Circle{Radius: 3.0},
			Expected:      28.274333882308138,
			ExpectedError: nil,
		},
		{
			Name:          "zero_radius",
			Input:         Circle{Radius: 0.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "negative_radius",
			Input:         Circle{Radius: -3.0},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "large_radius",
			Input:         Circle{Radius: 1e6},
			Expected:      math.Pi * 1e6 * 1e6,
			ExpectedError: nil,
		},
		{
			Name:          "huge_radius",
			Input:         Circle{Radius: 1e9},
			Expected:      math.Pi * 1e9 * 1e9,
			ExpectedError: nil,
		},
		{
			Name:          "tiny_radius",
			Input:         Circle{Radius: 1e-9},
			Expected:      math.Pi * 1e-9 * 1e-9,
			ExpectedError: nil,
		},
		{
			Name:          "small_radius",
			Input:         Circle{Radius: 1e-6},
			Expected:      math.Pi * 1e-6 * 1e-6,
			ExpectedError: nil,
		},
		{
			Name:          "nan_radius",
			Input:         Circle{Radius: math.NaN()},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "infinite_radius",
			Input:         Circle{Radius: math.Inf(1)},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "negative_infinite_radius",
			Input:         Circle{Radius: math.Inf(-1)},
			Expected:      0,
			ExpectedError: ErrInvalidDimension,
		},
		{
			Name:          "right_triangle_area",
			Input:         Triangle{A: 3.0, B: 4.0, C: 5.0},
			Expected:      6.0,
			ExpectedError: nil,
		},
		{
			Name:          "isosceles_triangle_area",
			Input:         Triangle{A: 5.0, B: 5.0, C: 6.0},
			Expected:      12.0,
			ExpectedError: nil,
		},
		{
			Name:          "equilateral_triangle_area",
			Input:         Triangle{A: 5.0, B: 5.0, C: 5.0},
			Expected:      10.825317547305486,
			ExpectedError: nil,
		},
		{
			Name:          "scalene_triangle_area",
			Input:         Triangle{A: 7.0, B: 5.0, C: 3.0},
			Expected:      6.49519052838329,
			ExpectedError: nil,
		},
		{
			Name:          "tiny_triangle_area",
			Input:         Triangle{A: 1e-9, B: 1e-9, C: 1e-9},
			Expected:      4.330127018922194e-19,
			ExpectedError: nil,
		},
		{
			Name:          "huge_triangle_area",
			Input:         Triangle{A: 1e6, B: 1e6, C: 1e6},
			Expected:      4.330127018922193e11,
			ExpectedError: nil,
		},
		{
			Name:          "zero_side_area",
			Input:         Triangle{A: 0.0, B: 4.0, C: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "negative_side_area",
			Input:         Triangle{A: -3.0, B: 4.0, C: 5.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "triangle_inequality_violation",
			Input:         Triangle{A: 10.0, B: 3.0, C: 4.0},
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
		{
			Name:          "degenerate_triangle_area",
			Input:         Triangle{A: 5.0, B: 5.0, C: 10.0}, // a + b == c
			Expected:      0,
			ExpectedError: ErrInvalidTriangle,
		},
	}

	testhelpers.RunTableTestWithError(t, tests, Shape.Area, testhelpers.AssertFloatEqual, testhelpers.AssertEqual)
}
