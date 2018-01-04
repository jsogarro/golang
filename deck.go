package main

import "fmt"

// Deck type 
type deck []string 

func createDeck() deck {
  cards:= deck{}
  suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
  values := []string{"Ace", "King", "Queen", "Jack", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
  
  for _, suit := range suits {
    for _, value := range values {
      cards = append(cards, value + " of " + suit)
    }
  }
  
  return cards
}

func dealCards(d deck, size int) (deck, deck) {
  return d[:size], d[size:]
}

func (d deck) printValues() {
  for _, card := range d {
    fmt.Println(card)
  }
}