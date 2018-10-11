package recipes

import (
	Recipes "Projects/TruoraTest/src/controllers/recipesV1/models/recipes"
	DB "Projects/TruoraTest/src/systems/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var recipe Recipes.Recipe

	recipe.Name = r.PostFormValue("name")

	recipe.Description = r.PostFormValue("description")

	if r.PostFormValue("oven") == "true"{
		recipe.Oven = 0
	}else {
		recipe.Oven = 1
	}

	i, err := strconv.ParseInt(r.PostFormValue("time"), 10, 64)
	if err != nil {
		panic(err)
	}
	recipe.Time = i

	i, err = strconv.ParseInt(r.PostFormValue("noPersons"), 10, 64)
	if err != nil {
		panic(err)
	}
	recipe.NoPersons = i

	packet, err := json.Marshal(recipe)
	if err = DB.Store(db, recipe); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store users", http.StatusInternalServerError)
		return
	}

	packet, err = json.Marshal(recipe)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
