package shapes

import (
	testhelpers "LearningGo/testhelpers"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tests := []testhelpers.TestCase[PerimeterOptions, float64]{
		{
			Name:     "rectangle_perimeter",
			Input:    PerimeterOptions{Width: 3.0, Height: 4.0},
			Expected: 14.0,
		},
		{
			Name:     "square_perimeter",
			Input:    PerimeterOptions{Width: 5.0, Height: 5.0},
			Expected: 20.0, // 2*(5 + 5) = 20
		},
		{
			Name:     "negative_width",
			Input:    PerimeterOptions{Width: -3.0, Height: 4.0},
			Expected: 14.0, // El valor negativo del ancho se convierte en positivo, 2*(3 + 4) = 14
		},
		{
			Name:     "negative_height",
			Input:    PerimeterOptions{Width: 3.0, Height: -4.0},
			Expected: 14.0, // El valor negativo de la altura se convierte en positivo, 2*(3 + 4) = 14
		},
		{
			Name:     "negative_both_dimensions",
			Input:    PerimeterOptions{Width: -3.0, Height: -4.0},
			Expected: 14.0, // Ambos valores negativos se convierten en positivos, 2*(3 + 4) = 14
		},
		{
			Name:     "large_dimensions",
			Input:    PerimeterOptions{Width: 1e6, Height: 1e6},
			Expected: 4e6, // 2*(1e6 + 1e6) = 4e6
		},
		{
			Name:     "huge_dimensions",
			Input:    PerimeterOptions{Width: 1e9, Height: 1e9},
			Expected: 4e9, // 2*(1e9 + 1e9) = 4e9
		},
		// Edge cases para valores pequeños
		{
			Name:     "small_dimensions",
			Input:    PerimeterOptions{Width: 1e-6, Height: 1e-6},
			Expected: 4e-6, // 2*(1e-6 + 1e-6) = 4e-6
		},
		{
			Name:     "tiny_dimensions",
			Input:    PerimeterOptions{Width: 1e-9, Height: 1e-9},
			Expected: 4e-9, // 2*(1e-9 + 1e-9) = 4e-9
		},
		// Casos con cero (deberían devolver -1.0)
		{
			Name:     "zero_width",
			Input:    PerimeterOptions{Width: 0.0, Height: 5.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "zero_height",
			Input:    PerimeterOptions{Width: 5.0, Height: 0.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
		{
			Name:     "zero_both_dimensions",
			Input:    PerimeterOptions{Width: 0.0, Height: 0.0},
			Expected: -1.0, // No es un rectángulo válido, devuelve -1.0
		},
	}

	testhelpers.RunTableTest(t, tests, Perimeter, testhelpers.AssertEqual[float64])
}
