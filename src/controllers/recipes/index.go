package recipes

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies idex"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies update"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies create"))
}

func Edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies edit"))
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipies destroy"))
}
