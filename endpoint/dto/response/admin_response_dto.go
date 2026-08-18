package dto

import "yosimaril/CourseEnrollment/constants"

type CreateAdminResponseDto struct {
	BaseIdDto

	Ulid string `json:"ulid" binding:"required"`
	Username   string             `json:"username" binding:"required"`
	Email      string             `json:"email" binding:"required,email"`

	BaseTimestampDto
}

type UpdateAdminResponseDto struct {
	BaseIdDto

	Ulid string `json:"ulid" binding:"required"`
	Username   *string            `json:"username"`
	Email      *string            `json:"email" binding:"email"`

	BaseTimestampDto
}
