package service

import (
	"github.com/maxheckel/covid_county/covid_count/domain"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"io"
	"strconv"
)

type DeathLoader struct {
	Data *repository.Manager
}

func NewDeathLoader(data *repository.Manager) *DeathLoader {
	return &DeathLoader{Data: data}
}

func (dl *DeathLoader) Load() error{
	err := dl.Data.DeathRecords().ClearPreviousMonthlyCountyDeathss()
	if err != nil {
		return err
	}
	var records []domain.MonthlyCountyDeaths
	for currentYear := 2018; currentYear <= 2020; currentYear++ {
		csvReader := GetReaderForCSV("./data/imports/"+strconv.Itoa(currentYear)+"deaths.csv")
		for {
			line, err := csvReader.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			record, err := domain.NewMonthlyCountyDeathsFromCSV(line, currentYear)
			records = append(records, record...)
		}
	}

	err = dl.Data.DeathRecords().CreateMultiple(records)
	if err != nil{
		return err
	}

	return nil
}



