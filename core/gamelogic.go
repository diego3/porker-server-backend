package core

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"
)


type Carta struct {
	naipe string
	codigo string
	valor int
}

type PokerManager struct {
	stack 		 Stack
	deck 		 []Carta
}

func (pm *PokerManager) inicializaDeck() {
	naipes := []string{"C", "D", "H", "S"}// Inicias em inglês dos naipes, TODO pesquisar legendas
	for _, v := range naipes {
		for i := 2; i<=10; i++ {
			pm.deck = append(pm.deck, Carta{v, string(i), i})
		}
	}

	for _, v := range naipes {
		for i := 11; i<=14; i++ {
			// J = 11, Q = 12, K = 13, A = 14
			pm.deck = append(pm.deck, Carta{v, string(i), i})
		}
	}
	fmt.Printf("Deck inicializado len: %d", len(pm.deck))
}

// todo, shuffle method ser chamado após a inicialização
func (pm *PokerManager) suffleDeck() {
	numbers := make([]int, 52)
	for i:=1; i <= 52; i++ {
		numbers = append(numbers, i)
	}

	// array shuffle
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[i], numbers[j]
	})

	// add shuffled array to stack
	for _, val := range numbers {
		pm.stack.Push(val)
	}
}

// dar cartas deve ser com uma stack, só ir fazendo pop
func (pm *PokerManager) darCartas(qtde int)  []Carta {
	var cartas []Carta
	if pm.stack.Empty() {
		return cartas
	}

	for i:=0; i < qtde; i++ {
		position, _ := pm.stack.Pop()
		cartas = append(cartas, pm.deck[position-1])
	}
	return cartas
}

func HandlerDeck(resp http.ResponseWriter, req *http.Request) {
	pm := PokerManager{}
	pm.inicializaDeck()

	fmt.Println(pm)

	fmt.Fprintf(resp, "Inicializando o deck")
}