package main

import "fmt"

type bot interface {
	getGreeting() string
}

// English bot Struct
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for gen english getGreeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	// very custom Logic for gen spanish greetings
	return "Hola!"
}
