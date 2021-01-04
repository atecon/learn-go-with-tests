package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	// return fmt.Sprintf("Hello, %s", name) // also works
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	NAME := "Chris"
	fmt.Println(Hello(NAME))
}
