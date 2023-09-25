package user

import (
	"context"

	"shifo-backend-website/internal/auth"
	"shifo-backend-website/internal/pkg"
	"shifo-backend-website/internal/repository/postgres/user"
)

type User interface {
	AdminGetList(ctx context.Context, filter user.Filter) ([]user.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (user.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request user.AdminCreateRequest) (user.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request user.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}

type Auth interface {
	GetTokenData(ctx context.Context, token string) (auth.TokenData, error)
}
