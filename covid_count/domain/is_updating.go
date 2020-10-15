package domain

type IsUpdating struct {
	IsUpdating bool `json:"is_updating" gorm:"column:is_updating"`
}

func (IsUpdating) TableName() string {
	return "imports.is_updating"
}
