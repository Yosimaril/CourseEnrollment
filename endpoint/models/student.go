package models

type Student struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	StudentID string `json:"studentId" gorm:"unique;not null"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique;not null"`
	Major     string `json:"major"`

	BaseModel
}