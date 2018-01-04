package main 

import (
    "os"
    "testing"
)

func TestCreateDeck(t *testing.T) {
  d := createDeck()
  
  if len(d) != 52 {
    t.Errorf("Expected deck len to be 52, but recieved %v", len(d))
  }
  
  if d[0] != "Ace of Clubs" {
    t.Errorf("Expected first card to be Ace of Clubs but got %v", d[0])
  }
  
  if d[len(d) - 1] != "10 of Spades" {
    t.Errorf("Expected first card to be Ace of Clubs but got %v", d[len(d) - 1])
  }
}

func TestStr(t *testing.T) {
  // TODO: implement
}

func TestSaveFileAndReadFileAndCreateDeck(t *testing.T) {
  // cleanup test files
  os.Remove("_testfile")
  
  d := createDeck()
  d.saveFile("_testfile")
  
  dPrime := readFileAndCreateDeck("_testfile")
  
  if len(dPrime) != 52 {
    t.Errorf("Expected deck len to be 52, but recieved %v", len(d))
  }
}

func TestDealCards(t *testing.T) {
  // TODO: implement
}

func TestShuffleDeck(t *testing.T) {
  // TODO: implement
}

func TestPrintValues(t *testing.T) {
  // TODO: implement
}