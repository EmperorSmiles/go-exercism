package airportrobot

import (
	"fmt"
)

// Write your code here.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}


func SayHello(name string, g Greeter) string {
	language := g.LanguageName()
	greeting := g.Greet(name)
	return fmt.Sprintf("I can speak %s: %s!", language, greeting)
}

type Italian struct {}
type Portuguese struct {}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s", name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s", name)
}
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

