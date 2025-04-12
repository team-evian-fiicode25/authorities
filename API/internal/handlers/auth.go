package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/team-evian-fiicode25/consumer/API/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(client *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: client,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		username string
		password string
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, err)
		return
	}

	login, err := h.authService.LogInWithPassword(context.Background(), request.username, request.password)
	if err != nil {
		writeError(w, err)
		return
	}

	var response struct {
		sessionId string
		login     struct {
			username string
		}
	}

	response.sessionId = login.GetSessionToken().GetToken()
	response.login.username = login.GetLogin().GetUsername()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		match := regexp.MustCompile(`^Bearer\s+(\S+)$`).FindStringSubmatch(header)

		if match == nil || !h.authService.VerifyToken(r.Context(), match[1]) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeError(w http.ResponseWriter, err error) {
	var response struct {
		error string
	}

	response.error = err.Error()

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}
