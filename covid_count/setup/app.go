package setup

import (
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/config"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/service"
)

type App struct {
	Config *config.Config
	Database *gorm.DB
	Data *repository.Manager
	Cache *service.Cache
}

func NewApp(cfg *config.Config, db *gorm.DB, mg *repository.Manager, ca *service.Cache) *App {
	return &App{
		Config:   cfg,
		Database: db,
		Data:     mg,
		Cache: ca,
	}
}
