package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aashish32/websocketchat/internal/handlers"
)

func main() {

	mux := routes()
	log.Println("Starting channel Listener")
	go handlers.ListenToWsChannel()
	fmt.Println("listening to port 8080 :")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
	}

}
