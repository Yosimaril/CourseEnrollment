package dto

import "yosimaril/CourseEnrollment/constants"

type CreateCoursePlanRequest struct {
	StudentID uint                       `json:"student_id" binding:"required"`
	Status    constants.CoursePlanStatus `json:"status" binding:"required,oneof=DRAFT APPROVED PARTIALLY_APPROVED REJECTED SUBMITTED"`
}

type UpdateCoursePlanRequest struct {
	StudentID *uint                       `json:"student_id"`
	Status    *constants.CoursePlanStatus `json:"status" binding:"omitempty,oneof=DRAFT APPROVED PARTIALLY_APPROVED REJECTED SUBMITTED"`
}
