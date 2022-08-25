package main

func main() {
	cards := Deck{"Ace of Diamonds", "Five of Diamonds"}
	cards = append(cards, "Six of Spades")
	cards.printCards()

}
