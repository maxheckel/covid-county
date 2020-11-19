package setup

import (
	"github.com/gorilla/mux"
	"github.com/maxheckel/covid_county/covid_count/web/handlers"
)

func Router(a *App) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/api/overview", handlers.Overview{
		Data: a.Data,
		Cache: a.Cache,
	}).Methods("GET")

	router.Handle("/api/counties/deaths", handlers.CountiesDeaths{
		Data:  a.Data,
		Cache: a.Cache,
	}).Methods("GET")

	router.Handle("/api/county/{county}", handlers.County{
		Data:  a.Data,
		Cache: a.Cache,
	}).Methods("GET")

	router.PathPrefix("/").Handler(handlers.SpaHandler{
		StaticPath: a.Config.SPARoot,
		IndexPath:  "index.html",
	}).Methods("GET")

	return router
}
