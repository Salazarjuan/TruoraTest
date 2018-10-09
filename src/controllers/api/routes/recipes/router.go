package router

import (
	"log"
	"net/http"

	"Projects/TruoraTest/pkg/types/routes"
	StatusHandler "Projects/TruoraTest/src/controllers/api/recipes"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V1 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB

	StatusHandler.Init(DB)
	UsersHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/api": {
			Routes: routes.Routes{
				routes.Route{"RecepieIndex", "GET", "/recipes", UsersHandler.Index},
				routes.Route{"RecepieCreate", "PUT", "/recipes/create", UsersHandler.Store},
				routes.Route{"RecepieEdit", "GET", "/recipes/{id}/edit", UsersHandler.Edit},
				routes.Route{"RecepieUpdate", "PUT", "/recipes/{id}", UsersHandler.Update},
				routes.Route{"RecepieDestroy", "DELETE", "/recipes/{id}", UsersHandler.Destroy},
			},
			Middleware: Middleware,
		},
	}

	return
}
