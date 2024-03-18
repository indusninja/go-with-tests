package main

import "fmt"

const englishHelloExpression = "Hello"
const spanishHelloExpression = "Hola"
const frenchHelloExpression = "Bonjour"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return GetHelloExpressionInLanguage(language) + ", " + name
}

func GetHelloExpressionInLanguage(language string) string {
	prefix := englishHelloExpression
	switch language {
	case "spanish":
		prefix = spanishHelloExpression
	case "french":
		prefix = frenchHelloExpression
	}
	return prefix
}

func main() {
	fmt.Println(Hello("", "english"))
}
