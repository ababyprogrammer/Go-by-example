package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
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
	varMatch := regexp.MustCompile(`/animal/([a-zA-Z]+)/([a-zA-Z]+)`).FindStringSubmatch(r.URL.Path)
	if len(varMatch) == 0 {
		http.NotFound(w, r)
		return
	}

	animal := varMatch[1]
	action := varMatch[2]
	fmt.Fprintf(w, "Welcome to the %s page! Here, you can %s!", animal, action)
}