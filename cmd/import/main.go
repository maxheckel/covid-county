package main

import (
	"github.com/maxheckel/covid_county/covid_count/config"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/service"
	"github.com/maxheckel/covid_county/covid_count/setup"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	database, err := setup.Database(cfg)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	manager := repository.NewManager(database)

	loader := service.NewLoader(manager)
	err = loader.Load()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
