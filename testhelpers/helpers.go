// Package testhelpers provides generic testing helpers for table-driven tests.
package testhelpers

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T, testName string) {
	t.Helper()
	if got != want {
		t.Errorf("For test case %q: got %v, want %v", testName, got, want)
	}
}

func RunTableTest[Input any, Expected comparable](t *testing.T, tests []struct {
	Name     string
	Input    Input
	Expected Expected
}, f func(Input) Expected,
) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := f(test.Input)

			AssertEqual(t, got, test.Expected, test.Name)
		})
	}
}
