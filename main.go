package main

import "fmt"

func fib() func() int {
	a, b:= 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func newCard() string {
	return "Five of Diamonds"
}

var deckSize = 52

func main() {
	var card string = "Ace of Spades"
	card1 := "Queen of Hearts"
	card3 := ""
	card4 := newCard()
	//cards := card1, card3
	fmt.Println(card, card1, card3, card4, deckSize)

	fmt.Println("HI")

	f := fib()
	for i := -1; i <= 10; i++ {
		fmt.Println(f(), f(), f(), f())
	}

	cards := deck{"Ace of Diamonds", newCard(), newCard()}
	cards = append(cards, "Five of Hearts")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)

	cards.print()

	card02 := newDeck()
	card02.print()
}

