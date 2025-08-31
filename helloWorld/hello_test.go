package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		opts     GreetOptions
		expected string
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

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Greet(testCase.opts)

			assertCorrectMessage(t, result, testCase.expected, testCase.name)
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string, testName string) {
	t.Helper()
	if got != want {
		t.Errorf("For test case %q: got %q, want %q", testName, got, want)
	}
}
