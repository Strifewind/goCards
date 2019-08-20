package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
}

func (d deck) print() {
	for i, card := range d{
		fmt.Println(i, card)
	}
}