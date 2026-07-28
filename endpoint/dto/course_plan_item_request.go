package dto

import "yosimaril/CourseEnrollment/constants"

type CreateCoursePlanItemRequest struct {
	CoursePlanID uint                           `json:"course_plan_id" binding:"required"`
	CourseID     uint                           `json:"course_id" binding:"required"`
	Status       constants.CoursePlanItemStatus `json:"status" binding:"required,oneof=PENDING APPROVED REJECTED"`
	Remarks      *string                        `json:"remarks"`
}

type UpdateCoursePlanItemRequest struct {
	Status  *constants.CoursePlanItemStatus `json:"status" binding:"omitempty,oneof=PENDING APPROVED REJECTED"`
	Remarks *string                         `json:"remarks"`
}
