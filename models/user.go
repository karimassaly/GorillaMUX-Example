package models

type User struct {
	ID       int    `gorm:primary_key;"AUTO_INCREMENT"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
