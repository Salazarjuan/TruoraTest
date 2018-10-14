package recipes

import "net/http"

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies update"))
}

