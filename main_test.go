package main

import "testing"

func TestGreetDefaultValue(t *testing.T) {
	got := Greet(GreetOptions{})
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreetWithName(t *testing.T) {
	got := Greet(GreetOptions{name: "Gopher"})
	want := "Hello, Gopher"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
