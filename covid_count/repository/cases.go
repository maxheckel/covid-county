package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"strconv"
)

type Cases struct {
	Database *gorm.DB
}



func (c *Cases) GetAllCasesForDays(numDays int) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(case_count) as count,
  records.onset_date as date, records.county,
  'Case' as type
FROM imports.records
WHERE records.onset_date > now() - INTERVAL '`+strconv.Itoa(numDays+7)+` DAYS' and records.onset_date < now() - INTERVAL '7 days'
GROUP BY records.county, records.onset_date`).Scan(&res).Error

	return res, err
}


func (c *Cases) GetAllDeathsForDays(numDays int) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(death_count) as count,
  records.death_date as date, records.county,
  'Death' as type
FROM imports.records
WHERE records.death_date > now() - INTERVAL '`+strconv.Itoa(numDays+7)+` DAYS' and records.death_date < now() - INTERVAL '7 days'
GROUP BY records.county, records.death_date`).Scan(&res).Error

	return res, err
}


func (c *Cases) GetAllHospitalizationsForDays(numDays int) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(hospitalized_count) as count,
  records.admission_date as date, records.county,
  'Hospitalization' as type
FROM imports.records
WHERE records.admission_date > now() - INTERVAL '`+strconv.Itoa(numDays+7)+` DAYS' and records.admission_date < now() - INTERVAL '7 days'
GROUP BY records.county, records.admission_date`).Scan(&res).Error

	return res, err
}
