package handlers

import (
	"fmt"
	"jwt-example/jwt"
	"net/http"
	"time"
)

var revokedTokens = make(map[string]time.Time)

func RevokeJWT(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == "" {
		http.Error(w, "Token required", http.StatusBadRequest)
		return
	}

	revokedTokens[token] = time.Now().Add(30 * time.Minute)
	_, err := w.Write([]byte("Token revoked"))
	if err != nil {
		return
	}
}

func RotateJWT(w http.ResponseWriter, r *http.Request) {
	oldToken := r.Header.Get("Authorization")
	if oldToken == "" {
		http.Error(w, "Token required", http.StatusBadRequest)
		return
	}

	claims, err := jwt.ParseJWT(oldToken)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if revokedTokens[oldToken].After(time.Now()) {
		http.Error(w, "Token has been revoked", http.StatusUnauthorized)
		return
	}

	newToken, err := jwt.GenerateJWT(claims.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", newToken)
	_, err = w.Write([]byte(fmt.Sprintf("New Token: %s", newToken)))
	if err != nil {
		return
	}
}
