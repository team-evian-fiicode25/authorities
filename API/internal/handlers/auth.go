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
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, err)
		return
	}

	login, err := h.authService.LogInWithPassword(context.Background(), request.Username, request.Password)
	if err != nil {
		writeError(w, err)
		return
	}

	var response struct {
		SessionId string `json:"sessionId"`
		Login     struct {
			Username string `json:"username"`
		} `json:"login"`
	}

	response.SessionId = login.GetSessionToken().GetToken()
	response.Login.Username = login.GetLogin().GetUsername()

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
		Error string `json:"error"`
	}

	response.Error = err.Error()

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}
