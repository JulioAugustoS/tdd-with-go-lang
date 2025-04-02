package main

import "fmt"

const portuguese = "Portuguese"
const spanish = "Spanish"
const french = "French"
const prefixEnglish = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "
const prefixPortuguese = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixSayHello(language) + name + "!"
}

func prefixSayHello(language string) (prefix string) {
	switch language {
	case french:
		prefix = prefixFrench
	case spanish:
		prefix = prefixSpanish
	case portuguese:
		prefix = prefixPortuguese
	default:
		prefix = prefixEnglish
	}

	return
}

func main() {
	fmt.Println(Hello("Julio", "Portuguese"))
}