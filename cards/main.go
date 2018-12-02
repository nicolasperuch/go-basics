package main

import "fmt"

func main(){
	cards := newDeck()
	//hand, remainingDeck := deal(cards, 3)
	fmt.Println("before shuffle")
	cards.print()
	cards.shuffle()
	fmt.Println("after shuffle")
	cards.print()
}
