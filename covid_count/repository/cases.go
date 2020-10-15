package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"strconv"
)

type Cases struct {
	Database *gorm.DB
}

func (c *Cases) GetCountyCasesForDays(numDays int, county string) ([]*domain.DailyCases, error) {
	var res []*domain.DailyCases
	daysStr := strconv.Itoa(numDays) + " days"
	err := c.Database.Raw(`SELECT
  sum(case_count),
  records.onset_date, records.county
FROM imports.records
WHERE records.onset_date > now() - INTERVAL ? and county = ?
GROUP BY records.county, records.onset_date`, daysStr, county).Scan(&res).Error
	return res, err
}


func (c *Cases) GetAllCasesForDays(numDays int) ([]*domain.DailyCases, error) {
	var res []*domain.DailyCases
	err := c.Database.Raw(`SELECT
  sum(case_count) as count,
  records.onset_date as date, records.county
FROM imports.records
WHERE records.onset_date > now() - INTERVAL '`+strconv.Itoa(numDays+7)+` DAYS' and records.onset_date < now() - INTERVAL '7 days'
GROUP BY records.county, records.onset_date`).Scan(&res).Error
	return res, err
}
