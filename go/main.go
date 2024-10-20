package main

import (
	"fmt"
	"log"
	"net/http"
)

const ADDR = ":8080"

func main() {
	fmt.Println("iniciando servidor")
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"message":"pong"}`)
	})

	fmt.Printf("servidor escutando em: %s\n", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, nil))
}
