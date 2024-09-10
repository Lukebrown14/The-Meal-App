package handlers

import (
	"log"
	"net/http"
	"recipeapp/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderHome(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
