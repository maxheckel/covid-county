package handlers

import (
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/service"
	"net/http"
)

type County struct{
	Data *repository.Manager
	Cache *service.Cache
}

func (c County) ServeHTTP(w http.ResponseWriter, r *http.Request){

}
