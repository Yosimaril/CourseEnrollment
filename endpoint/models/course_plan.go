package models

import (
	"yosimaril/CourseEnrollment/constants"
)

type CoursePlan struct {
	ID        uint                       `json:"id" gorm:"primaryKey"`
	StudentID uint                       `json:"student_id"`
	Status    constants.CoursePlanStatus `json:"status"`

	BaseModel

	Student User             `json:"user" gorm:"foreignKey:StudentID"`
	Items   []CoursePlanItem `json:"course_plan_item" gorm:"foreignKey:CoursePlanID"`
}
