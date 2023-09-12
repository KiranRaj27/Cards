package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card:= "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds",newCard()}
	cards := newDeck()
	cards.shuffle()
	// cards := newDeckFromFile("my_cards")
	// cards.saveToFile("my_cards")
	// cards = append(cards,"Six of Spades")
	// for i,card := range cards{
		// 	fmt.Println(i,card)
		// }
		
	cards.print()

	// hand,remainingDeck := deal(cards,5)
	// hand.print()
	fmt.Println("===========================")
	// remainingDeck.print()
}

// func newCard() string{
// 	return "Five of Diamonds"
// }