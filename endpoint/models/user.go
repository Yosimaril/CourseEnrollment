package models

type User struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Username   string   `gorm:"size:255;not null"`
	Email      string   `json:"email" gorm:"not null;unique"`
	Password   string   `json:"-" gorm:"not null"`
	Role       UserRole `json:"role" gorm:"type:enum('ADMIN','STUDENT');default:STUDENT"`
	Nrp        *string  `json:"nrp"`
	MaxCredits *int     `json:"max_credits"`

	BaseModel

	CoursePlans []CoursePlan `gorm:"foreignKey:StudentID"`
}
