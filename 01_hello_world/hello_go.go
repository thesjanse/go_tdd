package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := "Hello, "

	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	case "Latin":
		prefix = "Salve, "
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Jane", ""))
}
