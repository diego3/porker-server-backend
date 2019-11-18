package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Carta struct {
	naipe string
	valor int
}

type PokerManager struct {
	deck 		 []Carta // TODO deve ser uma stack
	deckSorteado []Carta // cartas já sorteadas
}

func (pm *PokerManager) inicializaDeck() {
	naipes := []string{"S", "P", "O", "V"}// Espadas, Paus, Ouro, Valete
	for _, v := range naipes {
		fmt.Printf("Simbol %s", v)
		var i int
		for i=2; i<=13; i++ {
			naipe := v
			switch i {
			case 11:
				naipe = "J"
			case 12:
				naipe = "Q"
			case 13:
				naipe = "K"
			}

			pm.deck = append(pm.deck, Carta{naipe, i})
		}
	}
	fmt.Printf("Deck inicializado len: %d", len(pm.deck))
}

// todo, shuffle method ser chamado após a inicialização

// dar cartas deve ser com uma stack, só ir fazendo pop
func (pm *PokerManager) darCartas(qtde int)  []Carta {
	// aleatoriamente sorteia uma carta do deck
	rand.Seed(52)
	var cartas []Carta
	for i:=0; i < qtde; i++ {
		var numero int = rand.Int()
		fmt.Printf("Random number: %d", numero)
		cartas = append(cartas, pm.deck[numero])
	}
	return cartas
}

func handlerDeck(resp http.ResponseWriter, req *http.Request) {
	pm := PokerManager{}
	pm.inicializaDeck()

	fmt.Println(pm)

	fmt.Fprintf(resp, "Inicializando o deck")
}

func main() {
	fmt.Println("Subindo servidor...")

	http.HandleFunc("/api/poker", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello")
	})

	http.HandleFunc("/api/poker/deck/init", handlerDeck)

	er := http.ListenAndServe("localhost:8080", nil)
	if er != nil {
		fmt.Printf("Erro %s", er.Error())
	}
}
