package domain

import (
	"errors"
	"strconv"
	"time"
)

type MonthlyCountyDeaths struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Month     int        `json:"month" gorm:"column:month"`
	County    string     `json:"county" gorm:"column:county"`
	Count     int        `json:"count" gorm:"column:count"`
	Year      int        `json:"year" gorm:"column:year"`
}

func (r MonthlyCountyDeaths) TableName() string {
	return "imports.death_records"
}

func NewMonthlyCountyDeathsFromCSV(row []string, year int) (mcd MonthlyCountyDeaths, err error) {
	mcd.Month, err = monthStringToInt(row[0])
	if err != nil {
		return mcd, err
	}
	mcd.Year = year
	mcd.County = row[1]
	mcd.Count, err = strconv.Atoi(row[2])
	// If the year is 0 then it's the average of 10 years
	if year == 0 {
		mcd.Count = mcd.Count/5
	}

	if err != nil {
		return mcd, err
	}

	return mcd, nil
}

func monthStringToInt(month string) (int, error) {
	switch month {
	case "January":
		return 1, nil
	case "February":
		return 2, nil
	case "March":
		return 3, nil
	case "April":
		return 4, nil
	case "May":
		return 5, nil
	case "June":
		return 6, nil
	case "July":
		return 7, nil
	case "August":
		return 8, nil
	case "September":
		return 9, nil
	case "October":
		return 10, nil
	case "November":
		return 11, nil
	case "December":
		return 12, nil
	}
	return 0, errors.New("invalid month representation")
}
