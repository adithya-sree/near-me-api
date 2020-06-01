package app

import (
	"fmt"
	"log"
	"nearme-api/src/app/db"
	"nearme-api/src/app/handler"
	"nearme-api/src/config"
	"net/http"

	"github.com/gorilla/mux"
)

//App app struct
type App struct {
	Router  *mux.Router
	Handler *handler.Handler
	config  config.AppConfig
}

//Initialize creates db & initializes connection
func (a *App) Initialize(c config.AppConfig) error {
	db, err := db.NewClient(c.ConnectionString())
	if err != nil {
		return err
	}

	a.Handler = handler.NewHandler(db)
	a.Router = mux.NewRouter().StrictSlash(true)
	a.config = c
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
func (a *App) Run() {
	fmt.Println("Starting NearMe app on port", a.config.AppPort)
	log.Fatal(http.ListenAndServe(a.config.AppPort, a.Router))
}
