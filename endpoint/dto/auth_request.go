package dto

import "yosimaril/CourseEnrollment/models"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Username   string          `json:"username" binding:"required"`
	Email      string          `json:"email" binding:"required,email"`
	Password   string          `json:"password" binding:"required,min=6"`
	Role       models.UserRole `json:"role" binding:"required,oneof=ADMIN STUDENT"`
	Nrp        *string         `json:"nrp"`
	MaxCredits *int            `json:"max_credits"`
}