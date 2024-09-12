package handlers

import (
	"net/http"
	"recipeapp/components"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    err := components.Index(nil).Render(r.Context(), w)
    if err != nil {
        http.Error(w, "Failed to render template", http.StatusInternalServerError)
    }
}
