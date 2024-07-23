package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	prefix := ""
	language = strings.ToLower(language)

	if name == "" {
		name = "World"
	}

	if language == "spanish" {
		prefix = spanishHelloPrefix
	} else if language == "english" {
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
