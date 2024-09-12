package handlers

import (
	"net/http"
	"recipeapp/components"
)

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	err := components.LandingBody().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
