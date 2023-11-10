package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := routes()
	fmt.Println("listening to port 8080 :")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
	}

}
