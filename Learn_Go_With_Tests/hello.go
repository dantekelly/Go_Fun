package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(n string, l string) string {
	var helloPrefix string

	switch l {
	case "es":
		helloPrefix = spanishHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}

	if n == "" {
		n = "World"
	}

	return helloPrefix + n + "!"
}

func main() {
	fmt.Println(Hello("Dante", "en"))
}
