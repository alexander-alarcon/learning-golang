package main

import "fmt"

type GreetOptions struct {
	name string
}

const englishGreeting = "Hello, "

func Greet(options GreetOptions) string {
	if options.name == "" {
		return englishGreeting + "World"
	}
	return fmt.Sprintf("%s%s", englishGreeting, options.name)
}

func main() {
	fmt.Println(Greet(GreetOptions{name: "Gopher"}))
}
