package steps

import (
	Step "Projects/TruoraTest-server/src/controllers/recipesV1/models/steps"
	DB "Projects/TruoraTest-server/src/systems/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var step Step.Step

	step.Step = r.PostFormValue("step")
	recipeId, err := strconv.ParseInt(r.PostFormValue("recipeID"), 10, 64)


	step.RecipeID = recipeId
	log.Println(step.Step)
	log.Println(step.RecipeID)

	log.Println("1")
	log.Println(step)
	if err = DB.Store(db, step); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store steps", http.StatusInternalServerError)
		return
	}

	log.Println("2")
	packet, err := json.Marshal(step)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse steps", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
