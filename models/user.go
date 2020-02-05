package models

type User struct {
	ID       int    `gorm:primary_key;"AUTO_INCREMENT"`
	Username string `json:"username"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
