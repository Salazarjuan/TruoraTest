package recipes

import (
	Recipes "Projects/TruoraTest-server/src/controllers/recipesV1/models/recipes"
	DB "Projects/TruoraTest-server/src/systems/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var recipe Recipes.Recipe
	r.ParseForm()

	name := r.PostFormValue("name")
	description := r.PostFormValue("description")
	oven := r.PostFormValue("oven")
	time := r.PostFormValue("time")
	noPersons := r.PostFormValue("noPersons")
	creatorID := r.PostFormValue("creatorID")

	log.Println(r)
	log.Println(name)
	log.Println(description)
	log.Println(oven)
	log.Println(time)
	log.Println(noPersons)

	recipe.Name = name

	recipe.Description = description

	recipe.Oven = oven

	log.Println("parsing time")
	i, err := strconv.ParseInt(time, 10, 64)
	if err != nil {
		panic(err)
	}
	recipe.Time = i

	log.Println("parsing noPersons")
	i, err = strconv.ParseInt(noPersons, 10, 64)
	if err != nil {
		panic(err)
	}
	recipe.NoPersons = i

	log.Println("parsing creatorID")
	i, err = strconv.ParseInt(creatorID, 10, 64)
	if err != nil {
		panic(err)
	}
	recipe.CreatorID = i

	packet, err := json.Marshal(recipe)
	if err = DB.Store(db, recipe); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store recipe", http.StatusInternalServerError)
		return
	}

	packet, err = json.Marshal(recipe)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse recipes", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
