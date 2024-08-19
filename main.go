package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	fmt.Println("after deal:")

	hand, remainingCards := deal(cards, 3)
	hand.print()
	fmt.Println("=====")
	remainingCards.print()

	fmt.Println("=====")
	fmt.Println(hand.toString())

	remainingCards.saveToFile("mycards")
	hand.saveToFile("mycards")

	cards = newDeckFromFile("mycards")
	cards.print()

	fmt.Println("=====")
	cards.shuffle()
	cards.print()
}
