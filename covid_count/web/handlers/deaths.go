package handlers

import (
	"github.com/maxheckel/covid_county/covid_count/domain"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/service"
	"github.com/maxheckel/covid_county/covid_count/web/responses"
	"net/http"
	"strings"
)

type CountiesDeaths struct {
	Data repository.Manager
	Cache service.Cache
}

func (c CountiesDeaths) ServeHTTP(w http.ResponseWriter, r *http.Request){
	countiesString := r.URL.Query()["counties"][0]
	counties := strings.Split(countiesString, ",")
	cacheKey := countiesString + "_deaths"
	cachedRes, found := c.Cache.Get(cacheKey)
	if found {
		responses.WriteJSON(w, 200, cachedRes)
		return
	}
	for k, v := range counties{
		counties[k] = strings.ToLower(v)
	}
	countyRes, err := c.Data.Cases().GetAllDeathsForCounties(counties)
	if err != nil {
		responses.WriteJSONError(w, r, err)
		return
	}
	res := responses.CountiesDeaths{
		Total: 0,
	}
	for _, val := range countyRes {
		res.Total += val.Count
		found := false
		for _, county := range res.Counties {
			if strings.ToLower(county.Name) == strings.ToLower(val.County) {
				found = true
				county.Total += val.Count
				county.Days = append(county.Days, val)
			}
		}
		if !found {
			res.Counties = append(res.Counties, &responses.CountyDeaths{
				Total: val.Count,
				Name: val.County,
				Days: []*domain.DailyInstances{
					val,
				},
			})
		}
	}
	c.Cache.Set(cacheKey, res, 10)
	responses.WriteJSON(w, 200, res)

}
