package main

import "fmt"

func main() {
	fmt.Println(hello("Ronan", ""))
}

const englishHelloPrefix = "Hello"

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return "Hola " + name
	}

	return englishHelloPrefix + " " + name
}
