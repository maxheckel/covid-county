package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"time"
)

type Record struct {
	Database *gorm.DB
}

func NewRecord(db *gorm.DB) *Record {
	return &Record{
		Database: db,
	}
}

func (r *Record) MaxDate(col string) (*time.Time, error) {
	var res string
	r.Database.LogMode(true).Select("MAX("+col+")").Table("imports.records").Limit(1).Row().Scan(&res)
	if res == "" {
		return nil, nil
	}
	timeRes, err := time.Parse(time.RFC3339, res)
	return &timeRes, err
}

func (r *Record) CreateMultiple(records []domain.Record) error {
	now := time.Now()
	var records2 = []domain.Record{
		{
			County:            "awgawe",
			Sex:               "wagaweg",
			Age:               "awgawe",
			OnsetDate:         time.Now(),
			DeathDate:         &now,
			AdmissionDate:     &now,
			CaseCount:         2,
			DeathCount:        3,
			HospitalizedCount: 5,
		},
	}
	return r.Database.LogMode(true).Create(&records2).Error
}
