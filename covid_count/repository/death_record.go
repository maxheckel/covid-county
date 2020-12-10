package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"sync"
	"time"
)
type DeathRecord interface {
	CreateMultiple(records []domain.MonthlyCountyDeaths) error
	ClearPreviousMonthlyCountyDeaths() error
	insertAsync(records []domain.MonthlyCountyDeaths, wg *sync.WaitGroup)
	GetForCounty(county string) ([]*domain.MonthlyCountyDeaths, error)
}

func NewDeathRecordRepository(db *gorm.DB) DeathRecord{
	return &deathRecord{
		Database: db,
	}
}

type deathRecord struct {
	Database *gorm.DB
}

func (dr *deathRecord) CreateMultiple(records []domain.MonthlyCountyDeaths) error {
	start := time.Now()
	batches := 10
	batchSize := len(records) / batches
	wg := sync.WaitGroup{}
	wg.Add(batches)
	for i := 0; i < batches; i++ {
		if i == batches-1 {
			go dr.insertAsync(records[batchSize*i:], &wg)
		} else {
			go dr.insertAsync(records[batchSize*i : batchSize*(i+1)], &wg)
		}

	}
	wg.Wait()
	end := time.Now()
	fmt.Println(start)
	fmt.Println(end)
	return nil
}

func (dr *deathRecord) ClearPreviousMonthlyCountyDeaths() error {
	dummy := domain.MonthlyCountyDeaths{}
	return dr.Database.Exec("TRUNCATE TABLE "+dummy.TableName()).Error
}

func (dr *deathRecord) insertAsync(records []domain.MonthlyCountyDeaths, wg *sync.WaitGroup){
	defer wg.Done()
	for _, rec := range records {
		dr.Database.LogMode(true).Create(&rec)
	}
}

func (dr *deathRecord) GetForCounty(county string) ([]*domain.MonthlyCountyDeaths, error) {
	var res []*domain.MonthlyCountyDeaths
	err := dr.Database.Where("lower(county) = lower(?)", county).Find(&res).Error
	return res, err
}
