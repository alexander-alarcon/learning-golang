// Package collections provides SumAll and SumAllTails functions
package collections

import "LearningGo/integers"

func SumAll(slices [][]int) []int {
	lengthOfNumbers := len(slices)

	if lengthOfNumbers == 0 {
		return []int{}
	}

	sums := make([]int, lengthOfNumbers)

	for index, list := range slices {
		sums[index] = integers.Adder(integers.AdderOptions{Numbers: list})
	}
	return sums
}

func SumAllTails(slices [][]int) []int {
	if len(slices) == 0 {
		return []int{}
	}

	var sums []int

	for _, list := range slices {
		if len(list) > 1 {
			tail := list[1:]
			sums = append(sums, integers.Adder(integers.AdderOptions{Numbers: tail}))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
