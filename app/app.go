package app

import (
	"log"
	"net/http"

	"github.com/Binod210/gomongoCRUD/api"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Api    *api.API
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.AddHandler()

}

func (a *App) AddHandler() {
	a.Router.HandleFunc("/user", a.Api.CreateUser).Methods("POST")
	a.Router.HandleFunc("/user", a.Api.GetAllUsers).Methods("GET")
	a.Router.HandleFunc("/user", a.Api.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id}", a.Api.DeleteUser).Methods("DELETE")
}

func (a *App) Run(addr string) {
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Fatalf("Could not start Server %v", err)
	}

}
