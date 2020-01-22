package main

import "fmt"

func main() {
	fmt.Println(hello("Ronan"))
}

const englishHelloPrefix = "Hello"

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + " " + name
}
