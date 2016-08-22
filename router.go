package main

import (
	"github.com/gorilla/pat"
	"github.com/markbates/goth/gothic"
)

func NewRouter() *pat.Router {

	p := pat.New()
	p.Get("/auth/{provider}", gothic.BeginAuthHandler)
	p.Get("/auth/{provider}/callback", AuthCallback)

	// for _, route := range routes {
	// 	var handler http.Handler
	//
	// 	handler = route.HandlerFunc
	// 	handler = Logger(handler, route.Name)
	//
	// 	router.
	// 		Methods(route.Method).
	// 		Path(route.Pattern).
	// 		Name(route.Name).
	// 		Handler(handler)
	//
	// }

	return p
}
