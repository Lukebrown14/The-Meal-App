package handlers

import (
	"log"
	"net/http"
	"recipeapp/utils"
)

func RandomRecipeHandler(w http.ResponseWriter, r *http.Request) {
	mealData, err := utils.FetchMealData("https://www.themealdb.com/api/json/v1/1/random.php")
	if err != nil {
		log.Fatal(err)
	}
	err = utils.RenderTemplate(w, "Meal_list", mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}
