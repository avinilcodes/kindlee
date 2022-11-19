package server

import (
	"net/http"
	"taskmanager/login"

	runtimemid "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

const (
// versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	router = mux.NewRouter()
	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost)

	//Add user

	ops := runtimemid.RedocOpts{SpecURL: "swagger.yaml"}
	sh := runtimemid.Redoc(ops, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	return
}
