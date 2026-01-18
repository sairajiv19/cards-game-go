package main

import (
	"fmt"
	"math/rand"
	"os"
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
	return strings.Join(d, ", ")
}

func (d deck) saveDeck(filename string) {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func loadDeck(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		os.Exit(1) // not recommended we should return err and if nothing then nil
	}
	return deck(strings.Split(string(data), ", "))
}

func (d deck) shuffleDeckRand() {
	n := len(d)
	for i := range n {
		j := rand.Intn(n)
		d[i], d[j] = d[j], d[i]
	}
}
