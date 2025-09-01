package shapes

import (
	testhelpers "LearningGo/testhelpers"
	"math"
	"testing"
)

func TestShapePerimeter(t *testing.T) {
	tests := []testhelpers.TestCase[Shape, float64]{
		{
			Name:     "rectangle_perimeter",
			Input:    Rectangle{Width: 3.0, Height: 4.0},
			Expected: 14.0,
		},
		{
			Name:     "square_perimeter",
			Input:    Rectangle{Width: 5.0, Height: 5.0},
			Expected: 20.0, // 2*(5 + 5) = 20
		},
		{
			Name:     "negative_width",
			Input:    Rectangle{Width: -3.0, Height: 4.0},
			Expected: 14.0, // El valor negativo del ancho se convierte en positivo, 2*(3 + 4) = 14
		},
		{
			Name:     "negative_height",
			Input:    Rectangle{Width: 3.0, Height: -4.0},
			Expected: 14.0, // El valor negativo de la altura se convierte en positivo, 2*(3 + 4) = 14
		},
		{
			Name:     "negative_both_dimensions",
			Input:    Rectangle{Width: -3.0, Height: -4.0},
			Expected: 14.0, // Ambos valores negativos se convierten en positivos, 2*(3 + 4) = 14
		},
		{
			Name:     "large_dimensions",
			Input:    Rectangle{Width: 1e6, Height: 1e6},
			Expected: 4e6, // 2*(1e6 + 1e6) = 4e6
		},
		{
			Name:     "huge_dimensions",
			Input:    Rectangle{Width: 1e9, Height: 1e9},
			Expected: 4e9, // 2*(1e9 + 1e9) = 4e9
		},
		// Edge cases para valores pequeños
		{
			Name:     "small_dimensions",
			Input:    Rectangle{Width: 1e-6, Height: 1e-6},
			Expected: 4e-6, // 2*(1e-6 + 1e-6) = 4e-6
		},
		{
			Name:     "tiny_dimensions",
			Input:    Rectangle{Width: 1e-9, Height: 1e-9},
			Expected: 4e-9, // 2*(1e-9 + 1e-9) = 4e-9
		},
		// Casos con cero (deberían devolver -1.0)
		{
			Name:     "zero_width",
			Input:    Rectangle{Width: 0.0, Height: 5.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "zero_height",
			Input:    Rectangle{Width: 5.0, Height: 0.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "zero_both_dimensions",
			Input:    Rectangle{Width: 0.0, Height: 0.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "circle_perimeter",
			Input:    Circle{Radius: 3.0},
			Expected: 18.84955592153876,
		},
		{
			Name:     "zero_radius",
			Input:    Circle{Radius: 0.0},
			Expected: -1.0, // No es una circunferencia valida, devuelve -1.0
		},
		{
			Name:     "negative_radius",
			Input:    Circle{Radius: -3.0},
			Expected: -1.0, // El radio negativo no es válido, devuelve -1.0
		},
		{
			Name:     "large_radius",
			Input:    Circle{Radius: 1e6},
			Expected: 2 * math.Pi * 1e6, // 2π * 1e6
		},
		{
			Name:     "huge_radius",
			Input:    Circle{Radius: 1e9},
			Expected: 2 * math.Pi * 1e9, // 2π * 1e9
		},
		{
			Name:     "tiny_radius",
			Input:    Circle{Radius: 1e-9},
			Expected: 2 * math.Pi * 1e-9, // 2π * 1e-9
		},
		{
			Name:     "small_radius",
			Input:    Circle{Radius: 1e-6},
			Expected: 2 * math.Pi * 1e-6, // 2π * 1e-6
		},
	}

	testhelpers.RunTableTest(t, tests, Shape.Perimeter, testhelpers.AssertFloatEqual)
}

func TestShapeArea(t *testing.T) {
	tests := []testhelpers.TestCase[Shape, float64]{
		{
			Name:     "rectangle_area",
			Input:    Rectangle{Width: 3.0, Height: 4.0},
			Expected: 12.0,
		},
		{
			Name:     "square_area",
			Input:    Rectangle{Width: 5.0, Height: 5.0},
			Expected: 25.0, // 5 * 5 = 25
		},
		{
			Name:     "negative_width",
			Input:    Rectangle{Width: -3.0, Height: 4.0},
			Expected: 12.0, // El valor negativo del ancho se convierte en positivo, 3 * 4 = 12
		},
		{
			Name:     "negative_height",
			Input:    Rectangle{Width: 3.0, Height: -4.0},
			Expected: 12.0, // El valor negativo de la altura se convierte en positivo, 3 * 4 = 12
		},
		{
			Name:     "negative_both_dimensions",
			Input:    Rectangle{Width: -3.0, Height: -4.0},
			Expected: 12.0, // Ambos valores negativos se convierten en positivos, 3 * 4 = 12
		},
		{
			Name:     "large_dimensions",
			Input:    Rectangle{Width: 1e6, Height: 1e6},
			Expected: 1e12, // 1e6 * 1e6 = 1e12
		},
		{
			Name:     "huge_dimensions",
			Input:    Rectangle{Width: 1e9, Height: 1e9},
			Expected: 1e18, // 1e9 * 1e9 = 1e18
		},
		// Edge cases para valores pequeños
		{
			Name:     "small_dimensions",
			Input:    Rectangle{Width: 1e-6, Height: 1e-6},
			Expected: 1e-12, // 1e-6 * 1e-6 = 1e-12
		},
		{
			Name:     "tiny_dimensions",
			Input:    Rectangle{Width: 1e-9, Height: 1e-9},
			Expected: 1e-18, // 1e-9 * 1e-9 = 1e-18
		},
		// Casos con cero (deberían devolver -1.0)
		{
			Name:     "zero_width",
			Input:    Rectangle{Width: 0.0, Height: 5.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "zero_height",
			Input:    Rectangle{Width: 5.0, Height: 0.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "zero_both_dimensions",
			Input:    Rectangle{Width: 0.0, Height: 0.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "circle_area",
			Input:    Circle{Radius: 3.0},
			Expected: 28.274333882308138,
		},
		{
			Name:     "zero_radius",
			Input:    Circle{Radius: 0.0},
			Expected: -1.0, // No es una circunferencia valida, devuelve -1.0
		},
		{
			Name:     "negative_radius",
			Input:    Circle{Radius: -3.0},
			Expected: -1.0, // El radio negativo no es válido, devuelve -1.0
		},
		{
			Name:     "large_radius",
			Input:    Circle{Radius: 1e6},
			Expected: math.Pi * 1e6 * 1e6, // π * (1e6)^2
		},
		{
			Name:     "huge_radius",
			Input:    Circle{Radius: 1e9},
			Expected: math.Pi * 1e9 * 1e9, // π * (1e9)^2
		},
		{
			Name:     "tiny_radius",
			Input:    Circle{Radius: 1e-9},
			Expected: math.Pi * 1e-9 * 1e-9, // π * (1e-9)^2
		},
		{
			Name:     "small_radius",
			Input:    Circle{Radius: 1e-6},
			Expected: math.Pi * 1e-6 * 1e-6, // π * (1e-6)^2
		},
	}

	testhelpers.RunTableTest(t, tests, Shape.Area, testhelpers.AssertFloatEqual)
}
