package auth

import (
	"context"
	"shifo-backend-website/internal/auth"
	"shifo-backend-website/internal/entity"
	"shifo-backend-website/internal/pkg"
	"shifo-backend-website/internal/repository/postgres/user"
)

type Auth interface {
	GenerateToken(ctx context.Context, data auth.GenerateToken) (string, error)
	IsValidToken(ctx context.Context, token string) (entity.User, error)
	GetTokenData(ctx context.Context, token string) (auth.TokenData, error)
}

type User interface {
	GetByFirstName(ctx context.Context, firstName string) (user.AdminGetDetail, *pkg.Error)
}
