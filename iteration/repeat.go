// Package iteration provides Repeat function
package iteration

type RepeatOptions struct {
	Character string
	Times     int
}

func Repeat(options RepeatOptions) string {
	if options.Times <= 0 || options.Character == "" {
		return ""
	}

	var repeated string

	for i := 0; i < options.Times; i++ {
		repeated += options.Character
	}

	return repeated
}
