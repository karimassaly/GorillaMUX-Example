package handlers

import (
	"Example/database"
	"Example/models"
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" || user.Password == "" || user.Name == "" {
		http.Error(w, "Password, Name or Email missing", http.StatusBadRequest)
		return
	}
	if InsertionUser(user.Name, user.Email, user.Password) == false {
		http.Error(w, "This Account already exists", http.StatusConflict)
	}

}

func InsertionUser(name, email, password string) bool {
	db, _ := database.InitDB()
	user := &models.User{Name: name, Email: email, Password: password}
	if err := db.Where(&models.User{Email: email}).Find(&user).Error; err == nil {
		return false
	} else {
		db.Create(&user)
	}
	return true
}
