package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of deck, slice of strings
type deck []string

type gameType string

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

func (d deck) toString() string {
	deckString := strings.Join(d, ",")
	return deckString
}

func (d deck) saveToFile(file string) error {
	return ioutil.WriteFile(file, []byte(d.toString()), 0666)
}

func newDeckFromFile(file string) deck {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		fmt.Println("ERROR:", e)
		os.Exit(-1)
		//return newDeck()
	}

	s := strings.Split(string(b), ",")
	fmt.Println("Read from file:", s)
	return deck(s)

}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() deck {
	nd := deck {}
	for i, _ := range d {
		newPosition := rand.Intn(len(d) - 1)
		nd[newPosition] = d[i]

		// d[i], d[newPosition] = d[newPosition], d[i]Yeah omgTeam,
	}
	return nd
}
