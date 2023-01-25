package main

import "fmt"

const SPANISH_LANGUAGE = "Spanish"
const GERMAN_LANGUAGE = "German"

const ENGLISH_HELLO_PREFIX = "Hello, "
const SPANISH_HELLO_PREFIX = "Hola, "
const GERMAN_HELLO_PREFIX = "Hallo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := ENGLISH_HELLO_PREFIX

	switch language {
	case SPANISH_LANGUAGE:
		prefix = SPANISH_HELLO_PREFIX
	case GERMAN_LANGUAGE:
		prefix = GERMAN_HELLO_PREFIX
	}

	return fmt.Sprintf("%s%s", prefix, name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
