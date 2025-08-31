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

func getGreeting(language Language) string {
	switch language {
	case Spanish:
		return spanishGreeting
	case French:
		return frenchGreeting
	case English:
		fallthrough
	default:
		return englishGreeting
	}
}

func getName(language Language, name string) string {
	if name != "" {
		return name
	}

	switch language {
	case Spanish:
		return "Mundo"
	case French:
		return "Monde"
	default:
		return "World"
	}
}

func Greet(options GreetOptions) string {
	greeting := getGreeting(options.language)
	name := getName(options.language, options.name)
	return greeting + name
}

func main() {
	fmt.Println(Greet(GreetOptions{name: "Gopher"}))
	fmt.Println(Greet(GreetOptions{name: "Gopher", language: Spanish}))
	fmt.Println(Greet(GreetOptions{name: "Gopher", language: French}))
}
