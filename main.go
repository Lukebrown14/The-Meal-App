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
	resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/search.php?s=Arrabiata")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var mealData MealResponse
	err = json.NewDecoder(resp.Body).Decode(&mealData)
	if err != nil {
		log.Fatal(err)
	}

	// Create and parse the template
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the template with the meal data
	err = tmpl.Execute(w, mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}

func randomMealHander(w http.ResponseWriter, r *http.Request) {
	// Fetch data from the URL
	resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/random.php")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var mealData MealResponse
	err = json.NewDecoder(resp.Body).Decode(&mealData)
	if err != nil {
		log.Fatal(err)
	}

	// Create and parse the template
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the template with the meal data
	err = tmpl.ExecuteTemplate(w, "Meal_list", mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}

func searchMealHander(w http.ResponseWriter, r *http.Request) {
	mealName := r.PostFormValue("meal-name")

	resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/search.php?s=" + mealName)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Parse the JSON response
	var mealData MealResponse
	err = json.NewDecoder(resp.Body).Decode(&mealData)
	if err != nil {
		log.Fatal(err)
	}

	// Create and parse the template
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the template with the meal data
	err = tmpl.ExecuteTemplate(w, "Meal_list", mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/", mealHandler)
	http.HandleFunc("/random-meal/", randomMealHander)
	http.HandleFunc("/search-meal/", searchMealHander)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
