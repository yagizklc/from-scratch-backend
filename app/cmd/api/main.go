package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var languages []Language

const (
	PORT = "8080"
	HOST = "localhost"
)

func main() {
	router := http.NewServeMux()

	languages = append(languages, Language{ID: "1", Name: "Go"})
	languages = append(languages, Language{ID: "2", Name: "Python"})

	router.HandleFunc("GET /languages", getLanguages)
	router.HandleFunc("GET /languages/{id}", getLanguage)
	router.HandleFunc("POST /languages", createLanguage)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), router))
}

func getLanguages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(languages)
}

func getLanguage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for _, item := range languages {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Language{})
}

func createLanguage(w http.ResponseWriter, r *http.Request) {
	var language Language
	_ = json.NewDecoder(r.Body).Decode(&language)
	languages = append(languages, language)
	json.NewEncoder(w).Encode(language)
}
