package models

import (
	"time"
)

type CoursePlan struct {
	ID        uint             `json:"id" gorm:"primaryKey"`
	StudentID uint             `json:"student_id"`
	Status    CoursePlanStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Student User             `gorm:"foreignKey:StudentID"`
	Items   []CoursePlanItem `gorm:"foreignKey:CoursePlanID"`
}
