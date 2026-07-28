package models

import (
	"time"
	"yosimaril/CourseEnrollment/constants"
)

type CoursePlan struct {
	ID        uint                       `json:"id" gorm:"primaryKey"`
	StudentID uint                       `json:"student_id"`
	Status    constants.CoursePlanStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Student User             `gorm:"foreignKey:StudentID"`
	Items   []CoursePlanItem `gorm:"foreignKey:CoursePlanID"`
}
