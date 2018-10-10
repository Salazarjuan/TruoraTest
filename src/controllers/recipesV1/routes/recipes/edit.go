package recipes

import "net/http"

func Edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies edit"))
}
