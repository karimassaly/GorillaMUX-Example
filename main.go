package main

import (
	"Example/database"
	"Example/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	db, _ := database.InitDB()
	database.CreateTable(db)
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
