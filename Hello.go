package main

import "fmt"

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
	fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}

func estimatePi() float64 {
	return 3.14
}
