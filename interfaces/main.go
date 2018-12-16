package main

import "fmt"

type bot interface {
	helloMessage() string
}

type englishBot struct {}
type brazilianBot struct {}

func main() {
	
	eb := englishBot{}
	bb := brazilianBot{}

	sayHello(bb)
	sayHello(eb)
}

func sayHello(b bot){
	fmt.Println(b.helloMessage())
}

func (englishBot) helloMessage() string {
	return "Hello from English bot"
}


func (brazilianBot) helloMessage() string {
	return "kkk eaeman"
}

