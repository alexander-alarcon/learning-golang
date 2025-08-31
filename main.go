package main

import (
	"fmt"
)

type Language int

const (
	English Language = iota
	Spanish
)

func (lang Language) String() string {
	names := [...]string{"English", "Spanish"}

	if int(lang) < 0 || int(lang) >= len(names) {
		return fmt.Sprintf("Unknown Language %d", lang)
	}

	return names[lang]
}

type GreetOptions struct {
	language Language
	name     string
}

const (
	englishGreeting = "Hello, "
	spanishGreeting = "Hola, "
)

func Greet(options GreetOptions) string {
	var greeting string

	switch options.language {
	case Spanish:
		greeting = spanishGreeting
	case English:
		greeting = englishGreeting
	default:
		greeting = englishGreeting
	}

	name := options.name
	if name == "" {
		name = "World"
	}

	return greeting + name
}

func main() {
	fmt.Println(Greet(GreetOptions{name: "Gopher"}))
	fmt.Println(Greet(GreetOptions{name: "Gopher", language: Spanish}))
}
