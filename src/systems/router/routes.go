package router

import (
	"Projects/TruoraTest-server/pkg/types/routes"

	"github.com/go-xorm/xorm"

	//AuthHandler "Projects/TruoraTest-server/src/controllers/auth"
	"log"
	"net/http"
	HomeHandler "Projects/TruoraTest-server/src/controllers/home"
	RecipeHandler "Projects/TruoraTest-server/src/controllers/recipesV1/routes/recipes"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	RecipeHandler.Init(db)
	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},

	}
}
