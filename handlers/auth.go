package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"main/db"
	"main/models"
	"main/utils"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	database := db.GetDB()
	defer database.Close()

	var existing string
	err = database.QueryRow("SELECT email FROM users WHERE email = ?", user.Email).Scan(&existing)
	if err != sql.ErrNoRows {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	token := utils.GenerateToken()
	_, err = database.Exec("INSERT INTO users (name, email, password, verified, token) VALUES (?, ?, ?, ?, ?)",
		user.Name, user.Email, string(hashedPassword), false, token)
	if err != nil {
		http.Error(w, "Database Error", http.StatusConflict)
		return
	}

	utils.SendMail(user.Email, token)
	w.Write([]byte("Signup successful! Please verify your email."))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	database := db.GetDB()
	defer database.Close()

	var hashedPassword, name string
	var verified bool

	err = database.QueryRow("SELECT name, password, verified FROM users WHERE email = ?", credentials.Email).
		Scan(&name, &hashedPassword, &verified)
	if err != nil || !verified {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateJWT(credentials.Email, name)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
