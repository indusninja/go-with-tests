package _1_helloworld

const englishHelloExpression = "Hello"
const spanishHelloExpression = "Hola"
const frenchHelloExpression = "Bonjour"

// Hello returns a personalised greeting, in a given language.
// If no name is provided, the greeting will default to the "World".
// If no language (or an unsupported language) is provided, the greeting will default to English.
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return GetHelloExpressionInLanguage(language) + ", " + name
}

// GetHelloExpressionInLanguage returns the hello expression in a given language.
// It currently supports English, Spanish and French.
// By default, the expression provided is in English.
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
