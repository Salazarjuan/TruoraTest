package recipes

import (
	Recipes "Projects/TruoraTest-server/src/controllers/recipesV1/models/recipes"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

var recipe Recipes.Recipe

func Index(w http.ResponseWriter, r *http.Request) {

	var offset int
	limit := int(25)

	if keys, ok := r.URL.Query()["limit"]; ok {
		limit, _ = strconv.Atoi(keys[0])
	}
	if keys, ok := r.URL.Query()["page"]; ok {
		page, _ := strconv.Atoi(keys[0])
		offset = (page - 1) * limit
	}

	recipes, err := Recipes.FindBy(db, limit, offset)

	packet, err := json.Marshal(recipes)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse recipes", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}


