package service

import (
	"backend-crud-user/models"
	"backend-crud-user/repository"
	"backend-crud-user/validation"
)

type ServiceUser interface {
	FetchAll() ([]models.Profile, error)
	Insert(inputUser validation.InsertData) (models.Profile, error)
	GetByID(ID int) (models.Profile, error)
	UpdateUser(ID int, UpdateUser validation.InsertData) (models.Profile, error)
	DeleteByID(ID int) (bool, error)
}

type service struct {
	repository repository.RepositoryUser
}

func NewServiceUser(repository repository.RepositoryUser) *service {
	return &service{repository}
}

func (s *service) DeleteByID(ID int) (bool, error) {
	_, err := s.repository.DeleteData(ID)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) GetByID(ID int) (models.Profile, error) {
	dataUser, err := s.repository.FindByUserID(ID)
	if err != nil {
		return models.Profile{}, err
	}

	return dataUser, nil
}

func (s *service) FetchAll() ([]models.Profile, error) {
	dataUser, err := s.repository.FindAll()

	if err != nil {
		return []models.Profile{}, err
	}

	return dataUser, nil
}

func (s *service) Insert(inputUser validation.InsertData) (models.Profile, error) {
	var profile models.Profile

	profile.AlamatKtp = inputUser.AlamatKtp
	profile.Pekerjaan = inputUser.Pekerjaan
	profile.NamaLengkap = inputUser.NamaLengkap
	profile.PendidikanTerakhir = inputUser.PendidikanTerakhir
	profile.NoHp = inputUser.NoHp
	profile.User.Email = inputUser.Email
	profile.User.Username = inputUser.Username
	profile.User.Email = inputUser.Email

	save, err := s.repository.InsertData(profile)

	if err != nil {
		return profile, err
	}

	return save, nil
}

func (s *service) UpdateUser(ID int, UpdateUser validation.InsertData) (models.Profile, error) {
	var profile models.Profile
	var user models.User

	user.Email = UpdateUser.Email
	user.Username = UpdateUser.Username
	user.Email = UpdateUser.Email

	profile.AlamatKtp = UpdateUser.AlamatKtp
	profile.Pekerjaan = UpdateUser.Pekerjaan
	profile.NamaLengkap = UpdateUser.NamaLengkap
	profile.PendidikanTerakhir = UpdateUser.PendidikanTerakhir
	profile.NoHp = UpdateUser.NoHp

	updateData, err := s.repository.UpdateData(ID, profile, user)

	if err != nil {
		return models.Profile{}, err
	}

	return updateData, nil

}
