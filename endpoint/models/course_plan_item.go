package models

import (
	"time"
	"yosimaril/CourseEnrollment/constants"
)

type CoursePlanItem struct {
	CoursePlanID uint                           `json:"course_plan_id" gorm:"primaryKey"`
	CourseID     uint                           `json:"course_id" gorm:"primaryKey"`
	Status       constants.CoursePlanItemStatus `json:"status"`
	Remarks      *string                        `json:"remarks"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CoursePlan CoursePlan `gorm:"foreignKey:CoursePlanID"`
	Course     Course     `gorm:"foreignKey:CourseID"`
}
