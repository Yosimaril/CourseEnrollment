package repositories

import (
	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/constants"
	"yosimaril/CourseEnrollment/models"
)

type CoursePlanRepository struct{}

func (r *CoursePlanRepository) GetAll(studentID uint, status string) ([]models.CoursePlan, error) {
	var coursePlans []models.CoursePlan

	db := config.DB
	db = db.Preload("Student").Preload("Items.Course").Order("created_at desc")

	if studentID != 0 {
		db = db.Where("student_id = ?", studentID)
	}
	if status != "" {
		db = db.Where("status = ?", status)
	}

	result := db.Find(&coursePlans)

	if result.Error != nil {
		return nil, result.Error
	}

	return coursePlans, nil
}

func (r *CoursePlanRepository) GetByID(id uint) (models.CoursePlan, error) {
	var coursePlan models.CoursePlan

	result := config.DB.
		Preload("Student").
		Preload("Items.Course").
		First(&coursePlan, "id = ?", id)

	return coursePlan, result.Error
}

func (r *CoursePlanRepository) GetDraftByStudent(studentID uint) (models.CoursePlan, error) {
	var coursePlan models.CoursePlan

	result := config.DB.
		Preload("Student").
		Preload("Items.Course").
		Where("student_id = ? AND status = ?", studentID, constants.CoursePlanDraft).
		Order("created_at desc").
		First(&coursePlan)

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
