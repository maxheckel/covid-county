package repository

import "github.com/jinzhu/gorm"

type Manager struct {
	cases *Cases
	record *Record
	isUpdating *IsUpdating
}

func (m Manager) IsUpdating() *IsUpdating {
	return m.isUpdating
}

func (m Manager) Record() *Record {
	return m.record
}

func (m Manager) Cases() *Cases {
	return m.cases
}

func NewManager(db *gorm.DB) *Manager{
	return &Manager{
		cases:      &Cases{Database: db},
		record:     &Record{Database: db},
		isUpdating: &IsUpdating{Database: db},
	}
}
