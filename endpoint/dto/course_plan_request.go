package dto

import "yosimaril/CourseEnrollment/models"

type CreateCoursePlanRequest struct {
	StudentID uint                     `json:"student_id" binding:"required"`
	Status    models.CoursePlanStatus `json:"status" binding:"required,oneof=DRAFT APPROVED PARTIALLY_APPROVED REJECTED SUBMITTED"`
}

type UpdateCoursePlanRequest struct {
	StudentID *uint                     `json:"student_id"`
	Status    *models.CoursePlanStatus `json:"status" binding:"omitempty,oneof=DRAFT APPROVED PARTIALLY_APPROVED REJECTED SUBMITTED"`
}