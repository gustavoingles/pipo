package main

import (
	"errors"
	"log"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPasword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error trying to hash the password", http.StatusInternalServerError)
		return
	}
	hashedPaswordString := string(hashedPasword)

	var usernameDB string
	err = dbpool.QueryRow(r.Context(), "SELECT username FROM users WHERE username = $1", username).Scan(&usernameDB)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			_, err := dbpool.Exec(r.Context(), "INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPaswordString)
			if err != nil {
				log.Printf("Error inserting data into the database: %v", err)
				http.Error(w, "Error inserting data into the database", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			return
		}

		log.Printf("User already exists, please choose another username or log in")
		http.Error(w, "User already exists, please choose another username or log in", http.StatusConflict)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		""
	})
}
