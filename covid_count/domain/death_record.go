package domain

import (
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

func NewMonthlyCountyDeathsFromCSV(row []string, year int) (mcd []MonthlyCountyDeaths, err error){
	county := row[0]
	for i := 1; i <= 12; i++ {
		deaths, err := strconv.Atoi(row[i])
		if err != nil {
			return mcd, err
		}
		mcd = append(mcd, MonthlyCountyDeaths{
			Month:  i,
			County: county,
			Count:  deaths,
			Year:   year,
		})
	}
	return mcd, nil
}
