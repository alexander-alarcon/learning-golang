package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Greet(GreetOptions{name: "Gopher"})
		want := "Hello, Gopher"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Greet(GreetOptions{name: ""})
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Greet(GreetOptions{name: "Gopher", language: Spanish})
		want := "Hola, Gopher"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
