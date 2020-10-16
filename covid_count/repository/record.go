package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"sync"
	"time"
)

type Record struct {
	Database *gorm.DB
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
	start := time.Now()
	batches := 10
	batchSize := len(records) / batches
	wg := sync.WaitGroup{}
	wg.Add(batches)
	for i := 0; i < batches; i++ {
		if i == batches-1 {
			go r.insertAsync(records[batchSize*i:], &wg)
		} else {
			go r.insertAsync(records[batchSize*i : batchSize*(i+1)], &wg)
		}

	}
	wg.Wait()
	end := time.Now()
	fmt.Println(start)
	fmt.Println(end)
	return nil
}

func (r *Record) ClearPreviousRecords() error {
	dummy := domain.Record{}
	return r.Database.Exec("TRUNCATE TABLE "+dummy.TableName()).Error
}

func (r *Record) insertAsync(records []domain.Record, wg *sync.WaitGroup){
	defer wg.Done()
	for _, rec := range records {
		r.Database.LogMode(true).Create(&rec)
	}

}
