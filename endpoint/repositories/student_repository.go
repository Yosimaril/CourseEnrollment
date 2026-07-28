package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/models"
)

type StudentRepository struct{}

func (r *StudentRepository) GetAll(firstName, lastName, studentID string) ([]models.Student, error) {
	var students []models.Student

	cacheKey := "student"
	if firstName != "" {
		cacheKey += ":firstName=" + firstName
	}
	if lastName != "" {
		cacheKey += ":lastName=" + lastName
	}
	if studentID != "" {
		cacheKey += ":studentID=" + studentID
	}

	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT for students")
		if err := json.Unmarshal([]byte(cached), &students); err == nil {
			return students, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS for students")

	db := config.DB

	if firstName != "" {
		db = db.Where("first_name LIKE ?", "%"+firstName+"%")
	}
	if lastName != "" {
		db = db.Where("last_name LIKE ?", "%"+lastName+"%")
	}
	if studentID != "" {
		db = db.Where("student_id LIKE ?", "%"+studentID+"%")
	}

	result := db.Find(&students)

	if result.Error != nil {
		return nil, result.Error
	}

	data, err := json.Marshal(students)
	if err == nil {
		config.Redis.Set(
			config.Ctx,
			cacheKey,
			data,
			time.Minute,
		)
	}

	return students, nil
}

func (r *StudentRepository) GetByID(id uint) (models.Student, error) {
	var student models.Student

	result := config.DB.
		First(&student, "id = ?", id)

	return student, result.Error
}

func (r *StudentRepository) Create(student *models.Student) error {
	return config.DB.Create(student).Error
}

func (r *StudentRepository) Update(student *models.Student) error {
	return config.DB.Updates(student).Error
}

func (r *StudentRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Student{}, id).Error
}
