package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create a new type of 'deck' which is slice of strings
type deck []string

// 1.) newDeck(): purpose: to create a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Daimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// 2.) func print(): purpose: loop through deck of cards and print all the name if cards
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

// 4.) func deal():
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// 5.) saveToFile()
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// create new deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1: log the error and return a call to newDeck()
		// option #2: log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
