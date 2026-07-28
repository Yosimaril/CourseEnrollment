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

func (r *UserRepository) GetByID(id uint) (models.User, error) { // Corrected return type from models.user to models.User
	var user models.User

	cacheKey := fmt.Sprintf("user:%d", id)
	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT")
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return user, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS")

	result := config.DB.
		First(&user, "id = ?", id)

	if result.Error != nil {
		return user, result.Error
	}

	data, err := json.Marshal(user)
	if err == nil {
		config.Redis.Set(
			config.Ctx,
			cacheKey,
			data,
			time.Minute,
		)
	}

	return user, result.Error
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	cacheKey := "user:email:" + email
	cached, err := config.Redis.Get(config.Ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("[Redis] CACHE HIT")
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return user, nil
		}
	}

	fmt.Println("[Redis] CACHE MISS")

	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	data, err := json.Marshal(user)
	if err == nil {
		config.Redis.Set(
			config.Ctx,
			cacheKey,
			data,
			time.Minute,
		)
	}

	return user, result.Error
}

func (r *UserRepository) Create(user *models.User) error { // Corrected parameter name from category to user
	return config.DB.Create(user).Error
}

func (r *UserRepository) Update(user *models.User) error { // Corrected parameter name from category to user
	return config.DB.Updates(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}