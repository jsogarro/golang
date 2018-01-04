package main

import (
  "fmt"
  "io/ioutil"
  "math/rand"
  "os"
  "strings"
  "time"
)

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

func (d deck) str() string {
  ss := []string(d)
  return strings.Join(ss, ",")
}

func (d deck) saveFile(file string) error {
  return ioutil.WriteFile(file, []byte(d.str()), 0777)
}

func readFileAndCreateDeck(file string) deck {
  bs, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Println("READ FILE ERROR: ", err)
    os.Exit(1)
  }
  
  s := strings.Split(string(bs), ",")
  return deck(s)
}

func (d deck) shuffleDeck() {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)
  
  for i := range d {
    newIndex := r.Intn(len(d) - 1)
    
    d[i], d[newIndex] = d[newIndex], d[i]
  }
}

func (d deck) printValues() {
  for _, card := range d {
    fmt.Println(card, ",")
  }
}