package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Deck []string

func newDeck() Deck {
	deck := Deck{}

	cardSuits := []string{"Spades", "Diamons", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, value+" of "+suit)
		}
	}
	return deck
}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (deck Deck) deal(handSize int) (Deck, Deck) {
	return deck[:handSize], deck[handSize:]
}

func (deck Deck) saveToFile(filename string) bool {
	// Deprecated
	// ioutil.WriteFile(filename, []byte(deck.toString()), 0666)

	error := os.WriteFile(filename, []byte(deck.toString()), 0644)
	if error != nil {
		fmt.Println("Error:", error)
		return false
	}

	return true
}

func (deck Deck) shuffle() {
	for i := range deck {
		newIndex := rand.Intn(len(deck) - 1)
		deck[i], deck[newIndex] = deck[newIndex], deck[i]
	}
}

func newDeckFromFile(filename string) Deck {
	bs, error := os.ReadFile(filename)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	content := strings.Split(string(bs), ",")

	return Deck(content)
}

func (deck Deck) toString() string {
	return strings.Join([]string(deck), ",")
}
