package main

import (
	"github.com/gorilla/pat"
)

func NewRouter() *pat.Router {

	p := pat.New()
	p.Get("/auth/{provider}/session", AuthSession)
	return p
}
