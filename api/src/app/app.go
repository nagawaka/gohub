package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home1!")
	fmt.Printf("api Started\n")
}

func main() {
	fmt.Printf("api Started\n")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}