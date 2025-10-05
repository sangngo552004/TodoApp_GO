package models

type Todo struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
}
