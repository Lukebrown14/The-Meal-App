package meal

type MealResponse struct {
	Meals []Meal `json:"meals"`
}

type Meal struct {
	IDMeal          string `json:"idMeal"`
	StrMeal         string `json:"strMeal"`
	StrCategory     string `json:"strCategory"`
	StrArea         string `json:"strArea"`
	StrInstructions string `json:"strInstructions"`
}
