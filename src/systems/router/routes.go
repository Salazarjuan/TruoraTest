package router

import (
	"Projects/TruoraTest/pkg/types/routes"

	"github.com/go-xorm/xorm"

	//AuthHandler "Projects/TruoraTest/src/controllers/auth"
	"log"
	"net/http"

	HomeHandler "Projects/TruoraTest/src/controllers/home"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	//AuthHandler.Init(db)
	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		//routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		//routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
