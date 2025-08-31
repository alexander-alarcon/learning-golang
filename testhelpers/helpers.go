// Package testhelpers provides generic testing helpers for table-driven tests.
package testhelpers

import (
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
