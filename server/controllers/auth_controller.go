package controllers

import (
	"encoding/json"
	"net/http"
	"tecnhical-test/config"
	"tecnhical-test/helpers"
	"tecnhical-test/middlewares"
	"tecnhical-test/models"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Fullname string `json:"full_name"`
	Role     string `json:"role"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		middlewares.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Username == "" || req.Password == "" {
		middlewares.ErrorResponse(w, http.StatusUnprocessableEntity, "Username and password required")
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		middlewares.ErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := middlewares.GenerateJWT(user)
	if err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := AuthResponse{
		Token: token,
		User:  user,
	}

	middlewares.SuccessResponse(w, "Login successful", response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		middlewares.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Username == "" || req.Password == "" || req.Email == "" {
		middlewares.ErrorResponse(w, http.StatusUnprocessableEntity, "Username, password, and email are required")
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	if req.Role == "" {
		req.Role = "staff"
	}

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Fullname: req.Fullname,
		Role:     req.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	token, err := middlewares.GenerateJWT(user)
	if err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := AuthResponse{
		Token: token,
		User:  user,
	}

	middlewares.SuccessResponse(w, "Registration successful", response)
}
