package main

import (
	"awesomeProject/core"
	"fmt"
)


func main() {

	deck := core.Stack{}
	deck.Push(10)
	deck.Push(20)
	deck.Push(30)
	fmt.Printf("pilha %v\n", deck)

	n, _ := deck.Pop()
	fmt.Printf("pop %d\n", n)
	fmt.Printf("pilha %v\n", deck)
}