package models

import (
	"yosimaril/CourseEnrollment/constants"
)

type CoursePlan struct {
	ID        uint                       `json:"id" gorm:"primaryKey"`
	StudentID uint                       `json:"student_id"`
	Status    constants.CoursePlanStatus `json:"status"`

	BaseModel

	Student User             `gorm:"foreignKey:StudentID"`
	Items   []CoursePlanItem `gorm:"foreignKey:CoursePlanID"`
}
