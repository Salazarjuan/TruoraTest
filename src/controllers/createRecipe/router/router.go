package router

import (
	"Projects/TruoraTest/pkg/types/routes"

	//AuthHandler "Projects/TruoraTest/src/controllers/auth"
	"log"
	"net/http"

	HomeHandler "Projects/TruoraTest/src/controllers/create"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside createRecipe middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		//routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		//routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
