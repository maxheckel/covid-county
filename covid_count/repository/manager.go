package repository

import "github.com/jinzhu/gorm"
type Manager interface {
	IsUpdating() IsUpdating
	Record() Record
	Cases() Cases
	DeathRecords() DeathRecord
}

type manager struct {
	cases Cases
	record Record
	isUpdating IsUpdating
	deathRecords DeathRecord
}

func (m manager) IsUpdating() IsUpdating {
	return m.isUpdating
}

func (m manager) Record() Record {
	return m.record
}

func (m manager) Cases() Cases {
	return m.cases
}

func (m manager) DeathRecords() DeathRecord {
	return m.deathRecords
}

func NewManager(db *gorm.DB) Manager{
	return &manager{
		cases:      NewCasesRepository(db),
		record:     NewRecordRepository(db),
		isUpdating: NewIsUpdatingRepository(db),
		deathRecords: NewDeathRecordRepository(db),
	}
}
