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
	deck := core.Deck{}

	deck.Init()


	fmt.Println("Tentando puxar 2 cartas do deck..")
	cards := deck.GetCards(2)
	fmt.Println(cards)

	fmt.Println("printando deck")
	deck.Print()

	fmt.Println("Tentando puxar 2 cartas do deck..")
	cards2 := deck.GetCards(2)
	fmt.Println(cards2)

	fmt.Println("printando deck")
	deck.Print()
}

func main() {

	testGameLogic()
}