package domain

import (
	"strconv"
	"time"
)

type Record struct {
	ID                uint `gorm:"primary_key"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
	County            string     `json:"county" gorm:"column:county"`
	Sex               string     `json:"sex" gorm:"column:sex"`
	Age               string     `json:"age" gorm:"column:age"`
	OnsetDate         time.Time  `json:"onset_date" gorm:"column:onset_date"`
	DeathDate         *time.Time `json:"death_date" gorm:"column:death_date"`
	AdmissionDate     *time.Time `json:"admission_date" gorm:"column:admission_date"`
	CaseCount         int        `json:"case_count" gorm:"column:case_count"`
	DeathCount        int        `json:"death_count" gorm:"column:death_count"`
	HospitalizedCount int        `json:"hospitalized_count" gorm:"column:hospitalized_count"`
}

func (r Record) TableName() string {
	return "imports.records"
}

func NewFromCSV(row []string) (Record, error) {
	layout := "1/2/2006"
	onsetDate, err := time.Parse(layout, row[3])
	var deathDate time.Time
	if row[4] != "" {
		deathDate, err = time.Parse(layout, row[4])
	}
	var admissionDate time.Time
	if row[5] != "" {
		admissionDate, err = time.Parse(layout, row[5])
	}

	caseCount, err := strconv.Atoi(row[6])
	deathCount, err := strconv.Atoi(row[7])
	hospitalizationCount, err := strconv.Atoi(row[8])
	if err != nil {
		return Record{}, err
	}
	return Record{
		County:            row[0],
		Sex:               row[1],
		Age:               row[2],
		OnsetDate:         onsetDate,
		DeathDate:         &deathDate,
		AdmissionDate:     &admissionDate,
		CaseCount:         caseCount,
		DeathCount:        deathCount,
		HospitalizedCount: hospitalizationCount,
	}, nil
}
