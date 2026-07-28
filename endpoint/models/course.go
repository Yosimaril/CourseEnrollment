package models

type Course struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Credits int    `json:"credits"`

	BaseModel

	CoursePlanItems []CoursePlanItem `gorm:"foreignKey:CourseID"`
}
