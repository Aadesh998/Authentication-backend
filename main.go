package main

import (
	"net/http"

	"main/handlers"
	"main/middleware"
)

func main() {
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/verify", handlers.VerifyHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/protected", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a protected route!"))
	}))

	http.ListenAndServe(":8000", nil)
}
