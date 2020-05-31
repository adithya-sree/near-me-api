package app

import (
	"fmt"
	"log"
	"nearme-api/app/db"
	"nearme-api/app/handler"
	"net/http"

	"github.com/gorilla/mux"
)

//App app struct
type App struct {
	Router  *mux.Router
	Handler *handler.Handler
}

//Initialize creates db & initializes connection
func (a *App) Initialize(connection string) error {
	db, err := db.NewClient(connection)
	if err != nil {
		return err
	}

	a.Handler = handler.NewHandler(db)
	a.Router = mux.NewRouter().StrictSlash(true)
	a.setRoutes()

	return nil
}

func (a *App) setRoutes() {
	a.Router.HandleFunc("/", a.Handler.Base)
	a.Router.HandleFunc("/api", a.Handler.Base)
	a.Router.HandleFunc("/api/uptime", a.Handler.Uptime)
	a.Router.HandleFunc("/api/running", a.Handler.Running)
	a.Router.HandleFunc("/api/location", a.Handler.AddLocation).Methods("POST")
	a.Router.HandleFunc("/api/location", a.Handler.GetLocation).Methods("GET")
}

//Run runs the mux server
func (a *App) Run(addr string) {
	fmt.Println("Starting NearMe app on port", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
