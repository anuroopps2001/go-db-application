package models

type User struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name" gorm:"not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	Age          int    `json:"age"`
	ProfileImage string `json:"profile_image"`
}

type Userparam struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
