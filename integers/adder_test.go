package integers

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestAdder(t *testing.T) {
	tests := []struct {
		Name     string
		Input    AdderOptions
		Expected int
	}{
		{
			"empty_list_returns_zero",
			AdderOptions{Numbers: []int{}},
			0,
		},
		{
			"sum_of_two_numbers",
			AdderOptions{Numbers: []int{1, 2}},
			3,
		},
		{
			"sum_of_several_numbers",
			AdderOptions{Numbers: []int{1, 2, 3, 4, 5}},
			15,
		},
		{
			"sum_with_negative_numbers",
			AdderOptions{Numbers: []int{-1, -2, 3}},
			0,
		},
		{
			"single_element",
			AdderOptions{Numbers: []int{42}},
			42,
		},
		{
			"sum_with_zeros",
			AdderOptions{Numbers: []int{0, 0, 5}},
			5,
		},
	}

	testhelpers.RunTableTest(t, tests, Adder)
}
