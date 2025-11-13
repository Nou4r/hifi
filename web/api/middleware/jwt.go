package middleware

import (
	"api/config"
	"api/types"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	users = make(map[string]*types.User)
	mu    sync.RWMutex

	results      = make(chan types.RegisterResult)
	registerJobs = make(chan types.SignupRequest, 10)

	tokenHashes = make(map[string][32]byte)
)

func RegistrationWorker() {
	for creds := range registerJobs {
		mu.Lock()
		if _, exists := users[creds.Username]; exists {
			mu.Unlock()
			results <- types.RegisterResult{Success: false, Message: "⚠️ User already exists"}
			continue
		}

		user := &types.User{
			ID:       uuid.NewString(),
			Username: creds.Username,
			Password: creds.Password,
		}
		users[creds.Username] = user
		mu.Unlock()

		results <- types.RegisterResult{Success: true, Message: "✅ Registered user"}
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds types.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	select {
	case registerJobs <- creds:
		res := <-results

		w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
		status := http.StatusCreated

		if !res.Success {
			status = http.StatusBadRequest
		}

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"message": res.Message,
		})

	default:
		http.Error(w, "Server busy, try again later", http.StatusServiceUnavailable)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds types.SigninRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	mu.RLock()
	user, ok := users[creds.Username]
	mu.RUnlock()

	if !ok || user.Password != creds.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	maxAge := 3600
	expiration := time.Now().Add(time.Second * time.Duration(maxAge))

	claims := &types.Claims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
	}

	hash := sha256.Sum256([]byte(claims.RegisteredClaims.ID))
	tokenHashes[claims.RegisteredClaims.ID] = hash

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"token":    tokenString,
		"username": user.Username,
		"maxAge":   maxAge,
	}
	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(resp)
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	var tokenString string
	fmt.Sscanf(authHeader, "Bearer %s", &tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &types.Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(*types.Claims)
	if !ok {
		http.Error(w, "Invalid claims", http.StatusBadRequest)
		return
	}

	computed := sha256.Sum256([]byte(claims.RegisteredClaims.ID))
	stored, ok := tokenHashes[claims.RegisteredClaims.ID]
	if !ok || subtle.ConstantTimeCompare(stored[:], computed[:]) != 1 {
		http.Error(w, "Token Unknown", http.StatusUnauthorized)
		return
	}

	mu.RLock()
	user, exists := users[claims.Username]
	mu.RUnlock()

	if !exists || user.ID != claims.ID {
		http.Error(w, "User does not exist or data mismatch", http.StatusBadRequest)
		return
	}

	resp := map[string]any{
		"description": "Welcome " + claims.Username + "! Verified via JWT.",
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(resp)
}
