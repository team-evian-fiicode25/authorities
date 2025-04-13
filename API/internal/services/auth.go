package services

import (
	"context"
	"log"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/team-evian-fiicode25/consumer/API/internal/config"
)

type AuthService struct {
	client graphql.Client
}

func NewAuthService() *AuthService {
	url, err := config.AuthServiceUrl()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &AuthService{
		client: graphql.NewClient(url, http.DefaultClient),
	}
}

func (s *AuthService) LogInWithPassword(ctx context.Context, username string, password string) (LogInWithUsernameLoginSessionIQueryableLoginSession, error) {
	response, err := LogInWithUsername(ctx, s.client, username, password)
	if err != nil {
		return nil, err
	}

	return response.GetLoginSession(), nil
}

func (s *AuthService) VerifyToken(ctx context.Context, token string) bool {
	resp, err := VerifyToken(ctx, s.client, token)

	if err != nil || resp.GetLogin() == nil {
		return false
	}

	return true
}
