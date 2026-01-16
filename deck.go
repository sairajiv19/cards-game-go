package main

import (
	"fmt"
	"strings"
)

type deck []string

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Printf("Card %d -> %s\n", i+1, card)
	}
}

func (d deck) Card(name string, pos int) { // reciever's arguments are only one letter matching the satarting of the type
	d[pos] = name
}

func newDeck() deck {
	var cardSuits = [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	var cardValues = [4]string{"Ace", "King", "Queen", "Jack"}
	var newd deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newd = append(newd, value+" of "+suit) // fmt.Sprintf can also be used but slow...
		}
	}
	return newd
}

func (d deck) Deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	var finalStr string = strings.Join(d, ", ")
	return finalStr
}
