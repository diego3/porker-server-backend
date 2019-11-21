package core

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


type Card struct {
	Naipe 	string `json:"naipe"`
	Code 	string `json:"code"`
	Value 	int	   `json:"value"`
}

type Deck struct {
	stack Stack
	cards []Card
}

func (pm *Deck) Init() {
	naipes := []string{"C", "D", "H", "S"}// Inicias em inglês dos naipes, TODO pesquisar legendas
	for _, v := range naipes {
		for i := 2; i<=10; i++ {
			pm.cards = append(pm.cards, Card{v, strconv.Itoa(i), i})
		}
	}

	coringas := []string{"J", "Q", "K", "A"}
	for _, v := range naipes {
		corVal := 10
		for _, coringa := range coringas {
			// J = 11, Q = 12, K = 13, A = 14
			corVal++
			pm.cards = append(pm.cards, Card{v, coringa, corVal})
		}
	}
	fmt.Printf("Deck inicializado len: %d\n", len(pm.cards))

	pm.Shuffle()
}

//
func (pm *Deck) Shuffle() {
	var numbers []int
	for i:=1; i <= 52; i++ {
		numbers = append(numbers, i)
	}
	// array shuffle https://programming.guide/go/shuffle-slice-array.html
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	// add shuffled array to stack
	for _, val := range numbers {
		pm.stack.Push(val)
	}
}

// get qty of cards from the cards
func (pm *Deck) GetCards(qty int)  []Card {
	var cards []Card
	if pm.stack.Empty() {
		return cards
	}

	for i:=0; i < qty; i++ {
		position, _ := pm.stack.Pop()
		cards = append(cards, pm.cards[position-1])
	}
	return cards
}

func (pm *Deck) Print()  {
	for _, v := range pm.cards {
		//fmt.Printf("{n:%s c:%s val:%d}\n", v.naipe, v.code, v.value)
		fmt.Println(v)
	}
}


