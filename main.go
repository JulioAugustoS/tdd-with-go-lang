package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const prefix = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

  if language == spanish {
		return prefixSpanish + name + "!"
	}

	if language == french {
		return prefixFrench + name + "!"
	}
	
	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Ellie", "French"))
}