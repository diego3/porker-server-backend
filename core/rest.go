package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var deck = Deck{}

func DealerCards(resp http.ResponseWriter, req *http.Request) {
	deck.Init()

	qty, err := strconv.Atoi(req.Form.Get("qty"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "qty param error: %v", err)
		return
	}

	cards := deck.GetCards(qty)
	fmt.Printf("deck size %d", deck.stack.Size())

	bytes, error := json.Marshal(cards)
	if error != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(resp, "Json marshal error %v", error)
		return
	}

	json := string(bytes)
	//resp.Write(bytes)
	_, _ = fmt.Fprintf(resp, "%v", json)
}

