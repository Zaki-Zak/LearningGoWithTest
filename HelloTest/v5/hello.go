package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	Spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	French             = "French"
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return

}
func main() {
	fmt.Println(Hello("world", ""))

}
