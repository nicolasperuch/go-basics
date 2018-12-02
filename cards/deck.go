package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck{
	cards := deck{}
	suits := buildSuits()
	values := buildValues()

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println("index: ", i, " card: ", card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return buildHand(d, handSize), buildRemainingDeck(d, handSize)
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] 
	}
}

func (d deck) nextCard() string {
	return d[0]
}

func (d deck) lastCard() string {
	return d[len(d) - 1]
}

func (d deck) size() int {
	return len(d)
}

func buildHand(d deck, handSize int) deck {
	return d[:handSize]
}

func buildRemainingDeck(d deck, handSize int) deck {
	return d[handSize:]
}

func buildSuits() []string {
	return []string {
		"Spades",
		"Hearts",
		"Diamonds",
		"Clubs"}
}

func buildValues() []string {
	return []string {
		"Ace",
		"Two",
		"Jack",
		"Queen",
		"King"}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if nil != err {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	deckAsSlice := strings.Split(string(bs), ",")
	return deck(deckAsSlice)
}