package main

import "fmt"

const englishHelloPrefix = "Hello, "
const exclamationMark = "!"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + exclamationMark
}

func main() {
	fmt.Println(Hello("Jane"))
}
