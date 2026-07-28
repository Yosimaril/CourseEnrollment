package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/models"
)

type CategoryRepository struct{}

func (r *CategoryRepository) GetAll(keyword string) ([]models.Category, error) {
	var categories []models.Category

	cacheKey := "categories"

	if keyword != "" {
		cacheKey += ":" + keyword
	}

	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT")

		if err := json.Unmarshal([]byte(cached), &categories); err == nil {
			return categories, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS")

	db := config.DB

	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	result := db.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	data, err := json.Marshal(categories)

	if err == nil {
		config.Redis.Set(
			config.Ctx,
			cacheKey,
			data,
			time.Minute,
		)
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(id uint) (models.Category, error) {
	var category models.Category

	result := config.DB.
		First(&category, "id = ?", id)

	return category, result.Error
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return config.DB.Create(category).Error
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return config.DB.Updates(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Category{}, id).Error
}
