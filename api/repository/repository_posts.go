package repository

import "golang_api_fullstack/api/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindByID(uint64) (models.Post, error)
	Update(uint64, models.Post) (int64, error)
	Delete(uint64) (int64, error)
}
