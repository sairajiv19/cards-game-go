package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := cards.Deal(3)
	// hand.printDeck()
	// fmt.Println()
	// remainingCards.printDeck()
	fmt.Println(hand.toString())
	fmt.Println()
	fmt.Println(remainingCards.toString())
	fmt.Println("\nDone")
}
