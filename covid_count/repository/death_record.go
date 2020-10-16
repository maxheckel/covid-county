package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"sync"
	"time"
)

type DeathRecord struct {
	Database *gorm.DB
}

func (dr *DeathRecord) CreateMultiple(records []domain.MonthlyCountyDeaths) error {
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

func (dr *DeathRecord) ClearPreviousMonthlyCountyDeathss() error {
	dummy := domain.MonthlyCountyDeaths{}
	return dr.Database.Exec("TRUNCATE TABLE "+dummy.TableName()).Error
}

func (dr *DeathRecord) insertAsync(records []domain.MonthlyCountyDeaths, wg *sync.WaitGroup){
	defer wg.Done()
	for _, rec := range records {
		dr.Database.LogMode(true).Create(&rec)
	}

}
