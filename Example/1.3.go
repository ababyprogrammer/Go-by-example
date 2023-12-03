package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/animal/", animalHandler)

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

func animalHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/animal/" {
		http.NotFound(w, r)
		return
	}

	animal := r.URL.Path[len("/animal/"):]
	fmt.Fprintf(w, "Welcome to the %s page!", animal)
}