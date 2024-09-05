package models

type MealResponse struct {
	Meals []Meal `json:"meals"`
}

type Meal struct {
	IdMeal          string `json:"idMeal"`
	StrMeal         string `json:"strMeal"`
	StrCategory     string `json:"strCategory"`
	StrArea         string `json:"strArea"`
	StrInstructions string `json:"strInstructions"`
}
