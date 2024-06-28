package main

import "fmt"

//we cannot intialize a variable outside a function
//we can intialize though

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card :=newCard()
	// cards:= []string{"Ace of Diamonds",newCard()}
	
	// cards = append(cards, "six of spades")
	// fmt.Println(cards)
	// fmt.Println(card)

	// for i, card:=range cards{
	// 	fmt.Println(i, card)
	// }
	fmt.Println("Hello!!")

	playingCards := newDeck()
	hand, remainingCards := deal(playingCards, 5)
	hand.print()
	remainingCards.print() 

	greeting := "Hi there"
	fmt.Println([]byte(greeting))
	playingCards.toString()

	playingCards.saveToFile("my_cards")

	newDeck := newDeckFromFile("my_cards")
	newDeck.print()
}

// func newCard() string{
// 	return "Five of Diamonds" 
// }

/*
static typed language - GO
Basic types:
bool - true,false
string
int 
float64


*/

/*
every elemet of array or slice should be of same data type
Array - fixed length
slice - array that can grow and shrink

*/