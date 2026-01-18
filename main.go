package main

import "fmt"

func main() {
	cards := newDeck()
	hand, _ := cards.Deal(3)
	hand.saveDeck("output.txt")
	newHand := loadDeck("output.txt")
	newHand.printDeck()
	fmt.Println("After Shuffle")
	newHand.shuffleDeckRand()
	newHand.printDeck()
}
