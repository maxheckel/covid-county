package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/maxheckel/covid_county/covid_count/domain"
)

type IsUpdating struct {
	Database *gorm.DB
}


func (i *IsUpdating) IsUpdating() (bool, error) {
	var res domain.IsUpdating
	err := i.Database.First(&res).Error
	if err != nil {
		return false, err
	}
	return res.IsUpdating, nil
}

func (i *IsUpdating) SetIsUpdating(isUpdating bool) error {
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
