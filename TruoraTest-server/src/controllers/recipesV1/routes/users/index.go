package users

import (
	Users "Projects/TruoraTest-server/src/controllers/recipesV1/models/users"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

var user Users.User

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

	users, err := Users.FindBy(db, limit, offset)

	packet, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}