package integers

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestAdder(t *testing.T) {
	tests := []testhelpers.TestCase[AdderOptions, int]{
		{
			Name:     "empty_list_returns_zero",
			Input:    AdderOptions{Numbers: []int{}},
			Expected: 0,
		},
		{
			Name:     "sum_of_two_numbers",
			Input:    AdderOptions{Numbers: []int{1, 2}},
			Expected: 3,
		},
		{
			Name:     "sum_of_several_numbers",
			Input:    AdderOptions{Numbers: []int{1, 2, 3, 4, 5}},
			Expected: 15,
		},
		{
			Name:     "sum_with_negative_numbers",
			Input:    AdderOptions{Numbers: []int{-1, -2, 3}},
			Expected: 0,
		},
		{
			Name:     "single_element",
			Input:    AdderOptions{Numbers: []int{42}},
			Expected: 42,
		},
		{
			Name:     "sum_with_zeros",
			Input:    AdderOptions{Numbers: []int{0, 0, 5}},
			Expected: 5,
		},
	}

	testhelpers.RunTableTest(t, tests, Adder, testhelpers.AssertEqual[int])
}
