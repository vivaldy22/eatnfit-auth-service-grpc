package token

import (
	"context"
	authservice "github.com/vivaldy22/eatnfit-auth-service/proto"
	"github.com/vivaldy22/eatnfit-auth-service/tools/jwttoken"
)

const (
	customKey = "keyEatNFit"
)

type Service struct {}

func (s Service) GenerateToken(ctx context.Context, credentials *authservice.LoginCredentials) (*authservice.Token, error) {
	token, err := jwttoken.JwtEncoder(credentials.UserEmail, customKey)
	if err != nil {
		return nil, err
	}
	return &authservice.Token{Token: token}, err
}

func NewService() authservice.JWTTokenServer {
	return &Service{}
}