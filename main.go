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

// Fetch meal data from API and decode it
func fetchMealData(url string) (MealResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return MealResponse{}, err
	}
	defer resp.Body.Close()

	var mealData MealResponse
	err = json.NewDecoder(resp.Body).Decode(&mealData)
	if err != nil {
		return MealResponse{}, err
	}

	return mealData, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, tmpl, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	mealData, err := fetchMealData("https://www.themealdb.com/api/json/v1/1/search.php?s=Arrabiata")
	if err != nil {
		log.Fatal(err)
	}

	err = renderTemplate(w, "Meal_list", mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}

func randomRecipeHandler(w http.ResponseWriter, r *http.Request) {
	mealData, err := fetchMealData("https://www.themealdb.com/api/json/v1/1/random.php")
	if err != nil {
		log.Fatal(err)
	}

	err = renderTemplate(w, "Meal_list", mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}

func searchRecipeHandler(w http.ResponseWriter, r *http.Request) {
	mealName := r.PostFormValue("meal-name")
	mealData, err := fetchMealData("https://www.themealdb.com/api/json/v1/1/search.php?s=" + mealName)
	if err != nil {
		log.Fatal(err)
	}

	err = renderTemplate(w, "Meal_list", mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/random-meal/", randomRecipeHandler)
	http.HandleFunc("/search-meal/", searchRecipeHandler)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
