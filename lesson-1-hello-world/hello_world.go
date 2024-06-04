package main

import "fmt"

const (
	Spanish = "Spanish"
	French  = "French"

	FrenchHelloPrefix  = "Bonjour, "
	SpanishHelloPrefix = "Hola, "
	EnglishHelloPrefix = "Hello, "
)

//One way to write a function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return resolveGreetingPrefix(language) + name
}

//Another way to write a function
func resolveGreetingPrefix(language string) (prefix string) {
	switch language {
	case French:
		prefix = FrenchHelloPrefix
	case Spanish:
		prefix = SpanishHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("my twin", ""))
}
