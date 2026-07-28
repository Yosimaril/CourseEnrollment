package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/models"
)

type CoursePlanRepository struct{}

func (r *CoursePlanRepository) GetAll(studentID uint) ([]models.CoursePlan, error) {
	var coursePlans []models.CoursePlan

	cacheKey := "course_plan"

	if studentID != 0 {
		cacheKey += ":" + fmt.Sprintf("%d", studentID)
	}

	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT")

		if err := json.Unmarshal([]byte(cached), &coursePlans); err == nil {
			return coursePlans, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS")

	db := config.DB

	if studentID != 0 {
		db = db.Where("student_id = ?", studentID)
	}

	result := db.Find(&coursePlans)

	if result.Error != nil {
		return nil, result.Error
	}

	data, err := json.Marshal(coursePlans)

	if err == nil {
		config.Redis.Set(
			config.Ctx,
			cacheKey,
			data,
			time.Minute,
		)
	}

	return coursePlans, nil
}

func (r *CoursePlanRepository) GetByID(id uint) (models.CoursePlan, error) {
	var coursePlan models.CoursePlan

	result := config.DB.
		First(&coursePlan, "id = ?", id)

	return coursePlan, result.Error
}

func (r *CoursePlanRepository) Create(coursePlan *models.CoursePlan) error {
	return config.DB.Create(coursePlan).Error
}

func (r *CoursePlanRepository) Update(coursePlan *models.CoursePlan) error {
	return config.DB.Updates(coursePlan).Error
}

func (r *CoursePlanRepository) Delete(id uint) error {
	return config.DB.Delete(&models.CoursePlan{}, id).Error
}