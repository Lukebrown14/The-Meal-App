package main

import (
	"log"
	"net/http"
	"recipeapp/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/random-meal/", handlers.RandomRecipeHandler)
	http.HandleFunc("/search-meal/", handlers.SearchRecipeHandler)
	http.HandleFunc("/body/", handlers.LandingHandler)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
