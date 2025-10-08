package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}
