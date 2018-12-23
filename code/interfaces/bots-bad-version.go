package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreetingSpanish(sb)

}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func OldrintGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSpanish(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
