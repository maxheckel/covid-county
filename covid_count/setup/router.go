package setup

import (
	"github.com/gorilla/mux"
	"github.com/maxheckel/covid_county/covid_count/web/handlers"
)

func Router(a *App) *mux.Router{
	router := mux.NewRouter()
	router.Handle("/api/overview", handlers.Overview{Manager: a.Data, Cache: a.Cache}).Methods("GET")
	router.PathPrefix("/").Handler(handlers.SpaHandler{
		StaticPath: "/server/web/dist",
		IndexPath:  "/server/web/dist/index.html",
	}).Methods("GET")


	return router
}


