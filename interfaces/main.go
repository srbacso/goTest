package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	botName string
}
type spanishBot struct {
	botName string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	eb.botName = "English"
	sb.botName = "Spanish"

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return eb.botName + ": Hi There"
}

func (sb spanishBot) getGreeting() string {
	return sb.botName + ": Hola!"
}
