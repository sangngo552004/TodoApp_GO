package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}
