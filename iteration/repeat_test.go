package iteration

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		Name     string
		Input    RepeatOptions
		Expected string
	}{
		{
			"repeat_a_character_3_times",
			RepeatOptions{Character: "a", Times: 3},
			"aaa",
		},
		{
			"repeat_character_0_times",
			RepeatOptions{Character: "a", Times: 0},
			"",
		},
		{
			"repeat_empty_string_multiple_times",
			RepeatOptions{Character: "", Times: 100000000},
			"",
		},
		{
			"repeat_unicode_character",
			RepeatOptions{Character: "🙂", Times: 2},
			"🙂🙂",
		},
		{
			"repeat_character_1_time",
			RepeatOptions{Character: "z", Times: 1},
			"z",
		},
		{
			"repeat_negative_times",
			RepeatOptions{Character: "a", Times: -100},
			"",
		},
	}

	testhelpers.RunTableTest(t, tests, Repeat)
}
