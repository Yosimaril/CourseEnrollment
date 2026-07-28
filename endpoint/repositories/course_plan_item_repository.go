package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/models"
)

type CoursePlanItemRepository struct{}

func (r *CoursePlanItemRepository) GetAll(coursePlanID, courseID uint) ([]models.CoursePlanItem, error) {
	var coursePlanItems []models.CoursePlanItem

	cacheKey := "course_plan_item"

	if coursePlanID != 0 {
		cacheKey += ":cp" + fmt.Sprintf("%d", coursePlanID)
	}
	if courseID != 0 {
		cacheKey += ":c" + fmt.Sprintf("%d", courseID)
	}

	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT")

		if err := json.Unmarshal([]byte(cached), &coursePlanItems); err == nil {
			return coursePlanItems, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS")

	db := config.DB

	if coursePlanID != 0 {
		db = db.Where("course_plan_id = ?", coursePlanID)
	}
	if courseID != 0 {
		db = db.Where("course_id = ?", courseID)
	}

	result := db.Find(&coursePlanItems)

	if result.Error != nil {
		return nil, result.Error
	}

	data, err := json.Marshal(coursePlanItems)

	if err == nil {
		config.Redis.Set(
			config.Ctx,
			cacheKey,
			data,
			time.Minute,
		)
	}

	return coursePlanItems, nil
}

func (r *CoursePlanItemRepository) GetByID(coursePlanID, courseID uint) (models.CoursePlanItem, error) {
	var coursePlanItem models.CoursePlanItem

	result := config.DB.
		Where("course_plan_id = ? AND course_id = ?", coursePlanID, courseID).
		First(&coursePlanItem)

	return coursePlanItem, result.Error
}

func (r *CoursePlanItemRepository) Create(coursePlanItem *models.CoursePlanItem) error {
	return config.DB.Create(coursePlanItem).Error
}

func (r *CoursePlanItemRepository) Update(coursePlanItem *models.CoursePlanItem) error {
	return config.DB.Updates(coursePlanItem).Error
}

func (r *CoursePlanItemRepository) Delete(coursePlanID, courseID uint) error {
	return config.DB.Delete(&models.CoursePlanItem{}, "course_plan_id = ? AND course_id = ?", coursePlanID, courseID).Error
}