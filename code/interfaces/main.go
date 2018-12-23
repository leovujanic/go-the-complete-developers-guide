package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engBot struct{}

type spanBot struct{}

func main() {

	eb := engBot{}
	sb := spanBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb engBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
