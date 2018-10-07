package routes

import (
	"net/http"
)

type Routes []Routes

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandleFunc
}

type SubRoutePackage struct {
	Routes      Routes
	Middleaware func(next http.Handler) http.Handler
}
