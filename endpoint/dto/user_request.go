package dto

import "yosimaril/CourseEnrollment/constants"

type CreateUserRequest struct {
	Username   string             `json:"username" binding:"required"`
	Email      string             `json:"email" binding:"required,email"`
	Password   string             `json:"password" binding:"required,min=6"`
	Role       constants.UserRole `json:"role" binding:"required,oneof=ADMIN STUDENT"`
	Nrp        *string            `json:"nrp"`
	MaxCredits *int               `json:"max_credits"`
}

type UpdateUserRequest struct {
	Username   string             `json:"username"`
	Email      string             `json:"email" binding:"email"`
	Role       constants.UserRole `json:"role" binding:"oneof=ADMIN STUDENT"`
	Nrp        *string            `json:"nrp"`
	MaxCredits *int               `json:"max_credits"`
}
