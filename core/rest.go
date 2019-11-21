package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func DealerCards(resp http.ResponseWriter, req *http.Request) {
	deck := Deck{}
	deck.Init()

	cards := deck.GetCards(2)
	fmt.Printf("deck size %d", deck.stack.Size())

	bytes, error := json.Marshal(cards)
	if error != nil {
		fmt.Fprintf(resp, "Json marshal error %v", error)
		return
	}

	json := string(bytes)
	//resp.Write(bytes)
	_, _ = fmt.Fprintf(resp, "%v", json)

}

