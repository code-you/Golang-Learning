package main

import "fmt"

type Deck []string

func newCards() {

}

func (d Deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) shuffle() {

}

func (d Deck) saveToFile() {

}
