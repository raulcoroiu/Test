package main

import "fmt"

//orice tip catre trece prin functia getGreeting si returneaza un string poate folosi tipul bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type germanBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	gb := germanBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(gb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func (gb germanBot) getGreeting() string {
	return "Danke!"
}
