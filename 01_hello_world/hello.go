package main

import "fmt"

const prefixEnglish = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Halo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return gettingPrefix(language) + name
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = prefixSpanish
	case "french":
		prefix = prefixFrench
	case "english":
		prefix = prefixEnglish
	default:
		prefix = prefixEnglish
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Manu", "spanish"))
}
