package rest

import (
	"awesomeProject/core"
	"fmt"
	"net/http"
)

var (
	// TODO map com route to method ex. {"/api/deck": core.ApiDeck}
	routes = []string{
		"/api",
		"/api/cards",
	}
)

func main() {
	fmt.Println("Subindo servidor: localhost:8080")

	for _, route := range routes {
		fmt.Printf("route: %s", route)
	}

	http.HandleFunc("/api/poker", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(resp, "Contador")
	})

	http.HandleFunc("/api/poker/deck/dealer", core.DealerCards)

	er := http.ListenAndServe("localhost:8080", nil)
	if er != nil {
		fmt.Printf("Erro %s", er.Error())
	}
}
