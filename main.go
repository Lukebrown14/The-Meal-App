package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type MealResponse struct {
	Meals []struct {
		IdMeal          string `json:"idMeal"`
		StrMeal         string `json:"strMeal"`
		StrCategory     string `json:"strCategory"`
		StrArea         string `json:"strArea"`
		StrInstructions string `json:"strInstructions"`
	} `json:"meals"`
}

func mealHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch data from the URL
	resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/lookup.php?i=52772")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var mealData MealResponse
	err = json.NewDecoder(resp.Body).Decode(&mealData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create and parse the template
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the meal data
	err = tmpl.Execute(w, mealData.Meals[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", mealHandler)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
