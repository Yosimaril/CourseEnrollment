package models

type Student struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Username   string  `json:"username" gorm:"size:255;not null"`
	Email      string  `json:"email" gorm:"not null;unique"`
	Password   string  `json:"-" gorm:"not null"`
	Nrp        *string `json:"nrp"`
	MaxCredits *int    `json:"max_credits"`

	BaseModel
}
