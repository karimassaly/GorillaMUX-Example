package handlers

import (
	"Example/database"
	"Example/models"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" || user.Password == "" {
		http.Error(w, "Password or Email missing", http.StatusBadRequest)
		return
	}
	if CheckUserIdentity(user.Email, user.Password) == false {
		http.Error(w, "Email or Password Incorrect", http.StatusUnauthorized)
		return
	}
}

func CheckUserIdentity(email, password string) bool {
	user := &models.User{Email: email, Password: password}
	db, _ := database.InitDB()

	if err := db.Where(&models.User{Email: email}).Find(&user).Error; err == nil {
		if password == user.Password {
			println("Connected welcome", user.Email)
			return true
		}
		return false
	}
	return false
}
