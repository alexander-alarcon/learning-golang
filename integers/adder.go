// Package integers provides an Adder function
package integers

type AdderOptions struct {
	Numbers []int
}

func Adder(options AdderOptions) int {
	sum := 0

	if len(options.Numbers) == 0 {
		return 0
	}

	for _, number := range options.Numbers {
		sum += number
	}

	return sum
}
