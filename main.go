package main

import "fmt"

type bot interface {
	getGreeting() string
}

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

// note because the receiver param is not being used, it's been removed
func (englishBot) getGreeting() string {
	return "Wotcha!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
