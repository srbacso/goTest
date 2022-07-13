package main

import "fmt"

func main() {
	fmt.Println("Starting")

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//cards := newDeckFromFile("saved_deck.txt")

	//fmt.Println(cards.toString()
	//
	//cards.saveToFile("saved_deck.txt")

	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//remainingCards.print()

	//cards.print()

	fmt.Println("Done")
}
