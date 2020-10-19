package responses

import (
	"github.com/maxheckel/covid_county/covid_count/domain"
	"sort"
	"time"
)

type County struct {
	Name                  string                                `json:"name"`
	Deaths                map[int][]*domain.MonthlyCountyDeaths `json:"deaths"`
	DailyCases            []*domain.DailyInstances              `json:"daily_cases"`
	DailyDeaths           []*domain.DailyInstances              `json:"daily_deaths"`
	DailyHospitalizations []*domain.DailyInstances              `json:"daily_hospitalizations"`
}

func (c *County) Sort() {

	for _, year := range c.Deaths {
		sort.SliceStable(year, func(i, j int) bool {
			return year[i].Month < year[j].Month
		})
	}
	sort.SliceStable(c.DailyCases, func(i, j int) bool {
		return c.DailyCases[i].Date.Before(c.DailyCases[j].Date)
	})
	sort.SliceStable(c.DailyDeaths, func(i, j int) bool {
		return c.DailyDeaths[i].Date.Before(c.DailyDeaths[j].Date)
	})
	sort.SliceStable(c.DailyHospitalizations, func(i, j int) bool {
		return c.DailyHospitalizations[i].Date.Before(c.DailyHospitalizations[j].Date)
	})
}

func (c *County) Fill(numDays int) {
	c.fillCases(numDays)
	c.fillDeaths(numDays)
	c.fillHospitalizations(numDays)
	c.Sort()
}

func (c *County) fillCases(numDays int) {
	casesToCheck := c.DailyCases
	if len(casesToCheck) == numDays {
		return
	}
	casesToAdd := populateMissingDays(numDays, casesToCheck, c, domain.Case)
	c.DailyCases = append(c.DailyCases, casesToAdd...)
}

func (c *County) fillDeaths(numDays int) {
	casesToCheck := c.DailyDeaths
	if len(casesToCheck) == numDays {
		return
	}
	casesToAdd := populateMissingDays(numDays, casesToCheck, c, domain.Death)
	c.DailyDeaths = append(c.DailyDeaths, casesToAdd...)
}

func (c *County) fillHospitalizations(numDays int) {
	casesToCheck := c.DailyHospitalizations
	if len(casesToCheck) == numDays {
		return
	}
	casesToAdd := populateMissingDays(numDays, casesToCheck, c, domain.Hospitalization)
	c.DailyHospitalizations = append(c.DailyHospitalizations, casesToAdd...)
}

func populateMissingDays(numDays int, casesToCheck []*domain.DailyInstances, c *County, instanceType domain.InstanceType) []*domain.DailyInstances {
	var casesToAdd []*domain.DailyInstances
	for i := 0; i < numDays; i++ {
		tfmt := "2006-01-02"
		if i > (len(casesToCheck) - 1) {
			casesToAdd = appendCase(casesToAdd, instanceType, c, i)
			continue
		}

		found := false
		for _, c := range casesToCheck {
			if c.Date.Format(tfmt) == time.Now().AddDate(0, 0, i*-1).Format(tfmt) {
				found = true
				break
			}
		}
		if !found {
			casesToAdd = appendCase(casesToAdd, instanceType, c, i)
		}
	}
	return casesToAdd
}

func appendCase(casesToAdd []*domain.DailyInstances, instanceType domain.InstanceType, c *County, i int) []*domain.DailyInstances {
	dateOnly, _ := time.Parse("2006-01-02", time.Now().AddDate(0, 0, i*-1).Format("2006-01-02"))
	casesToAdd = append(casesToAdd, &domain.DailyInstances{
		Type:   instanceType,
		County: c.Name,
		Date:   dateOnly,
		Count:  0,
	})
	return casesToAdd
}
