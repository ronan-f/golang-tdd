package main

import "fmt"

func main() {
	fmt.Println(hello("Ronan"))
}

func hello(name string) string {
	return "Hello" + " " + name
}
