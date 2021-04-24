package main

import "fmt"

// Create a new type of deck, slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	return cards
}

func (d deck) copyDeck() deck {
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
