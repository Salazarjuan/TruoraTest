package recipes

import "net/http"

func Destroy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies destroy"))
}
