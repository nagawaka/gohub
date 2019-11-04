package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"ngwk.org/test/app/starred"
)

type ApiVersion struct {
	Version string
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	m := ApiVersion{"1.0"}
	b, err := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func Init() {
	fmt.Printf("api Started\n")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/starred/{username}", starred.FetchStarred).Methods("GET")
	// router.HandleFunc("/tags", fetchTags).Methods("GET")
	// router.HandleFunc("/tags/{tag_id}", fetchTags).Methods("POST")
	// router.HandleFunc("/tags/{tag_id}", fetchTags).Methods("PUT")
	// router.HandleFunc("/tags/{tag_id}", fetchTags).Methods("DELETE")
	// router.HandleFunc("/repositories/tag/{tag_id}", fetchReposByTag).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
