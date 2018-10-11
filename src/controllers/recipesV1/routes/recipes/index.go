package recipes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Recipe struct {
	Id string `json:id`
	Name string `json: name`
	Description string `jason: description`
	Oven bool `json: oven`
	Time int `json: time`
	noPersons int `json: noPersons`
}

func Index(w http.ResponseWriter, r *http.Request) {
	recipe := Recipe{
		Id:"",
		Name: "recipe name",
		Description: "recipe description",
		Oven: true,
		Time: 50,
		noPersons: 4,
	}

	recipeJson, err := json.Marshal(recipe)
	if err != nil{
		fmt.Fprintf(w, "error: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(recipeJson)
}
