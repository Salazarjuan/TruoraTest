package router

import (
	"Projects/TruoraTest/pkg/types/routes"

	"github.com/go-xorm/xorm"

	//AuthHandler "Projects/TruoraTest/src/controllers/auth"
	"log"
	"net/http"

	RecipeHandler "Projects/TruoraTest/src/controllers/recipes"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	//AuthHandler.Init(db)
	RecipeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", RecipeHandler.Index},

	}
}
