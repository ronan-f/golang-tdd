package main

import "fmt"

func main() {
	fmt.Println(hello("Ronan", ""))
}

const spanish = "Spanish"
const french = "French"
const irish = "Irish"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const irishHelloPrefix = "Dia duit"

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return getGreetingPrefix(language) + " " + name
}

func getGreetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case irish:
		prefix = irishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
