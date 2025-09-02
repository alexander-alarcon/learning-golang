// Package testhelpers provides generic testing utilities for table-driven tests.
package testhelpers

import (
	"math"
	"reflect"
	"testing"
)

// AssertionFunc is a generic function type used to compare expected and actual results in tests.
type AssertionFunc[Expected any] func(t *testing.T, got, want Expected, testName string)

// TestCase represents a standard table-driven test case for pure functions.
type TestCase[Input any, Expected any] struct {
	Name     string   // Name of the test case
	Input    Input    // Input to the function under test
	Expected Expected // Expected result
}

// StateTestCase is used for tests that involve setting up mutable state.
// It allows testing methods or functions that rely on internal state or side effects.
type StateTestCase[Subject any, Result any] struct {
	Name     string                // Name of the test case
	Setup    func() Subject        // Initializes the subject (e.g., object or system under test)
	Test     func(Subject) Result  // Performs the test logic and returns the result
	Expected Result                // Expected result from the test logic
	Assert   AssertionFunc[Result] // Assertion function to compare actual vs expected
}

// StateTestCaseNoExpected is for state-based tests that don't return a value,
// and instead perform internal assertions.
type StateTestCaseNoExpected[Subject any] struct {
	Name   string                          // Name of the test case
	Setup  func() Subject                  // Initializes the subject
	Test   func(Subject)                   // Performs the test logic
	Assert func(t *testing.T, name string) // Custom assertion logic (manual)
}

// AssertEqual compares two values using == and fails the test if they differ.
func AssertEqual[T comparable](t *testing.T, got, want T, testName string) {
	t.Helper()
	if got != want {
		t.Errorf("For test case %q: got %v, want %v", testName, got, want)
	}
}

// AssertFloatEqual compares two float64 values using relative or absolute tolerance.
func AssertFloatEqual(t *testing.T, got, want float64, testName string) {
	t.Helper()

	relativeTolerance := 1e-9 // Standard relative tolerance

	// Use absolute tolerance for very small expected values
	if math.Abs(want) < 1e-9 {
		absoluteTolerance := 1e-9
		if math.Abs(got-want) > absoluteTolerance {
			t.Errorf("For test case %q: got %.9f, want %.9f (tolerance: %.9f)", testName, got, want, absoluteTolerance)
		}
	} else {
		if math.Abs(got-want)/math.Abs(want) > relativeTolerance {
			t.Errorf("For test case %q: got %.9f, want %.9f (tolerance: %.9f)", testName, got, want, relativeTolerance)
		}
	}
}

// AssertSlicesEqual compares two slices using reflect.DeepEqual and fails the test if they differ.
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

// RunTableTest executes a list of TestCase for pure functions with deterministic output.
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

// RunStateTests executes a list of StateTestCase for scenarios involving internal state or side effects.
func RunStateTests[Subject any, Result any](
	t *testing.T,
	tests []StateTestCase[Subject, Result],
) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			subject := test.Setup()
			result := test.Test(subject)
			test.Assert(t, result, test.Expected, test.Name)
		})
	}
}

// RunStateTestsNoExpected executes StateTestCaseNoExpected where no return value is expected.
// Assertions must be handled manually inside the Assert function.
func RunStateTestsNoExpected[Subject any](
	t *testing.T,
	tests []StateTestCaseNoExpected[Subject],
) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			subject := test.Setup()
			test.Test(subject)
			test.Assert(t, test.Name)
		})
	}
}
