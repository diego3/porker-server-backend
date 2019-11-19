package main

import (
	"awesomeProject/core"
	"fmt"
	"net/http"
)


func main() {
	fmt.Println("Subindo servidor...")

	http.HandleFunc("/api/poker", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello")
	})

	http.HandleFunc("/api/poker/deck/init", core.HandlerDeck)

	er := http.ListenAndServe("localhost:8080", nil)
	if er != nil {
		fmt.Printf("Erro %s", er.Error())
	}

}
