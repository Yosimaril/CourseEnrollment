package models

import (
	"yosimaril/CourseEnrollment/constants"
)

type User struct {
	ID         uint               `json:"id" gorm:"primaryKey"`
	Username   string             `json:"username" gorm:"size:255;not null"`
	Email      string             `json:"email" gorm:"not null;unique"`
	Password   string             `json:"-" gorm:"not null"`
	Role       constants.UserRole `json:"role" gorm:"type:enum('ADMIN','STUDENT');default:STUDENT"`
	Nrp        *string            `json:"nrp"`
	MaxCredits *int               `json:"max_credits"`

	BaseModel

	CoursePlans []CoursePlan `gorm:"foreignKey:StudentID"`
}
