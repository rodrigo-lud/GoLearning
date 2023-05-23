package main

import "fmt"

type englishBot struct{}

func (englishBot) getGreeting() string {
	return ("Hello!")
}

type spanishBot struct{}

func (spanishBot) getGreeting() string {
	return ("Hola!")
}

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// fmt.Println(eb.getGreeting())
	// fmt.Println(sb.getGreeting())

	printGreeting(eb)
	printGreeting(sb)

}
