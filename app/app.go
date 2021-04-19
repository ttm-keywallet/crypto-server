package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	h "github.com/ttmbank/crypto-server/handlers"
	"log"
	"net/http"
)

type App struct {
	Router     *mux.Router
}

func (a *App) Run(host string) {

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Origin", "Content-Type"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowCredentials(),
	)(a.Router)

	log.Fatal(http.ListenAndServe(host, cors))
}

// Initialize
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Get
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}


// setRouters
func (a *App) setRouters() {
	a.Post("/api/script", a.GetScript)
}


func (a *App) GetScript(w http.ResponseWriter, r *http.Request) {
	h.GetScript(w, r)
}
