package core

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)


type Card struct {
	naipe 	string
	code 	string
	value 	int
}

type PokerManager struct {
	stack 		 Stack
	deck 		 []Card
}

func (pm *PokerManager) InitDeck() {
	naipes := []string{"C", "D", "H", "S"}// Inicias em inglês dos naipes, TODO pesquisar legendas
	for _, v := range naipes {
		for i := 2; i<=10; i++ {
			pm.deck = append(pm.deck, Card{v, strconv.Itoa(i), i})
		}
	}

	coringas := []string{"J", "Q", "K", "A"}
	for _, v := range naipes {
		corVal := 10
		for _, coringa := range coringas {
			// J = 11, Q = 12, K = 13, A = 14
			corVal++
			pm.deck = append(pm.deck, Card{v, coringa, corVal})
		}
	}
	fmt.Printf("Deck inicializado len: %d\n", len(pm.deck))
}

//
func (pm *PokerManager) ShuffleDeck() {
	var numbers []int
	for i:=1; i <= 52; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	// array shuffle https://programming.guide/go/shuffle-slice-array.html
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Println("Shuffed numbers")
	fmt.Println(numbers)

	// add shuffled array to stack
	for _, val := range numbers {
		pm.stack.Push(val)
	}
}

// get qty of cards from the deck
func (pm *PokerManager) GetCards(qty int)  []Card {
	var cards []Card
	if pm.stack.Empty() {
		return cards
	}

	for i:=0; i < qty; i++ {
		position, _ := pm.stack.Pop()
		cards = append(cards, pm.deck[position-1])
	}
	return cards
}

func (pm *PokerManager) PrintDeck()  {
	for _, v := range pm.deck {
		fmt.Printf("{n:%s c:%s val:%d}\n", v.naipe, v.code, v.value)
	}
}


func HandlerDeck(resp http.ResponseWriter, req *http.Request) {
	pm := PokerManager{}
	pm.InitDeck()
	pm.ShuffleDeck()

	fmt.Fprintf(resp, "Stack %v", pm.stack)
}