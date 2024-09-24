package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID    int
	Name  string
	Title string
}

var books = []Book{
	{ID: 1, Name: "Koori", Title: "Koori-saramago"},
	{ID: 2, Name: "Dota2", Title: "Dota-saramago"},
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var Name_h []string
	for _, b := range books {
		Name_h = append(Name_h, b.Name)
		json.NewEncoder(w).Encode(Name_h)
	}
}

func main() {
	http.HandleFunc("/books", getBooksHandler)

	port := "localhost:8080"
	fmt.Println("Server is listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
