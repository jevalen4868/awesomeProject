package main

import (
	"fmt"
)

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

type car struct {
	make string
	model string
}

func (c *car) print() {
	fmt.Println(c)
	fmt.Println(*c)

}

func print(c *car) {
	fmt.Println(c)
	fmt.Println(*c)
}

type bot interface {
	getGreeting() string
}

type engBot struct{}
type spanBot struct{}

func (engBot) getGreeting() string {
	return "SUP YO"
}

func (spanBot) getGreeting() string {
	return "Que tal"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	printGreeting(engBot{})
	printGreeting(spanBot{})

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

	hand, remainingCards := deal(card02, 5)

	hand.print()
	remainingCards.print()


	writeFileError := hand.saveToFile("./hand.txt")
	fmt.Println(writeFileError)

	cardsFile := newDeckFromFile("./hand.txt")
	cardsFile.print()

	fmt.Println("-------------------------------")

	cardsFile.shuffle()

	cardsFile.print()

	fmt.Println("-------------------------------")

	cardsFile.shuffle()

	cardsFile.print()

	fmt.Println("-------------------------------")

	cardsFile.shuffle()

	cardsFile.print()

	fmt.Println("-------------------------------")

	fordFocus := car{
		make: "Ford",
		model: "Focus",
	}

	nameBill := "bill"
	fmt.Println(*&nameBill)

	print(&fordFocus)
	fordFocus.print()

	makeModel := map[string]string {
		"Ford" : "Focus",
		//"Ford" : "F150",
	}

	fmt.Println(makeModel)
}

