package handlers

import (
	"net/http"
	"recipeapp/utils"
)

func BodyHandler(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "body_area", nil)

}
