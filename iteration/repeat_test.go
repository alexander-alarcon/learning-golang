package iteration

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestRepeat(t *testing.T) {
	tests := []testhelpers.TestCase[RepeatOptions, string]{
		{
			Name:     "repeat_a_character_3_times",
			Input:    RepeatOptions{Character: "a", Times: 3},
			Expected: "aaa",
		},
		{
			Name:     "repeat_character_0_times",
			Input:    RepeatOptions{Character: "a", Times: 0},
			Expected: "",
		},
		{
			Name:     "repeat_empty_string_multiple_times",
			Input:    RepeatOptions{Character: "", Times: 100000000},
			Expected: "",
		},
		{
			Name:     "repeat_unicode_character",
			Input:    RepeatOptions{Character: "🙂", Times: 2},
			Expected: "🙂🙂",
		},
		{
			Name:     "repeat_character_1_time",
			Input:    RepeatOptions{Character: "z", Times: 1},
			Expected: "z",
		},
		{
			Name:     "repeat_negative_times",
			Input:    RepeatOptions{Character: "a", Times: -100},
			Expected: "",
		},
	}

	testhelpers.RunTableTest(t, tests, Repeat, testhelpers.AssertEqual[string])
}
