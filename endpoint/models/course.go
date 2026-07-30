package models

type Course struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Credits int    `json:"credits"`

	BaseModel

	CoursePlanItems []CoursePlanItem `json:"course_plan_items" gorm:"foreignKey:CourseID"`
}
