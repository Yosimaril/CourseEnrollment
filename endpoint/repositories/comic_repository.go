package repositories

import (
	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/models"
)

type ComicRepository struct{}

func (r *ComicRepository) GetAll(keyword string) ([]models.Comic, error) {
	var comics []models.Comic

	db := config.DB.
		Preload("Creator").
		Preload("Categories")

	if keyword != "" {
		db = db.Where("title LIKE ?", "%"+keyword+"%")
	}

	result := db.Find(&comics)

	return comics, result.Error
}

func (r *ComicRepository) GetByID(id string) (models.Comic, error) {
	var comic models.Comic

	result := config.DB.
		Preload("Creator").
		Preload("Categories").
		First(&comic, id)

	return comic, result.Error
}
