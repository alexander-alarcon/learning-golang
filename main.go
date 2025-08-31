package main

import "fmt"

type GreetOptions struct {
	name string
}

func Greet(options GreetOptions) string {
	if options.name == "" {
		return "Hello, World"
	}
	return fmt.Sprintf("Hello, %s", options.name)
}

func main() {
	fmt.Println(Greet(GreetOptions{name: "Gopher"}))
}
