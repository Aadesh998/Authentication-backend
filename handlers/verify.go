package handlers

import (
	"net/http"

	"main/db"
)

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	database := db.GetDB()
	defer database.Close()

	res, err := database.Exec("UPDATE users SET verified = true WHERE token = ?", token)
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Email verified successfully! You can now log in."))
}
