package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", NewCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func NewCard() string {
	return "Five of Diamonds"
}
