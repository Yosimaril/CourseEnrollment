package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/models"
)

type UserRepository struct{}

func (r *UserRepository) GetAll(username string) ([]models.User, error) {
	var users []models.User

	cacheKey := "user"

	if username != "" {
		cacheKey += ":" + username
	}

	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT")

		if err := json.Unmarshal([]byte(cached), &users); err == nil {
			return users, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS")

	db := config.DB

	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}

	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	data, err := json.Marshal(users)

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

func (r *UserRepository) GetByID(id uint) (models.user, error) {
	var user models.User

	result := config.DB.
		First(&user, "id = ?", id)

	return user, result.Error
}

func (r *UserRepository) Create(category *models.User) error {
	return config.DB.Create(user).Error
}

func (r *UserRepository) Update(category *models.User) error {
	return config.DB.Updates(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
