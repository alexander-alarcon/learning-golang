package main

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []testhelpers.TestCase[GreetOptions, string]{
		{
			Name:     "default_name_and_language",
			Input:    GreetOptions{},
			Expected: "Hello, World",
		},
		{
			Name:     "greet_world_in_spanish",
			Input:    GreetOptions{language: Spanish},
			Expected: "Hola, Mundo",
		},
		{
			Name:     "greet_world_in_french",
			Input:    GreetOptions{language: French},
			Expected: "Bonjour, Monde",
		},
		{
			Name:     "greet_with_name_in_english",
			Input:    GreetOptions{name: "Gopher"},
			Expected: "Hello, Gopher",
		},
		{
			Name:     "greet_with_name_in_spanish",
			Input:    GreetOptions{name: "Gopher", language: Spanish},
			Expected: "Hola, Gopher",
		},
		{
			Name:     "greet_with_name_in_french",
			Input:    GreetOptions{name: "Gopher", language: French},
			Expected: "Bonjour, Gopher",
		},
	}

	testhelpers.RunTableTest(t, tests, Greet, testhelpers.AssertEqual[string])
}
