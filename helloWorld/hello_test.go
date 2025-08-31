package main

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		Name     string
		Input    GreetOptions
		Expected string
	}{
		{
			"default_name_and_language",
			GreetOptions{},
			"Hello, World",
		},
		{
			"greet_world_in_spanish",
			GreetOptions{language: Spanish},
			"Hola, Mundo",
		},
		{
			"greet_world_in_french",
			GreetOptions{language: French},
			"Bonjour, Monde",
		},
		{
			"greet_with_name_in_english",
			GreetOptions{name: "Gopher"},
			"Hello, Gopher",
		},
		{
			"greet_with_name_in_spanish",
			GreetOptions{name: "Gopher", language: Spanish},
			"Hola, Gopher",
		},
		{
			"greet_with_name_in_french",
			GreetOptions{name: "Gopher", language: French},
			"Bonjour, Gopher",
		},
	}

	testhelpers.RunTableTest(t, tests, Greet)
}
