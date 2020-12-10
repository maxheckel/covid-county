package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
)
type IsUpdating interface {
	IsUpdating() (bool, error)
	SetIsUpdating(isUpdating bool) error
}

type isUpdating struct {
	Database *gorm.DB
}

func NewIsUpdatingRepository(db *gorm.DB) IsUpdating{
	return &isUpdating{Database: db}
}

func (i *isUpdating) IsUpdating() (bool, error) {
	var res domain.IsUpdating
	err := i.Database.First(&res).Error
	if err != nil {
		return false, err
	}
	return res.IsUpdating, nil
}

func (i *isUpdating) SetIsUpdating(isUpdating bool) error {
	var res domain.IsUpdating
	err := i.Database.First(&res).Error
	if err != nil && !gorm.IsRecordNotFoundError(err){
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		res.IsUpdating = isUpdating
		return i.Database.Create(&res).Error
	}

	return i.Database.Model(&domain.IsUpdating{}).Where("is_updating = ?", res.IsUpdating).Update("is_updating", isUpdating).Error
}
