package apis

import (
	"fmt"
	"net/http"

	"github.com/team-evian-fiicode25/consumer/API/internal/handlers"
	"github.com/team-evian-fiicode25/consumer/API/internal/services"
)

func SetupRoutes() {
	mux := http.NewServeMux()

	authHandler := handlers.NewAuthHandler(services.NewAuthService())

	mux.HandleFunc("/auth/login", authHandler.Login)

	mux.Handle("/hello", authHandler.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Hello, world!"}`))
	})))

	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", mux)
}
