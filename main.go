package main 
import "fmt"
func main() {
    cards := createDeck()
    hand, mainDeck := dealCards(cards, 2);
    
    hand.printValues()
    mainDeck.printValues()
    
    
    greeting := "Hello"
    fmt.Println([]byte(greeting))
}