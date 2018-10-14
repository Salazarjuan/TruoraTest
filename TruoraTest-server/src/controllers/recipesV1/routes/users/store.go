package users

import (
	Users "Projects/TruoraTest-server/src/controllers/recipesV1/models/users"
	DB "Projects/TruoraTest-server/src/systems/db"
	Passwords "Projects/TruoraTest-server/src/systems/passwords"

	"encoding/json"
	"log"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var user Users.User

	user.First = r.PostFormValue("first")
	user.Last = r.PostFormValue("last")
	user.Email = r.PostFormValue("email")
	user.Password = r.PostFormValue("password")

	encryptedPassword, err := Passwords.Encrypt(user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to encrypt password", http.StatusInternalServerError)
		return
	}

	encryptedPassword = encryptedPassword

	if err = DB.Store(db, &user); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store users", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
