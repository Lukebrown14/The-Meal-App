package handlers

import (
	"net/http"
	"recipeapp/components"
	"recipeapp/utils"
)

func RandomRecipeHandler(w http.ResponseWriter, r *http.Request) {
	mealData, err := utils.FetchMealData("https://www.themealdb.com/api/json/v1/1/random.php")
	if err != nil {
		http.Error(w, "Failed to fetch meal data", http.StatusInternalServerError)
		return
	}

	if len(mealData.Meals) > 0 {
		err = components.RecipeList(&mealData.Meals[0]).Render(r.Context(), w)
	} else {
		err = components.RecipeList(nil).Render(r.Context(), w)
	}

	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
