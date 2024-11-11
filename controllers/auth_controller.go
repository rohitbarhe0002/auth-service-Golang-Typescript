package controllers

import (
	"auth-service/models"
	"auth-service/repository"
	"auth-service/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passwordHash, err := utils.HashPassword(user.Password) 
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}


	fmt.Printf("Raw Password: %s\n", user.Password) 
	fmt.Printf("Stored Password Hash: %s\n", passwordHash)


	user.PasswordHash = passwordHash

	if err := repository.CreateUser(&user); err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}


func SignIn(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := repository.FindUserByEmail(credentials.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	passwordMatch := utils.CheckPasswordHash(credentials.Password, user.PasswordHash)
	if !passwordMatch {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Email)
	if err != nil {
		http.Error(w, "Could not generate refresh token", http.StatusInternalServerError)
		return
	}


	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}



func Protected(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "You have accessed a protected route"}
	json.NewEncoder(w).Encode(response)
}


func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
    RefreshToken string `json:"refresh_token"`
}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email, err := utils.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateJWT(email)
	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return
	}


	response := map[string]string{
		"access_token": accessToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
