package main

import (
	"github.com/maxheckel/covid_county/covid_count/config"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/setup"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(cfg)
	}
	db, err := setup.Database(cfg)
	if err != nil {
		log.Fatal(err)
	}
	manager := repository.NewManager(db)
	cacheDriver := cache.New(5*time.Minute, 10*time.Minute)
	cacheService := setup.NewCache(cacheDriver)
	app := setup.NewApp(cfg, db, manager, cacheService)
	router := setup.Router(app)
	server, err := setup.New(router)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
