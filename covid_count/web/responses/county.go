package responses

import "github.com/maxheckel/covid_county/covid_count/domain"

type County struct {
	Deaths map[int][]*domain.MonthlyCountyDeaths `json:"deaths"`
}
