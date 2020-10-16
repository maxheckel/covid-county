package handlers

import (
	"github.com/maxheckel/covid_county/covid_count/domain"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/service"
	"github.com/maxheckel/covid_county/covid_count/web"
	"github.com/maxheckel/covid_county/covid_count/web/responses"
	"net/http"
	"sort"
)

type County struct{
	Data *repository.Manager
	Cache *service.Cache
}

func (c County) ServeHTTP(w http.ResponseWriter, r *http.Request){
	res, err := c.Data.DeathRecords().GetForCounty("Franklin")
	if err != nil {
		web.WriteJSONError(w, r, err)
	}
	output := responses.County{
		Deaths: map[int][]*domain.MonthlyCountyDeaths{},
	}
	for _, record := range res{
		output.Deaths[record.Year] = append(output.Deaths[record.Year], record)
	}
	for _, year := range output.Deaths{
		sort.SliceStable(year, func(i, j int) bool {
			return year[i].Month < year[j].Month
		})
	}
	web.WriteJSON(w, 200, output)
}
