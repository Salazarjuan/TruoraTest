package recipes

import (
	"encoding/json"
	"log"
	"net/http"
)

type store_struct struct {

	Recipe string
}

func Store(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var s store_struct
	err := decoder.Decode(&s)
	if err != nil {
		panic(err)
	}
	log.Println(s.Recipe)
}
