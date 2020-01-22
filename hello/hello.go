package main

import "fmt"

func main() {
	fmt.Println(hello("Ronan", ""))
}

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + " " + name
	}

	if language == french {
		return frenchHelloPrefix + " " + name
	}

	return englishHelloPrefix + " " + name
}
