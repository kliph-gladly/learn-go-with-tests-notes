package main

import "fmt"

const SpanishLanguage = "Spanish"
const GermanLanguage = "German"

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const GermanHelloPrefix = "Hallo, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case SpanishLanguage:
		prefix = SpanishHelloPrefix
	case GermanLanguage:
		prefix = GermanHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}

	return prefix
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return fmt.Sprintf("%s%s", prefix, name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
