package handlers

import (
	"encoding/json"
	"latlongapi/backend/auth"
	"latlongapi/backend/models"
	"latlongapi/backend/store"
	"log"
	"net/http"
	"strings"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	userStore models.UserStore
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(userStore models.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

// RegisterRequest represents a registration request
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents an authentication response
type AuthResponse struct {
	Token string      `json:"token"`
	User  *models.User `json:"user"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		respondError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 6 {
		respondError(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	_, err := h.userStore.GetUserByEmail(req.Email)
	if err == nil {
		respondError(w, "User already exists", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		respondError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create user
	user, err := h.userStore.CreateUser(req.Email, hashedPassword)
	if err != nil {
		if err == store.ErrUserExists {
			respondError(w, "User already exists", http.StatusConflict)
			return
		}
		log.Printf("Error creating user: %v", err)
		respondError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		respondError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJSON(w, AuthResponse{
		Token: token,
		User:  user,
	}, http.StatusCreated)
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		respondError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Get user
	user, err := h.userStore.GetUserByEmail(req.Email)
	if err != nil {
		if err == store.ErrUserNotFound {
			respondError(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		log.Printf("Error getting user: %v", err)
		respondError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Check password
	if !auth.CheckPasswordHash(req.Password, user.Password) {
		respondError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		respondError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJSON(w, AuthResponse{
		Token: token,
		User:  user,
	}, http.StatusOK)
}

// Me returns the current authenticated user
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user from context (set by middleware)
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		respondError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	respondJSON(w, user, http.StatusOK)
}

// Logout handles user logout (client-side token removal)
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Logout is handled client-side by removing the token
	// This endpoint just confirms the logout
	respondJSON(w, map[string]string{"message": "Logged out successfully"}, http.StatusOK)
}

// Helper functions

func respondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func respondError(w http.ResponseWriter, message string, statusCode int) {
	respondJSON(w, ErrorResponse{Error: message}, statusCode)
}

// GetTokenFromRequest extracts JWT token from request
func GetTokenFromRequest(r *http.Request) string {
	// Try Authorization header first
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	// Try cookie
	cookie, err := r.Cookie("token")
	if err == nil {
		return cookie.Value
	}

	return ""
}

