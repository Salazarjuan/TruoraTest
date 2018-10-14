package router

import (
	"log"
	"net/http"

	"Projects/TruoraTest-server/pkg/types/routes"
	RecipeHandler "Projects/TruoraTest-server/src/controllers/recipesV1/routes/recipes"
	StepHandler "Projects/TruoraTest-server/src/controllers/recipesV1/routes/steps"
	UsersHandler "Projects/TruoraTest-server/src/controllers/recipesV1/routes/users"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*token := r.Header.Get("X-App-Token")

		if len(token) < 1 {
			log.Println(token)
			//log.Println(w)
			//log.Println(r)
			http.Error(w, "Not authorized into recipesV1", http.StatusUnauthorized)

			return
		}*/

		log.Println("Inside controller Middleware")

		next.ServeHTTP(w, r)
	})

}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB

	RecipeHandler.Init(DB)
	UsersHandler.Init(DB)
	StepHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/recipes": {
			Routes: routes.Routes{
				routes.Route{"RecipeIndex", "GET", "/", RecipeHandler.Index},
				routes.Route{"RecipeCreate", "POST", "/store", RecipeHandler.Store},
				routes.Route{"RecipeEdit", "GET", "/{id}/edit", RecipeHandler.Edit},
				routes.Route{"RecipeUpdate", "PUT", "/{id}", RecipeHandler.Update},
				routes.Route{"RecipeDestroy", "DELETE", "/{id}", RecipeHandler.Destroy},
			},
			Middleware: Middleware,
		},
		"/users": {
			Routes: routes.Routes{
				routes.Route{"userseIndex", "GET", "/", UsersHandler.Index},
				routes.Route{"usersCreate", "POST", "/store", UsersHandler.Store},
				routes.Route{"usersEdit", "GET", "/{id}/edit", UsersHandler.Edit},
				routes.Route{"usersUpdate", "PUT", "/{id}", UsersHandler.Update},
				routes.Route{"usersDestroy", "DELETE", "/{id}", UsersHandler.Destroy},
			},
			Middleware: Middleware,
		},
		"/steps": {
			Routes: routes.Routes{
				routes.Route{"stepIndex", "GET", "/", StepHandler.Index},
				routes.Route{"stepCreate", "POST", "/store", StepHandler.Store},
				routes.Route{"stepEdit", "GET", "/{id}/edit", StepHandler.Edit},
				routes.Route{"stepUpdate", "PUT", "/{id}", StepHandler.Update},
				routes.Route{"stepDestroy", "DELETE", "/{id}", StepHandler.Destroy},
			},
			Middleware: Middleware,
		},
	}

	return
}
