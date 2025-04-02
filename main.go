package main

import "fmt"

const portuguese = "Portuguese"
const spanish = "Spanish"
const french = "French"
const prefix = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "
const prefixPortuguese = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case french:
		return prefixFrench + name + "!"
	case spanish:
		return prefixSpanish + name + "!"
	case portuguese:
		return prefixPortuguese + name + "!"
	}
	
	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Julio", "Portuguese"))
}