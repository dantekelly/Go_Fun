package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(n string, l string) string {
	if n == "" {
		n = "World"
	}

	return greetingPrefix(l) + n + "!"
}

func greetingPrefix(l string) (p string) {
	switch l {
	case "fr":
		p = frenchHelloPrefix
	case "es":
		p = spanishHelloPrefix
	default:
		p = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Dante", "en"))
}
