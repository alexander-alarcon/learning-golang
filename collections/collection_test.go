package collections

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestCollections(t *testing.T) {
	t.Run("testing SumAll", func(t *testing.T) {
		tests := []testhelpers.TestCase[[][]int, []int]{
			{
				Name:     "sum_all_empty",
				Input:    [][]int{},
				Expected: []int{},
			},
			{
				Name:     "sum_all_non_empty_slices",
				Input:    [][]int{{1, 2}, {3, 4}, {5, 6}},
				Expected: []int{3, 7, 11},
			},
			{
				Name:     "sum_all_single_slices",
				Input:    [][]int{{10, 20, 30}},
				Expected: []int{60},
			},
			{
				Name:     "sum_all_with_empty_and_non_empty_slices",
				Input:    [][]int{{1, 2}, {3, 4}, {}, {5, 6}},
				Expected: []int{3, 7, 0, 11},
			},
			{
				Name:     "sum_all_with_negative_values",
				Input:    [][]int{{-1, -2}, {3, -3}, {0}},
				Expected: []int{-3, 0, 0},
			},
		}

		testhelpers.RunTableTest(t, tests, SumAll, testhelpers.AssertSlicesEqual[int])
	})

	t.Run("testing SumAllTails", func(t *testing.T) {
		tests := []testhelpers.TestCase[[][]int, []int]{
			{
				Name:     "sum_all_tails_empty_slices",
				Input:    [][]int{},
				Expected: []int{},
			},
			{
				Name:     "sum_all_tails_non_empty_slices",
				Input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				Expected: []int{5, 11, 17}, // sum of tails: 2+3, 5+6, 8+9
			},
			{
				Name:     "sum_all_tails_with_empty_and_single_element_slices",
				Input:    [][]int{{1}, {}, {10, 20}},
				Expected: []int{0, 0, 20}, // tails: empty, empty, 20
			},
		}

		testhelpers.RunTableTest(t, tests, SumAllTails, testhelpers.AssertSlicesEqual[int])
	})
}
