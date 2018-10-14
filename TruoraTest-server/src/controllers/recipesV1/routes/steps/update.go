package steps

import (
	Step "Projects/TruoraTest-server/src/controllers/recipesV1/models/steps"
	DB "Projects/TruoraTest-server/src/systems/db"

	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)
	step := Step.Step{Id: int64(newId)}
	step.Step = r.PostFormValue("steps")
	recipeId, err := strconv.ParseInt(r.PostFormValue("recipeID"), 10, 64)
	step.RecipeID = recipeId

	if err := DB.Update(db, step.Id, &step); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get steps", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(step)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse steps", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
