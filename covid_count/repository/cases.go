package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"strconv"
)
type Cases interface {
	GetAllCasesForDays(numDays int) ([]*domain.DailyInstances, error)
	GetAllDeathsForDays(numDays int) ([]*domain.DailyInstances, error)
	GetAllHospitalizationsForDays(numDays int) ([]*domain.DailyInstances, error)
	GetCountyCasesForDays(numDays int, county string) ([]*domain.DailyInstances, error)
	GetCountyDeathsForDays(numDays int, county string) ([]*domain.DailyInstances, error)
	GetAllDeathsForCounties(counties []string) ([]*domain.DailyInstances, error)
	GetCountyHospitalizationsForDays(numDays int, county string) ([]*domain.DailyInstances, error)
}

type cases struct {
	Database *gorm.DB
}

func NewCasesRepository(db *gorm.DB) Cases{
	return &cases{
		Database: db,
	}
}

func (c *cases) GetAllCasesForDays(numDays int) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(case_count) as count,
  records.onset_date as date, records.county,
  'Case' as type
FROM imports.records
WHERE records.onset_date > now() - INTERVAL '` + strconv.Itoa(numDays+7) + ` DAYS' and records.onset_date < now() - INTERVAL '7 days'
GROUP BY records.county, records.onset_date`).Scan(&res).Error

	return res, err
}

func (c *cases) GetAllDeathsForDays(numDays int) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(death_count) as count,
  records.death_date as date, records.county,
  'Death' as type
FROM imports.records
WHERE records.death_date > now() - INTERVAL '` + strconv.Itoa(numDays+7) + ` DAYS' and records.death_date < now() - INTERVAL '7 days'
GROUP BY records.county, records.death_date`).Scan(&res).Error

	return res, err
}

func (c *cases) GetAllHospitalizationsForDays(numDays int) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(hospitalized_count) as count,
  records.admission_date as date, records.county,
  'Hospitalization' as type
FROM imports.records
WHERE records.admission_date > now() - INTERVAL '` + strconv.Itoa(numDays+7) + ` DAYS' and records.admission_date < now() - INTERVAL '7 days'
GROUP BY records.county, records.admission_date`).Scan(&res).Error

	return res, err
}

func (c *cases) GetCountyCasesForDays(numDays int, county string) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(case_count) as count,
  records.onset_date as date, records.county,
  'Case' as type
FROM imports.records
WHERE records.onset_date > now() - INTERVAL '`+strconv.Itoa(numDays)+` DAYS'
AND lower(county) = lower(?)
GROUP BY records.county, records.onset_date`, county).Scan(&res).Error

	return res, err
}

func (c *cases) GetCountyDeathsForDays(numDays int, county string) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(death_count) as count,
  records.death_date as date, records.county,
  'Death' as type
FROM imports.records
WHERE records.death_date > now() - INTERVAL '`+strconv.Itoa(numDays)+` DAYS'
AND lower(county) = lower(?)
GROUP BY records.county, records.death_date`, county).Scan(&res).Error

	return res, err
}

func (c *cases) GetAllDeathsForCounties(counties []string) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(death_count) AS count,
  'Death'          AS type,
  death_date       AS date,
  county
FROM imports.records
WHERE lower(county) IN (?) AND death_count > 0 AND death_date > '2020-01-01'
GROUP BY date, county
ORDER BY death_date desc`, counties).Scan(&res).Error

	return res, err
}

func (c *cases) GetCountyHospitalizationsForDays(numDays int, county string) ([]*domain.DailyInstances, error) {
	var res []*domain.DailyInstances
	err := c.Database.Raw(`SELECT
  sum(hospitalized_count) as count,
  records.admission_date as date, records.county,
  'Hospitalization' as type
FROM imports.records
WHERE records.admission_date > now() - INTERVAL '`+strconv.Itoa(numDays)+` DAYS'
AND lower(county) = lower(?)
GROUP BY records.county, records.admission_date`, county).Scan(&res).Error

	return res, err
}
