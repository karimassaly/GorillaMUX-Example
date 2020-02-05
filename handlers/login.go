package handlers

import (
	"Example/database"
	"Example/models"
	"Example/security"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Password or Username missing", http.StatusBadRequest)
		return
	}
	if CheckUserIdentity(user.Username, user.Password) == false {
		http.Error(w, "Username or Password Incorrect", http.StatusUnauthorized)
		return
	}
}

func CheckUserIdentity(username, password string) bool {
	user := &models.User{Username: username, Password: password}
	db, _ := database.InitDB()

	if err := db.Where(&models.User{Username: username}).Find(&user).Error; err == nil {
		Cipherpass, _ := hex.DecodeString(user.Password)
		user.Password = security.UserCypherOFF(Cipherpass)
		if password == user.Password {
			println("Connected welcome", user.Username)
			return true
		}
		return false
	}
	return false
}
