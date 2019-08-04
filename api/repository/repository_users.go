package repository

import "golang_api_fullstack/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindByID(uint32) (models.User, error)
	Update(uint32, models.User) (int64, error)
	Delete(uint32) (int64, error)
}
