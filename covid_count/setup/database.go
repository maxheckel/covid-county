package setup

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/maxheckel/covid_county/covid_count/config"
	"time"
)

func Database(cfg *config.Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s connect_timeout=%d",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPassword,
		cfg.DBSSLMode,
		cfg.DBTimeout,
	)
	var err error
	var db *gorm.DB

	// ping until ready (max 10 attempts)
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open("postgres", connStr)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
