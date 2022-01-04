package main

import "fmt"

type bot interface {
	getGreeting() string
	//getGreeting(string, int) (string, error)
	//getBotVersion() float
	//respondToUser(user) string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb) //output-> Hi there!
	printGreeting(sb) //output-> Hola!
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	//
	return "Hola!"
}
