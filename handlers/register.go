package handlers

import (
	"Example/database"
	"Example/models"
	"Example/security"
	"encoding/json"
	"net/http"
	"regexp"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(user.Email) == false {
		http.Error(w, "Email Format is Invalid", http.StatusBadRequest)
		return
	}
	if security.IsAlphaNumeric(user.Username) == false {
		http.Error(w, "Username Format is Invalid", http.StatusBadRequest)
		return
	}
	return
	if user.Email == "" || user.Password == "" || user.Username == "" {
		http.Error(w, "Password, Username or Email missing", http.StatusBadRequest)
		return
	}
	UserSuccess := InsertionUser(user.Username, user.Email, user.Password)
	switch UserSuccess {
	case 0:
		http.Error(w, "Registration Successful", http.StatusOK)
	case 1:
		http.Error(w, "This Account already exists", http.StatusConflict)
	case 2:
		http.Error(w, "This Username already exists", http.StatusConflict)
	}

}

func InsertionUser(username, email, password string) int {
	db, _ := database.InitDB()
	user := &models.User{Username: username, Email: email, Password: password}
	if err := db.Where(&models.User{Email: email}).Find(&user).Error; err == nil {
		return 1
	} else if err := db.Where(&models.User{Username: username}).Find(&user).Error; err == nil {
		return 2
	} else {
		user.Password = security.UserCypherON(user.Password)
		db.Create(&user)
		return 0
	}
}
