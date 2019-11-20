package main

import (
	"awesomeProject/core"
	"fmt"
)

func testStack() {
	deck := core.Stack{}
	deck.Push(10)
	deck.Push(20)
	deck.Push(30)
	fmt.Printf("pilha %v\n", deck)

	n, _ := deck.Pop()
	fmt.Printf("pop %d\n", n)
	fmt.Printf("pilha %v\n", deck)
}

func testGameLogic() {
	fmt.Println("Testando gamelogic")
	poker := core.PokerManager{}

	poker.InitDeck()
	poker.ShuffleDeck()
	//fmt.Println("printando deck")
	//poker.PrintDeck()

	fmt.Println("Tentando puxar 2 cartas do deck..")
	cards := poker.GetCards(2)
	fmt.Println(cards)
}

func main() {

	testGameLogic()
}