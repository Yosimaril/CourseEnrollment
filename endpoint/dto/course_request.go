package dto

type CreateCourseRequest struct {
	Code    string `json:"code" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Credits int    `json:"credits" binding:"required,min=0"`
}

type UpdateCourseRequest struct {
	Code    *string `json:"code"`
	Name    *string `json:"name"`
	Credits *int    `json:"credits,omitempty" binding:"omitempty,min=0"`
}