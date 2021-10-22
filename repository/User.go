package repository

import (
	"backend-crud-user/models"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	FindAll() ([]models.Profile, error)
	FindByUserID(ID int) (models.Profile, error)
	InsertData(profile models.Profile) (models.Profile, error)
	UpdateData(ID int, profile models.Profile, user models.User) (models.Profile, error)
	DeleteData(ID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) FindAll() ([]models.Profile, error) {
	var Profile []models.Profile

	err := r.db.Preload("User").Find(&Profile).Error

	if err != nil {
		return Profile, err
	}

	return Profile, nil
}

func (r *repository) FindByUserID(ID int) (models.Profile, error) {
	var Profile models.Profile

	err := r.db.Preload("User").Where("id = ?", ID).Find(&Profile).Error
	if err != nil {
		return Profile, err
	}

	return Profile, nil
}

func (r *repository) InsertData(profile models.Profile) (models.Profile, error) {

	err := r.db.Create(&profile).Error
	if err != nil {
		return profile, err
	}

	err = r.db.Save(&profile).Error
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *repository) UpdateData(ID int, profile models.Profile, user models.User) (models.Profile, error) {
	err := r.db.Save(&profile).Error

	if err != nil {
		return profile, err
	}

	err = r.db.Preload("Profile").Take(&user, "id = ?", ID).Error

	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *repository) DeleteData(ID int) (bool, error) {
	err := r.db.Select("User").Delete(&models.Profile{ID: ID}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
