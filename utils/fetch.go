package utils

import (
	"encoding/json"
	"net/http"
	"recipeapp/models"
)

func FetchMealData(url string) (models.MealResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.MealResponse{}, err
	}
	defer resp.Body.Close()

	var mealData models.MealResponse
	err = json.NewDecoder(resp.Body).Decode(&mealData)
	if err != nil {
		return models.MealResponse{}, err
	}
	return mealData, nil
}
