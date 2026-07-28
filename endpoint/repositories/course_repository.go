package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/models"
)

type CourseRepository struct{}

func (r *CourseRepository) GetAll(name string) ([]models.Course, error) {
	var courses []models.Course

	cacheKey := "course"

	if name != "" {
		cacheKey += ":" + name
	}

	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT")

		if err := json.Unmarshal([]byte(cached), &courses); err == nil {
			return courses, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS")

	db := config.DB

	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	result := db.Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}

	data, err := json.Marshal(courses)

	if err == nil {
		config.Redis.Set(
			config.Ctx,
			cacheKey,
			data,
			time.Minute,
		)
	}

	return users, nil
}

func (r *CourseRepository) GetByID(id uint) (models.Course, error) {
	var course models.Course

	result := config.DB.
		First(&course, "id = ?", id)

	return course, result.Error
}

func (r *CourseRepository) Create(course *models.Course) error {
	return config.DB.Create(course).Error
}

func (r *CourseRepository) Update(course *models.Course) error {
	return config.DB.Updates(course).Error
}

func (r *CourseRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Course{}, id).Error
}
