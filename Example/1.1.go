package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/cat/", catHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Welcome to the home page!")
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the dog page!")
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the cat page!")
}