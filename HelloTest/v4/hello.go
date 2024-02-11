package main

import "fmt"

const englishHelloPrefix = "Hello, "
const Spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const French = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	Prefix := englishHelloPrefix
	switch language {
	case "French":
		Prefix = frenchHelloPrefix
	case "Spanish":
		Prefix = spanishHelloPrefix
	}
	return Prefix + name
}
func main() {
	fmt.Println(Hello("world", ""))

}
