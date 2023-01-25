package main

import "fmt"

const SPANISH_LANGUAGE = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	helloPrefix := englishHelloPrefix

	if language == SPANISH_LANGUAGE {
		helloPrefix = spanishHelloPrefix
	}

	return fmt.Sprintf("%s%s", helloPrefix, name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
