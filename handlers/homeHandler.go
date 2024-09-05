package handlers

import (
	"log"
	"net/http"
	"recipeapp/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	mealData, err := utils.FetchMealData("https://www.themealdb.com/api/json/v1/1/search.php?s=Arrabiata")
	if err != nil {
		log.Fatal(err)
	}
	err = utils.RenderHome(w, mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}
