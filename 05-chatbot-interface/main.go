package main

import "fmt"

// if a type implements `getGreeting`, it implements `bot`.
// type: interface.
type bot interface {
	getGreeting() string
}

// type: concrete.
// note: you can use interfaces as atts.
type englishBot struct{}

type spanishBot struct{}

func main() {
	_englishBot := englishBot{}
	_spanishBot := spanishBot{}
	printGreeting(_englishBot)
	printGreeting(_spanishBot)
}

func printGreeting(_bot bot) {
	fmt.Println(_bot.getGreeting())
}

func (bot englishBot) getGreeting() string {
	return "Hello"
}

func (bot spanishBot) getGreeting() string {
	return "Hola"
}
