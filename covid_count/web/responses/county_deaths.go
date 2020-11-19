package responses

import "github.com/maxheckel/covid_county/covid_count/domain"

type CountiesDeaths struct {
	Total    int            `json:"total"`
	Counties []*CountyDeaths `json:"counties"`
}

type CountyDeaths struct {
	Name  string                   `json:"name"`
	Total int                      `json:"total"`
	Days  []*domain.DailyInstances `json:"days"`
}
