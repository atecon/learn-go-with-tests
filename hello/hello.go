package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello creates a friendly 'hello'
func Hello(name string, language string) string {
	// return fmt.Sprintf("Hello, %s", name) // also works

	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	if language == "French" {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	NAME := "Chris"
	fmt.Println(Hello(NAME, ""))
}
