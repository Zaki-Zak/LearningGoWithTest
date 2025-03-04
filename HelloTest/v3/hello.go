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
	if language == Spanish {
		return spanishHelloPrefix + name
	}
	if language == French {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world", ""))

}
