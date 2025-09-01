// Package testhelpers provides generic testing helpers for table-driven tests.
package testhelpers

import (
	"math"
	"reflect"
	"testing"
)

type AssertionFunc[Expected any] func(t *testing.T, got, want Expected, testName string)

type TestCase[Input any, Expected any] struct {
	Name     string
	Input    Input
	Expected Expected
}

func AssertEqual[T comparable](t *testing.T, got, want T, testName string) {
	t.Helper()
	if got != want {
		t.Errorf("For test case %q: got %v, want %v", testName, got, want)
	}
}

func AssertFloatEqual(t *testing.T, got, want float64, testName string) {
	t.Helper()

	// Usamos una tolerancia relativa que se adapta al tamaño de los valores
	relativeTolerance := 1e-9 // Tolerancia relativa estándar

	// Si el valor esperado es muy pequeño, usamos una tolerancia absoluta para evitar errores
	// de punto flotante en números muy pequeños
	if math.Abs(want) < 1e-9 {
		// Si el valor esperado es muy pequeño, aplicamos una tolerancia absoluta
		absoluteTolerance := 1e-9
		if math.Abs(got-want) > absoluteTolerance {
			t.Errorf("For test case %q: got %.9f, want %.9f (tolerance: %.9f)", testName, got, want, absoluteTolerance)
		}
	} else {
		// Si el valor esperado no es pequeño, usamos la tolerancia relativa
		if math.Abs(got-want)/math.Abs(want) > relativeTolerance {
			t.Errorf("For test case %q: got %.9f, want %.9f (tolerance: %.9f)", testName, got, want, relativeTolerance)
		}
	}
}

func AssertSlicesEqual[T any](t *testing.T, got, want []T, testName string) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("For test case %q: length mismatch - got length %d, want length %d", testName, len(got), len(want))
		return
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("For test case %q: got %v, want %v", testName, got, want)
	}
}

func RunTableTest[Input any, Expected any](
	t *testing.T,
	tests []TestCase[Input, Expected],
	function func(Input) Expected,
	assert AssertionFunc[Expected],
) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := function(test.Input)

			assert(t, got, test.Expected, test.Name)
		})
	}
}
