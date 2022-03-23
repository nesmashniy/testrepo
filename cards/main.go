package main

func main() {
	cards := deck{"Ace of Diamonds", NewCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func NewCard() string {
	return "Five of Diamonds"
}
