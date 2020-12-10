package handlers

import (
	"github.com/gorilla/mux"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/service"
	"github.com/maxheckel/covid_county/covid_count/web/responses"
	"net/http"
	"strings"
	"time"
)

type County struct{
	Data repository.Manager
	Cache service.Cache
}

func (c County) ServeHTTP(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	res, err := c.Data.DeathRecords().GetForCounty(vars["county"])
	if err != nil {
		responses.WriteJSONError(w, r, err)
	}
	output := responses.County{
		Deaths: map[int][]*domain.MonthlyCountyDeaths{},
	}
	for _, record := range res{
		output.Deaths[record.Year] = append(output.Deaths[record.Year], record)
	}

	t1, _ := time.Parse("2006-01-02", "2020-02-01")


	t2 := time.Now()
	numDays := int(t2.Sub(t1).Hours() / 24)
	output.DailyCases, err = c.Data.Cases().GetCountyCasesForDays(numDays, vars["county"])
	if err != nil {
		responses.WriteJSONError(w, r, err)
	}
	output.DailyDeaths, err = c.Data.Cases().GetCountyDeathsForDays(numDays, vars["county"])
	if err != nil {
		responses.WriteJSONError(w, r, err)
	}
	output.DailyHospitalizations, err = c.Data.Cases().GetCountyHospitalizationsForDays(numDays, vars["county"])
	if err != nil {
		responses.WriteJSONError(w, r, err)
	}
	output.Name =  strings.Title(vars["county"])
	output.Sort()
	output.Fill(numDays)

	responses.WriteJSON(w, 200, output)
}

