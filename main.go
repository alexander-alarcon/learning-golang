package main

import (
	"fmt"
)

type Language int

const (
	English Language = iota
	Spanish
	French
)

func (lang Language) String() string {
	names := [...]string{"English", "Spanish", "French"}

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
	frenchGreeting  = "Bonjour, "
)

func Greet(options GreetOptions) string {
	var greeting string

	switch options.language {
	case Spanish:
		greeting = spanishGreeting
	case English:
		greeting = englishGreeting
	case French:
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}

	name := options.name
	if name == "" {
		switch options.language {
		case Spanish:
			name = "Mundo"
		case French:
			name = "Monde"
		default:
			name = "World"
		}
	}

	return greeting + name
}

func main() {
	fmt.Println(Greet(GreetOptions{name: "Gopher"}))
	fmt.Println(Greet(GreetOptions{name: "Gopher", language: Spanish}))
	fmt.Println(Greet(GreetOptions{name: "Gopher", language: French}))
}
